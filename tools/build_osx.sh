#!/usr/bin/env bash

### VERSION
VERSION="0.0.9"
BUILD="20210513"
###

export GO_SRC_DIR=$PWD

echo -n "Updates release version in build_osx? [y/n]"
read answer
if [ "$answer" != "${answer#[Nn]}" ] ;then
  exit;
fi

# MAGIC
cd ${GO_SRC_DIR}

# Clean up before building
if [ -d "dist" ]; then
  rm -fR dist
fi
if [ -f wellness-buddy ]; then
  rm wellness-buddy
fi

# Build go application
go build -v

# Clean up after building
if [ -f wellness-buddy ]; then
  mkdir dist
  mv wellness-buddy dist/wellness-buddy
fi

# Package the executable
fyne package -os darwin -appBuild ${BUILD} -appID "com.bounteous.wellness-buddy" -appVersion "0.0.1" -executable "dist/wellness-buddy" -icon "assets/logo.png" -name "Wellness Buddy"

# Clean up after building
#if [ -f wellness-buddy ]; then
#    rm wellness-buddy
#fi

echo "Build complete."