FROM golang:1.24-alpine AS build

WORKDIR /app

COPY . .
RUN go mod tidy \
&& GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -C ./cmd/auction -o auction

FROM scratch
WORKDIR /app
COPY --from=build /app/cmd/auction/auction auction
COPY --from=build /app/cmd/auction/.env cmd/auction/.env
COPY --from=build /etc/ssl/certs/ /etc/ssl/certs/
EXPOSE 8080
ENTRYPOINT ["./auction"]