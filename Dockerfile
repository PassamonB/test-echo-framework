FROM golang:1.19 as GO_BUILD

RUN mkdir -p /home/app
WORKDIR /home/app
COPY . /home/app

RUN CGO_ENABLED=0 make build

ENTRYPOINT [ "/home/app/test-echo-framework" ]
