FROM ubuntu:14.04
LABEL "purpose"="practice"
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install git -y
RUN apt-get install apache2 -y
RUN apt-get install golang -y
RUN apt-get install vim -y
WORKDIR /var/www/html
RUN git clone https://github.com/nsyu/gows.git
EXPOSE 80 5000
CMD apachectl -DFOREGROUND