#!/bin/sh

buildlibsass() {
  local current_dir=`pwd`
  mkdir -p libs
  cd libsass
  autoreconf --install
  ./configure --prefix=$current_dir/libs
  make && make install
  cd ..
}

git submodule update --init
buildlibsass
