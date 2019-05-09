#!/bin/sh
while true; do sleep $INTERVAL; date; wget -q -S --timeout $TIMEOUT $ENDPOINT -O /dev/null; done
