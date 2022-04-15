#!/bin/zsh

# This is the bare minimum to run in development. For full list of flags,
# run ./vigilate -help

go build -o goDict cmd/*.go && ./goDict
#-dbuser='yaji' \
#-pusherHost='localhost' \
#-pusherKey='abc123' \
#-pusherSecret='123abc' \
#-pusherApp="1" \
#-pusherPort="4001" \
#-pusherSecure=false