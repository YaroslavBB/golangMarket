#!/usr/bin/env bash

if [ -z "$1" ]
  then
    echo "Укажите название модуля"
    exit 1
fi


MODULE=$1

mockgen -destination=../internal/modules/$MODULE/mocks.go -package=${MODULE} -source=../internal/modules/$MODULE/interface.go