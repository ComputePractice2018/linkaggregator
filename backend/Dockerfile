FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/linkaggregator/backend
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/linkaggregator github.com/ComputePractice2018/linkaggregator/backend/

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/linkaggregator /opt/linkaggregator

EXPOSE 8080

ENV NAME "Dmitry"
ENTRYPOINT /opt/linkaggregator -name ${NAME}