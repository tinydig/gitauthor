#!/bin/sh

VERSION=$1
PLATFORM=$2
URL=https://github.com/tinydig/gitauthor/releases/download/${VERSION}/gitauthor-${PLATFORM}

curl -o gitauthor ${URL}
chmod 755 gitauthor
mv gitauthor /usr/local/bin
