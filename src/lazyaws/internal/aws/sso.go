package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/skratchdot/open-golang/open"
)

const (
	// Default SSO configuration
	DefaultSSORegion = "us-east-1"
	ClientName       = "lazyaws"
	ClientType       = "public"
)

// SSOSession represents an active SSO session with cached tokens
type SSOSession struct {
	StartURL              string    `json:"start_url"`
	Region                string    `json:"region"`
	AccessToken           string    `json:"access_token"`
	ExpiresAt             time.Time `json:"expires_at"`
	ClientID              string    `json:"client_id"`
	ClientSecret          string    `json:"client_secret"`
	RegistrationExpiresAt time.Time `json:"registration_expires_at"`
}

// SSOAccount represents an AWS account available through SSO
type SSOAccount struct {
	AccountID    string
	AccountName  string
	EmailAddress string
	RoleName     string
}

// SSOAuthenticator handles AWS SSO authentication
type SSOAuthenticator struct {
	startURL string
	region   string
	session  *SSOSession
}

// NewSSOAuthenticator creates a new SSO authenticator
func NewSSOAuthenticator(startURL, region string) *SSOAuthenticator {
	if region == "" {
		region = DefaultSSORegion
	}

	return &SSOAuthenticator{
		startURL: startURL,
		region:   region,
	}
}

// Authenticate performs the SSO device authorization flow
func (a *SSOAuthenticator) Authenticate(ctx context.Context) error {
	// Try to load cached session first
	if err := a.loadCachedSession(); err == nil &&
		a.session.AccessToken != "" &&
		time.Now().Before(a.session.ExpiresAt) {
		return nil
	}

	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(a.region))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	oidcClient := ssooidc.NewFromConfig(cfg)

	// Register client if needed
	if a.session == nil || time.Now().After(a.session.RegistrationExpiresAt) {
		if err := a.registerClient(ctx, oidcClient); err != nil {
			return fmt.Errorf("failed to register SSO client: %w", err)
		}
	}

	deviceAuthResp, err := oidcClient.StartDeviceAuthorization(ctx, &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     &a.session.ClientID,
		ClientSecret: &a.session.ClientSecret,
		StartUrl:     &a.startURL,
	})
	if err != nil {
		return fmt.Errorf("failed to start device authorization: %w", err)
	}

	verificationURL := *deviceAuthResp.VerificationUriComplete
	if err := openBrowser(verificationURL); err != nil {
		fmt.Fprintf(os.Stderr, "\nUnable to open browser automatically.\n")
		fmt.Fprintf(os.Stderr, "Please navigate to the following URL:\n\n  %s\n\n", verificationURL)
	}

	deviceCode := *deviceAuthResp.DeviceCode
	interval := time.Duration(deviceAuthResp.Interval) * time.Second
	expiresAt := time.Now().Add(time.Duration(deviceAuthResp.ExpiresIn) * time.Second)

	for time.Now().Before(expiresAt) {
		time.Sleep(interval)

		tokenResp, err := oidcClient.CreateToken(ctx, &ssooidc.CreateTokenInput{
			ClientId:     &a.session.ClientID,
			ClientSecret: &a.session.ClientSecret,
			GrantType:    strPtr("urn:ietf:params:oauth:grant-type:device_code"),
			DeviceCode:   &deviceCode,
		})

		if err != nil {
			if isAuthPending(err) {
				continue
			}
			return fmt.Errorf("failed to create token: %w", err)
		}

		a.session.AccessToken = *tokenResp.AccessToken
		a.session.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

		if err := a.saveCachedSession(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to cache SSO session: %v\n", err)
		}

		return nil
	}

	return fmt.Errorf("SSO authorization timed out")
}

// ListAccounts returns all AWS accounts available through SSO
func (a *SSOAuthenticator) ListAccounts(ctx context.Context) ([]SSOAccount, error) {
	if a.session == nil || a.session.AccessToken == "" {
		return nil, fmt.Errorf("not authenticated")
	}

	if time.Now().After(a.session.ExpiresAt) {
		return nil, fmt.Errorf("SSO session expired")
	}

	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(a.region))
	if err != nil {
		return nil, err
	}

	ssoClient := sso.NewFromConfig(cfg)

	resp, err := ssoClient.ListAccounts(ctx, &sso.ListAccountsInput{
		AccessToken: &a.session.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	var accounts []SSOAccount
	for _, acc := range resp.AccountList {
		roles, err := ssoClient.ListAccountRoles(ctx, &sso.ListAccountRolesInput{
			AccessToken: &a.session.AccessToken,
			AccountId:   acc.AccountId,
		})
		if err != nil {
			continue
		}

		for _, role := range roles.RoleList {
			accounts = append(accounts, SSOAccount{
				AccountID:    *acc.AccountId,
				AccountName:  *acc.AccountName,
				EmailAddress: *acc.EmailAddress,
				RoleName:     *role.RoleName,
			})
		}
	}

	return accounts, nil
}

// SSOCredentials represents credentials from SSO
type SSOCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

// GetCredentials retrieves temporary credentials
func (a *SSOAuthenticator) GetCredentials(ctx context.Context, accountID, roleName string) (*SSOCredentials, error) {
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(a.region))
	if err != nil {
		return nil, err
	}

	ssoClient := sso.NewFromConfig(cfg)

	resp, err := ssoClient.GetRoleCredentials(ctx, &sso.GetRoleCredentialsInput{
		AccessToken: &a.session.AccessToken,
		AccountId:   &accountID,
		RoleName:    &roleName,
	})
	if err != nil {
		return nil, err
	}

	creds := resp.RoleCredentials

	return &SSOCredentials{
		AccessKeyID:     *creds.AccessKeyId,
		SecretAccessKey: *creds.SecretAccessKey,
		SessionToken:    *creds.SessionToken,
		Expiration:      time.UnixMilli(creds.Expiration),
	}, nil
}

func (a *SSOAuthenticator) registerClient(ctx context.Context, oidc *ssooidc.Client) error {
	resp, err := oidc.RegisterClient(ctx, &ssooidc.RegisterClientInput{
		ClientName: strPtr(ClientName),
		ClientType: strPtr(ClientType),
	})
	if err != nil {
		return err
	}

	if a.session == nil {
		a.session = &SSOSession{
			StartURL: a.startURL,
			Region:   a.region,
		}
	}

	a.session.ClientID = *resp.ClientId
	a.session.ClientSecret = *resp.ClientSecret
	a.session.RegistrationExpiresAt = time.Unix(resp.ClientSecretExpiresAt, 0)

	return nil
}

func (a *SSOAuthenticator) getCacheFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".lazyaws", "sso-cache")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(dir, fmt.Sprintf("session-%s.json", hashString(a.startURL))), nil
}

func (a *SSOAuthenticator) loadCachedSession() error {
	path, err := a.getCacheFilePath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var s SSOSession
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	a.session = &s
	return nil
}

func (a *SSOAuthenticator) saveCachedSession() error {
	path, err := a.getCacheFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(a.session)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func openBrowser(url string) error {
	return open.Run(url)
}

func isAuthPending(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "AuthorizationPendingException") ||
		strings.Contains(msg, "authorization_pending")
}

func hashString(s string) string {
	var h uint32
	for _, c := range s {
		h = h*31 + uint32(c)
	}
	return fmt.Sprintf("%08x", h)
}

func strPtr(s string) *string {
	return &s
}
