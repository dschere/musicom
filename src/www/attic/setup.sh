#!/usr/bin/env bash

export VEXFLOW_TAG=4.1.0

set -e

# populate lib directory and add our modification for events
git clone https://github.com/0xfe/vexflow.git 
cd vexflow 
git fetch --tags && git checkout $VEXFLOW_TAG -b latest 
cp ../factory-$VEXFLOW_TAG-modified.ts src/factory.ts && \
npm install vexflow 
npm link
cd ..

git clone https://github.com/0xfe/vextab.git
cd vextab
#MAYBE
# may need to fork this repo to sync version with vexflow
npm link vexflow
npm install



	


