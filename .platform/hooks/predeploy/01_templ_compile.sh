#!/usr/bin/env bash

# Install templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Compile templ
templ generate