#!/usr/bin/env sh

# install tooling commands
go install fyne.io/fyne/v2/cmd/fyne@latest

# install direct dependencies
go install
