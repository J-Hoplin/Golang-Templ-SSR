#!/usr/bin bash

# Install cURL
sudo yum -y install curl
# Install Node.js
yum install nodejs

# Enroll Go binary path to Path
export PATH=$PATH:$HOME/go/bin
source ~/.bashrc

# templ CLI
go install github.com/a-h/templ/cmd/templ@latest