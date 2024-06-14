#!/bin/bash

#cd cmd/gobot

BASEDIR=$(dirname "$0")
cd $BASEDIR/cmd/gobot

go build || echo

mv ./gobot ../../build
