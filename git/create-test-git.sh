#!/bin/sh

echo "mkdir test"
mkdir test || exit 1
echo "cd to test"
cd test
git init

echo "file1" > file1
git add file1
git commit -m "feat: file1"
git tag 0.0.1

echo "file2" > file2
git add file2
git commit -m "fix: file2"



