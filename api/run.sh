#!/bin/bash

if [ ${APP_ENV} = production ]; then
  ./api;
else
  go get github.com/pilu/fresh && fresh
fi
