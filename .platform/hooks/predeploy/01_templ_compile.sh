#!/usr/bin/env bash

# Download dependencies
go mod download

# Install templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Compile templ
templ generate