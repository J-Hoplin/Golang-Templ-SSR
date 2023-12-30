#!/usr/bin/env bash


export PATH=$PATH:$HOME/go/bin
source ~/.bashrc
go install github.com/a-h/templ/cmd/templ@latest

templ generate