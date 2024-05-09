A.create enveroment

1. install docker on ec2

install sudo apt update

sudo apt install apt-transport-https ca-certificates curl software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update

sudo apt update

apt-cache policy docker-ce

sudo apt install docker-ce

sudo systemctl status docker

2. install postgres on docker

docker pull postgres

docker run --name some-postgres -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres 

3. install redis

docker pull redis

docker run --name some-redis  -d -p 6379:6379 -e REDIS_PASSWORD=mypas  redis sh -c 'exec redis-server --requirepass "$REDIS_PASSWORD"'

1. backend
   
FrameWork: echo
   
ORM: gorm
   
Languagle: golang
   
Database: postgresql ,redis , dynamoDB.
 
Deploy: aws

2. frontend

FrameWork: reactjs

Languagle: typescript

Deploy: aws

3. mobile

FrameWork: react native
   
Languagle: typescript
   
Deploy: ch play

![image](https://github.com/PTH-IT/karaoke/assets/56516439/7a5c4ae5-8ac2-4e0c-a9e9-ffd2a1495380)
