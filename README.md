# Z-Score Calculator

A simple command-line tool to calculate z-scores made in Golang.

# #How To Use
```bash
## Input
zscore -value 85 -mean 100 -stddev 15
## Expected Outut
Z-Score: -1.0000
```

## Prerequisites
You must have Go installed (this is for the 'Go Install' method below)

## Installation

Download the pre-built binary from the releases page or build from source:

```bash
go install github.com/darius/zscore-calculator/cmd/zscore@latest


## Help Command
zscore -h

## Verify installation (may require adding GOPATH/bin to PATH)
zscore -version

#If go install worked but you still cant do zscore run these commands

1.) Check Your path
echo $GOPATH
2.)Make sure the bin directory is in your PATH
echo $PATH | grep "$GOPATH/bin"
3.)
export PATH=$PATH:$GOPATH/bin
```
##LINUX / MAC OS INSTRUCTIONS

If you are using a linux or Mac based OS try using these commands after using the ones above for installation

```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc  # or ~/.zshrc

// THEN AFTERWARDS RUN THIS

source ~/.bashrc
```
## Happy Scoring! :bar_chart:
