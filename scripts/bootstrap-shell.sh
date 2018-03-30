#!/bin/bash
set -euo pipefail

# Define variables.
BASH_GO="~/.bash_go"

# Update .bash_profile if need be.
if ! grep "[ -f ${BASH_GO} ] && . ${BASH_GO}" ~/.bash_profile; then
  echo '' >> ~/.bash_profile
  echo '# Configure Go.' >> ~/.bash_profile
  echo "[ -f ${BASH_GO} ] && . ${BASH_GO}" >> ~/.bash_profile
fi

# Update the $PATH only if needed.
if ! echo $PATH | grep -i "$(go env GOPATH)"; then
  PATH=$PATH:$(go env GOPATH)/bin
fi

# Create the dotfile for bash_go.
cat <<EOF > "${HOME}/.bash_go"
export GOPATH=$(go env GOPATH)
export PATH=$PATH
EOF
source "${HOME}/.bash_go"
