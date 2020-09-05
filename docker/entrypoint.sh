#!/bin/bash
MODD_DIR=modd
cd docker
if [ ! -d "$MODD_DIR" ]; then
  git clone https://github.com/cortesi/modd
  git checkout 45e41bb091af8d3c7bfdc523851a9fc993df8a80
  cd modd
  go get mvdan.cc/sh@8aeb0734cd0f
  cd -
fi
cd modd
go install ./cmd/modd
cd -
cd ..
modd -f ./docker/modd.conf
