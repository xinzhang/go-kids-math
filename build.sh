#!/bin/bash

# Build script for kids learning programs
# Builds all Go programs and copies config to dist folder

set -e  # Exit on error

echo "Building kids learning programs..."

# Create dist directory if it doesn't exist
mkdir -p dist

# Build each Go program
echo "Building mathtest..."
go build -o dist/mathtest mathtest.go

echo "Building multiply..."
go build -o dist/multiply multiply.go

echo "Building thinkingskills..."
go build -o dist/thinkingskills thinkingskills.go

echo "Building algebra..."
go build -o dist/algebra algebra.go

echo "Building tstrial..."
go build -o dist/tstrial tstrial.go

# Copy config file
echo "Copying config file..."
cp config dist/config

echo ""
echo "✓ Build complete!"
echo "All programs are in the dist/ folder:"
echo "  - mathtest"
echo "  - multiply"
echo "  - thinkingskills"
echo "  - algebra"
echo "  - tstrial"
echo "  - config"
