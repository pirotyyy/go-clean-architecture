FROM golang:1.18-alpine3.15 as builder
ENV APPDIR /go/app
COPY go.mod go.sum $APPDIR/
WORKDIR $APPDIR
RUN go mod download
WORKDIR /
COPY ./ $APPDIR/
WORKDIR $APPDIR
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ca-server main.go

FROM gcr.io/distroless/static
COPY --from=builder /go/app/ca-server ./app/
CMD ["/app/ca-server"]
