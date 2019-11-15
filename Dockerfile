# Backend build stage
FROM golang:1.13 AS back_builder

ENV GO111MODULE=on

ADD . /src

RUN cd /src && go build

# Front build stage
FROM node:12.13.0-buster AS front_builder

COPY --from=back_builder /src/frontend /src

RUN cd /src && npm install && npm run-script build

# Final stage
FROM ubuntu:18.04 AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=back_builder /src/ffmpeg /usr/bin/

COPY --from=front_builder /src/dist /app/frontend

COPY --from=back_builder /src/go-ass /app/

CMD ["sh", "-c", "./go-ass -connection_string=$CONNECTION_STRING -storage_location=$STORAGE_LOCATION -api_only=$API_ONLY"]