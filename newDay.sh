#! /bin/bash

mkdir $1
cd $1
go mod init $1
go get -u github.com/spf13/cobra/cobra
cobra init
cobra add solve
