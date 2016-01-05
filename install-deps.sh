#!/bin/sh

buildlibsass() {
  mkdir -p libs
  cd libsass
  autoreconf --install
  ./configure
  make && make install
  cd ..
}

git submodule update --init
buildlibsass
