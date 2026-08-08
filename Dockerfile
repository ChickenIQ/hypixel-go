FROM golang:1.26.4-alpine AS build

RUN apk add --no-cache ca-certificates
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/hypixel-go ./cmd/
RUN mkdir -p /out/data

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /out/hypixel-go /hypixel-go

USER 65532:65532
WORKDIR /data

ENTRYPOINT ["/hypixel-go"]