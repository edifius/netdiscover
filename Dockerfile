#FROM ulexus/go-minimal
#COPY netdiscover /app

# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM ulexus/go-minimal
COPY --from=builder /app/netdiscover /app/
ENTRYPOINT ["/app/netdiscover"]