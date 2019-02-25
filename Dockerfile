FROM golang:alpine as builder
RUN apk add --no-cache git build-base
ENV GOBIN=$GOPATH/bin
COPY . /go
WORKDIR /go/cmd
RUN go get -v && go install

FROM segment/chamber:2 as chamber

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /go/bin/cmd /app/trello-burndown
COPY --from=chamber /chamber /bin/chamber

ENTRYPOINT ["/bin/chamber", "exec", "trello-burndown", "--", "/app/trello-burndown"]