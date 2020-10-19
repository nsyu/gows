## Docker nginx Setting

```
PS C:\dev\github\gows> docker pull nginx
PS C:\dev\github\gows> docker ps
PS C:\dev\github\gows> docker ps -a
PS C:\dev\github\gows> docker images
PS C:\dev\github\gows> docker run -dp 80:80 nginx
PS C:\dev\github\gows> docker ps
PS C:\dev\github\gows> docker exec -it d2824a1d3b41 /bin/bash
```

```
root@d2824a1d3b41:/# apt-get update
root@d2824a1d3b41:/# apt-get install vim
root@d2824a1d3b41:/# mkdir /home/html
root@d2824a1d3b41:/home/html# vim index.html
```

```
# docker run --name nginx-test -v /home/html:/usr/share/nginx/html:ro -d -p 80:80 nginx
```

## Docker image container Delete

```
PS C:\dev\github\gows> docker stop $(docker ps -a -q)
PS C:\dev\github\gows> docker rm $(docker ps -a -q)
PS C:\dev\github\gows> docker rmi $(docker images -a -q)
PS C:\dev\github\gows> docker images
PS C:\dev\github\gows> docker ps -a
```

## Docker Node Start

```
PS C:\dev\github\gows> docker run -it node
PS C:\dev\github\gows> docker exec -it pensive_hofstadter
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
