#!/bin/bash

go build -o bin/$1 days/$1/*

if [ $? -eq 0 ]
then
    bin/$1;
fi
