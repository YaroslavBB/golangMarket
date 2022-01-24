#!/bin/bash
source variables.sh

echo 'COMPILING...'
rm $ROOT/bin/market

cd $ROOT/cmd
go build -o $ROOT/bin/market || { echo 'build failed' ; exit 1; }

echo 'RUN'
cd $ROOT/bin
./market