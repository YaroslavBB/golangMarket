#!/bin/bash
export CONFIG_TYPE=""

echo 'COMPILING...'
rm ../bin/market

cd ../
go build -o bin/market || { echo 'build failed' ; exit 1; }

echo 'RUN'
./bin/market