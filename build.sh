rm pkg/*.go
cp src/**/*.go pkg/
cp src/*.go pkg/

go build -o robocup pkg/*.go
rm bin/*
mv robocup bin/
