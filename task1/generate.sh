#!/bin/bash

title=$1
file="randoms/$title.md"
file_path="content/$file"

hugo new "$file"

fortune -s | sed -e 's/^/> /' >> "$file_path"
sed -i '' -e 's/^draft:.*/draft: false/' "$file_path"
