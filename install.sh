#!/bin/bash
set -e

REPO="fardhanrasya/dream-journal"
INSTALL_DIR="$HOME/.dream-journal/bin"
BIN_NAME="dream"

echo "Fetching latest release..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest")
TAG_NAME=$(echo "$LATEST_RELEASE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$TAG_NAME" ]; then
    echo "Error: Could not find latest release tag."
    exit 1
fi

echo "Latest version: $TAG_NAME"

OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" == "Linux" ]; then
    ASSET_OS="Linux"
elif [ "$OS" == "Darwin" ]; then
    ASSET_OS="Darwin"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

if [ "$ARCH" == "x86_64" ]; then
    ASSET_ARCH="x86_64"
elif [ "$ARCH" == "aarch64" ] || [ "$ARCH" == "arm64" ]; then
    ASSET_ARCH="arm64"
else
    echo "Unsupported Architecture: $ARCH"
    exit 1
fi

ASSET_NAME="dream-journal_${ASSET_OS}_${ASSET_ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$TAG_NAME/$ASSET_NAME"

mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR"

echo "Downloading $DOWNLOAD_URL..."
curl -sL "$DOWNLOAD_URL" | tar xz

if [ ! -f "$BIN_NAME" ]; then
    echo "Error: Extraction failed or binary name mismatch."
    exit 1
fi

chmod +x "$BIN_NAME"

echo "Checking PATH..."
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Adding to PATH..."
    SHELL_CONFIG=""
    if [ -f "$HOME/.zshrc" ]; then
        SHELL_CONFIG="$HOME/.zshrc"
    elif [ -f "$HOME/.bashrc" ]; then
        SHELL_CONFIG="$HOME/.bashrc"
    fi

    if [ -n "$SHELL_CONFIG" ]; then
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_CONFIG"
        echo "Added to $SHELL_CONFIG. Please restart your terminal or run 'source $SHELL_CONFIG'"
    else
        echo "Could not detect shell config. Please manually add '$INSTALL_DIR' to your PATH."
    fi
fi

# Ensure EDITOR is set
if [ -z "$EDITOR" ]; then
    echo "EDITOR environment variable not set. Adding default (nano) to shell config..."
    if [ -n "$SHELL_CONFIG" ]; then
        echo "export EDITOR=nano" >> "$SHELL_CONFIG"
    fi
fi

echo ""
echo "Success! Dream Journal $TAG_NAME installed."
echo "Run 'dream' to start."
