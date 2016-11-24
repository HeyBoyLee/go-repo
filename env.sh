#!/bin/sh

BASE_DIR=`cd $(dirname $0);pwd -P`
#HOME_ = /home/mi
echo ${BASE_DIR}
GOPATH=$BASE_DIR
export GOPATH
