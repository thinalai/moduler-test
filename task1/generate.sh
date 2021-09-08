#!/bin/bash

title=$1
file="randoms/$title.md"
file_path="content/$file"

function cmd_check() {
    which $1 &> /dev/null
    if [ $? -ne 0 ]; then
        echo "$1 not found, please install it first" > /dev/stderr
        exit 1
    fi
}

cmd_check hugo
cmd_check fortune

hugo new "$file"

fortune -s | sed -e 's/^/> /' >> "$file_path"
sed -i '' -e 's/^draft:.*/draft: false/' "$file_path"
