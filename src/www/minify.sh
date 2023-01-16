#!/usr/bin/env sh 

rm -f app.min.js 2> /dev/null 
uglifyjs app/*.js > app.min.js


