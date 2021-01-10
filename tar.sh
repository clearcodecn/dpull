#!/bin/bash

cd dist
for i in ./* ; do
  if [[ -d $i ]] ; then
      if [[ -f "${i}/dpull.exe" ]] ;then
          tar -zcf ${i}.tar.gz $i/dpull.exe
      else
          tar -zcf ${i}.tar.gz $i/dpull
      fi
  fi
done