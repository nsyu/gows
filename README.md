## Docker nginx Setting

```
docker pull nginx
docker ps
docker ps -a
docker images
docker run -dp 80:80 nginx
docker ps
docker exec -it d2824a1d3b41 /bin/bash
```

```
apt-get update
apt-get install vim
mkdir /home/html
vim index.html
```

```
docker run --name nginx-test -v /home/html:/usr/share/nginx/html:ro -d -p 80:80 nginx
```

## Docker image container Delete

```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -a -q)
docker images
docker ps -a
```

## Docker Node Start

```
docker run -it node
docker exec -it pensive_hofstadter
```

## Docker apache Build

```
docker build -t mybuild:0.0 ./
docker run -d -P -p 80:80 --name myserver mybuild:0.0
docker exec -it myserver bash
http://localhost/
```

```
FROM ubuntu:14.04
LABEL "purpose"="practice"
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install git -y
RUN apt-get install apache2 -y
ADD test.html /var/www/html
WORKDIR /var/www/html
RUN git clone https://github.com/nsyu/gows.git
RUN ["/bin/bash", "-c", "echo hello >> test2.html"]
EXPOSE 80
CMD apachectl -DFOREGROUND
```
