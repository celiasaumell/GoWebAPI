##
## Build
##

FROM golang:alpine

RUN mkdir -p /build/app/cmd/main
WORKDIR /build/app/cmd/main 

COPY ./app/cmd/main /build/app/cmd/main

COPY ./app /build/app

RUN cd /build/app/cmd/main && go build 

EXPOSE 8080

CMD ["/bin/sh"]

ENTRYPOINT ["/build/app/cmd/main/main"]