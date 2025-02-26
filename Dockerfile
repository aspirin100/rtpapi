FROM golang:1.23.3-alpine3.20 AS build

RUN apk --no-cache add make git

COPY . /go/src
WORKDIR /go/src

RUN make build

FROM alpine:3.21.3

COPY --from=build /go/src/bin/rtpapi-server /usr/local/bin/rtpapi-server
COPY --from=build /go/src/main /usr/local/bin/test

EXPOSE 64333

ENTRYPOINT ["rtpapi-server"]