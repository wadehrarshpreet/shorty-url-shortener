## FrontEnd Build
FROM node:10.19.0-alpine3.9 as web

WORKDIR /web

COPY web .

RUN yarn
RUN yarn build

# generate server build
FROM golang:1.15.0-alpine as builder

#RUN apt-get install make

WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache make
RUN make swagger

RUN go get -d -v ./...
RUN go install -v ./...


RUN ls /go/bin

EXPOSE 1234

## Final Build
FROM alpine:3.7
# add ca-certificates in case you need them
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# set working directory
WORKDIR /app
# copy the binary from builder
COPY --from=builder /go/bin/short .
# copy docs
COPY --from=builder /go/src/app/docs ./docs
# copy config files
COPY configs ./configs/
# copy web files
COPY --from=web web/dist ./web/dist
# copy index.html
COPY --from=web web/index.html ./web/index.html
# run the binary
CMD ["/app/short"]
