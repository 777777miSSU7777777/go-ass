# Build stage
FROM golang:1.13 AS builder

ENV GO111MODULE=on

ADD . /src

RUN cd /src && go build

# Final stage
FROM ubuntu:18.04 AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/ffmpeg /usr/bin/

COPY --from=builder /src/frontend/. /app/frontend

COPY --from=builder /src/go-ass /app/

CMD ["sh", "-c", "./go-ass -connection_string=$CONNECTION_STRING -storage_location=$STORAGE_LOCATION"]