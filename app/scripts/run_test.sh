#!/bin/bash

export ROOT=../../..
# source variables.sh
export DEBUG=true
export DEV=true

if [ -z "$1" ]
  then
    echo "Укажите название модуля"
    exit 1
fi

MODULE=$1
FUNC_NAME=""

if [ -n "$2" ]
  then
    FUNC_NAME="--run $2"
    echo $FUNC_NAME
fi

cd ../internal/modules/$MODULE
go test -v $FUNC_NAME