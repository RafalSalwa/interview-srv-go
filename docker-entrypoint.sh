#!/bin/sh
set -e

echo "Container's IP address: `awk 'END{print $1}' /etc/hosts`"
echo 'Hello World!'
echo "$APP_ENV $APP_NAME"
ECHO 
#if [ "$1" = 'server' ]; then
#    sh examples/server/$@ -bi
#else
#    if [ "$1" = 'client' ]; then
#        sh examples/client/$@
#    else
#        exec "$@"
#    fi
#fi