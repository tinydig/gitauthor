#!/bin/sh

VERSION=$1
PLATFORM=$2
URL=https://github.com/tinydig/gitauthor/releases/download/${VERSION}/gitauthor-${PLATFORM}

curl -L -o gitauthor ${URL}
mv gitauthor /usr/local/bin

chmod 755 /usr/local/bin/gitauthor
