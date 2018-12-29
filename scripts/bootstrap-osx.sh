#!/bin/bash
set -euo pipefail

# Install Brew.
brew --version || /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Add new taps.
brew tap alecthomas/homebrew-tap

# Update Brew.
brew update

# Install brew formulas.
brew install go dep gometalinter
