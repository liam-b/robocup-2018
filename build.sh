echo "> building"
env GOOS=linux GOARCH=arm GOARM=5 go build -o robocup src/*.go
if [[ $? != 0 ]]; then
  echo "> build failed"
  exit
else
  echo "> moving to bin"
  rm bin/* 2> /dev/null
  mv robocup bin/
  echo "> done"
fi
