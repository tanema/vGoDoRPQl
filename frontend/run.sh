#!/bin/bash

if [ "${NODE_ENV}" = production ]; then
  yarn install -g http-server && yarn run build && cd build && hs -p 3000
else
  yarn run start
fi
