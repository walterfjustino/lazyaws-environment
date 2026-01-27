FROM golang:1.24-bookworm
# -------------------------
# Dependências básicas
# -------------------------
RUN apt-get update && apt-get install -y \
    bash \
    curl \
    unzip \
    ca-certificates \
    less \
    groff \
    wget \
    git \
 && rm -rf /var/lib/apt/lists/*

# -------------------------
# AWS CLI v2 (glibc compatível)
# -------------------------
RUN wget https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -O /tmp/awscliv2.zip \
 && unzip /tmp/awscliv2.zip \
 && ./aws/install \
 && rm -rf /tmp/awscliv2.zip aws

# -------------------------
# kubectl
# -------------------------
RUN curl -fsSL https://dl.k8s.io/release/stable.txt -o /tmp/k8s_version \
 && curl -fsSL https://dl.k8s.io/release/$(cat /tmp/k8s_version)/bin/linux/amd64/kubectl \
    -o /usr/local/bin/kubectl \
 && chmod +x /usr/local/bin/kubectl \
 && rm /tmp/k8s_version

# -------------------------
# k9s
# -------------------------
RUN wget https://github.com/derailed/k9s/releases/latest/download/k9s_Linux_amd64.tar.gz -O /tmp/k9s.tar.gz \
 && tar -xzf /tmp/k9s.tar.gz -C /usr/local/bin k9s \
 && chmod +x /usr/local/bin/k9s \
 && rm /tmp/k9s.tar.gz

# -------------------------
# LazyAWS
# -------------------------
COPY src/lazyaws /tmp/lazyaws
RUN cd /tmp/lazyaws \
 && go mod tidy \
 && go build -o /usr/local/bin/lazyaws . \
 && chmod +x /usr/local/bin/lazyaws \
 && rm -rf /tmp/lazyaws

# -------------------------
# Workspace
# -------------------------
WORKDIR /workspace

CMD ["sh"]
