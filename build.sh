echo "> building"
# go build -o robocup pkg/*.go
env GOOS=linux GOARCH=arm GOARM=5 go build -o robocup src/*.go
echo "> moving to bin"
rm bin/*
mv robocup bin/
echo "> done"
