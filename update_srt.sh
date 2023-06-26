#!/bin/bash

# Setup location to download new copy of srt into
echo "Creating temporary workspace"
if [ -d "./tmp" ];
then
    rm -rf ./tmp
fi
mkdir -p ./tmp

# Pull down the new headers
echo "Fetching srt repo"
cd tmp
git clone https://github.com/Haivision/srt.git
cd ..
cp ./tmp/srt/srtcore/*.h .

# Remove #include "version.h" from srt.h
echo "Fixing srt.h"
grep -v "#include \"version.h\"" ./srt.h > ./tmp/srt.h
mv ./tmp/srt.h ./srt.h

# Cleanup
echo "Cleaning up"
rm -rf ./tmp
