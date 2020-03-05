FROM golang:1.13 as builder

WORKDIR /go/src/github.com/dcherman/homeassistant-roborock-monitor
COPY . .

RUN CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static

COPY --from=builder /go/src/github.com/dcherman/homeassistant-roborock-monitor/main /bin/app
ENTRYPOINT [ "/bin/app" ]
