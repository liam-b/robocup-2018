# set -e
echo "> clearing package folder"
rm pkg/*.go
echo "> packing files"
cp src/**/*.go pkg/
cp src/*.go pkg/
echo "> removing mock file"
rm pkg/mock.go

echo "> building"
# go build -o robocup pkg/*.go
env GOOS=linux GOARCH=arm GOARM=5 go build -o robocup pkg/*.go
echo "> moving to bin"
rm bin/*
mv robocup bin/
echo "> done"
