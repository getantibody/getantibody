#!/bin/bash
set -ex

# build and bush to docker
docker build -t "caarlos0/getantibody:$CIRCLE_BUILD_NUM" .
docker push "caarlos0/getantibody:$CIRCLE_BUILD_NUM"
docker tag "caarlos0/getantibody:$CIRCLE_BUILD_NUM" caarlos0/getantibody:latest
docker push caarlos0/getantibody:latest

# push to beanstalk
cd .deploy
sed -i'' -e "s;%BUILD_NUM%;$CIRCLE_BUILD_NUM;g" Dockerrun.aws.json
eb init -r us-east-1 getantibody
eb deploy -l $CIRCLE_BUILD_NUM
