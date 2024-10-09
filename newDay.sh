#! /bin/bash

mkdir day$1
cd day$1
go mod init day$1
go get -u github.com/spf13/cobra/cobra
cobra-cli init
cobra-cli add solve
