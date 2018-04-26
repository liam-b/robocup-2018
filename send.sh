# set -e
# rm pkg/*.go
# cp src/**/*.go pkg/
# cp src/*.go pkg/
echo "> sending executable to $1"
scp -q bin/robocup $1:/home/robot/robocup-2018/bin/
echo "> sent"

# scp -r pkg/ $1:/home/robot/robocup-2018/