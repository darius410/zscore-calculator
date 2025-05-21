# Z-Score Calculator

A simple command-line tool to calculate z-scores made in Golang.

##Prerequisites
You must have Go installed (this is for the 'Go Install' method below)

## Installation

Download the pre-built binary from the releases page or build from source:

```bash
go install github.com/darius/zscore-calculator/cmd/zscore@latest

## Verify installation (may require adding GOPATH/bin to PATH)
zscore -version

#If go install worked but you still cant do zscore run these commands

1.) Check Your path
echo $GOPATH
2.)Make sure the bin directory is in your PATH
echo $PATH | grep "$GOPATH/bin"
3.)
export PATH=$PATH:$GOPATH/bin

##Happy Scoring!
