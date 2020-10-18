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

```
PS C:\dev\github\gows> docker stop $(docker ps -a -q)
PS C:\dev\github\gows> docker rm $(docker ps -a -q)
PS C:\dev\github\gows> docker rmi $(docker images -a -q)
PS C:\dev\github\gows> docker images
PS C:\dev\github\gows> docker ps -a
```

```
PS C:\dev\github\gows> docker run -it node
PS C:\dev\github\gows> docker exec -it pensive_hofstadter
```

```
docker build -t mybuild:0.0 ./
docker run -d -P --name myserver mybuild:0.0
http://localhost:32768/
```
