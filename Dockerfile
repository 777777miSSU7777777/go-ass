# Backend build stage
FROM golang:1.15 AS back_builder

ADD . /src

RUN cd /src && go build

# Front build stage
# FROM node:12.13.0-buster AS front_builder

# COPY --from=back_builder /src/frontend /src

# RUN cd /src && npm install && npm run-script build

# Final stage
FROM ubuntu:20.04 AS runtime

EXPOSE 8080

WORKDIR /app

RUN apt update
RUN apt install -y ffmpeg

# COPY --from=front_builder /src/dist /app/frontend

COPY --from=back_builder /src/go-ass /app/

CMD ["sh", "-c", "./go-ass -connection_string=$CONNECTION_STRING -storage_location=$STORAGE_LOCATION"]