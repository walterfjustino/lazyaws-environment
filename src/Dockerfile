FROM golang:1.24-alpine

# Dependências básicas
RUN apk add --no-cache \
    bash \
    unzip \
    wget \
    ca-certificates \
    less \
    groff

# -------------------------
# Instalar AWS CLI v2 (via wget - evita bug do curl)
# -------------------------
RUN wget https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -O awscliv2.zip \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf aws awscliv2.zip

# -------------------------
# Session Manager Plugin
# -------------------------
RUN wget https://s3.amazonaws.com/session-manager-downloads/plugin/latest/linux_64bit/session-manager-plugin \
    -O /usr/local/bin/session-manager-plugin \
    && chmod +x /usr/local/bin/session-manager-plugin

WORKDIR /workspace

CMD ["sh"]
