FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/celiasaumell/GoWebAPI/main
RUN cd /build && git clone https://github.com/celiasaumell/GoWebAPI.git

RUN cd /build/GoWebAPI/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/GoWebAPI/main/main" ]