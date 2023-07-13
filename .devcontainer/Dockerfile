ARG GO_VERSION=1.20
FROM mcr.microsoft.com/devcontainers/go:0-${GO_VERSION}

# install dependencies to build
RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev \
    xorg-dev \
    && rm -rf /var/lib/apt/lists/*

# install tooling commands
# RUN go install fyne.io/fyne/v2/cmd/fyne@latest
