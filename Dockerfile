# Use the official Ubuntu 22.04 image
FROM node:18.17 AS builder
RUN apt-get update \
    && apt-get install -y golang-go \
    && apt-get install -y make
WORKDIR /app
COPY . .
RUN make build

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/bin /app/bin
EXPOSE 3000
CMD ["/app/bin/server/main"]