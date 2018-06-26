export GO_BUILD_COUNT=$(($GO_BUILD_COUNT+1))
rm pkg/* 2> /dev/null
cp src/**/*.go pkg
cp src/*.go pkg
echo -e "> \033[0;32mbuilding\033[0;0m \033[0;30m"$GO_BUILD_COUNT"\033[0;0m"
env GOOS=linux GOARCH=arm GOARM=5 go build -o robocup pkg/*
if [[ $? != 0 ]]; then
  echo -e "< \033[0;31mbuild failed\033[0;0m"
else
  echo -e "| \033[0;32mbuild finished\033[0;0m"
  rm bin/* 2> /dev/null
  mv robocup bin/
  echo -e "| \033[0;32msending executable\033[0;0m"
  scp -q bin/robocup $1:/home/robot/src/bin/
  if [[ $? != 0 ]]; then
    echo -e "< \033[0;31msend failed\033[0;0m"
  else
    echo -e "< \033[0;32mdone\033[0;0m"
  fi
fi