#!/bin/bash

# Quick build script for billing_go on Linux AMD64 VPS
# Usage: ./quick_build.sh

set -e  # Exit on error

echo "=== Quick Build for billing_go ==="
echo "Target: Linux AMD64 VPS"
echo ""

# Check if running on Linux
if [[ "$OSTYPE" != "linux-gnu"* ]]; then
    echo "Error: This script is designed for Linux systems only"
    exit 1
fi

# Clean previous build artifacts
echo "Cleaning previous build artifacts..."
make clean 2>/dev/null || true

# Build the project with production flags for VPS compatibility
echo "Building billing_go for production deployment..."
echo "Using CGO_ENABLED=0 for static binary compatible with older GLIBC"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o billing

# Check if build was successful
if [ -f "./billing" ]; then
    echo ""
    echo "✓ Build successful!"
    echo "Binary created: ./billing"
    
    # Make binary executable
    chmod +x billing
    
    # Show binary info
    echo ""
    echo "Binary information:"
    file billing
    
    echo ""
    echo "Usage:"
    echo "  ./billing         - Run in foreground"
    echo "  ./billing up -d   - Run as daemon"
    echo "  ./billing stop    - Stop service"
    echo "  ./billing version - Show version"
else
    echo ""
    echo "✗ Build failed!"
    exit 1
fi

echo ""
echo "Build completed at: $(date)"