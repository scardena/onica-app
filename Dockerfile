FROM golang

RUN apt-get update && apt-get install vim -y

WORKDIR /var/app

COPY . ./

CMD ["go","run","server.go"]
