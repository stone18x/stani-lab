FROM golang:1.15.7-alpine3.13 AS builder

WORKDIR /dockersrc/
COPY main.go /dockersrc/
COPY go.* /dockersrc/
COPY app.go /dockersrc/
COPY ./model /dockersrc/model
COPY ./services /dockersrc/services

RUN ls

RUN CGO_ENABLED=0 go build -o /bin/image

FROM scratch
COPY --from=builder /bin/image /bin/app
ENTRYPOINT ["/bin/app"]