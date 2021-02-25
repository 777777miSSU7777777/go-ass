# Front build stage
FROM node:lts-buster AS builder

ARG REACT_APP_API_URL

ADD . /src

RUN cd /src && npm install && npm run-script build

# Front serve stage
FROM nginx:stable as runtime

EXPOSE 3000

COPY --from=builder /src/build /app

COPY --from=builder /src/config/nginx/nginx.conf /config/nginx.conf

CMD ["nginx", "-g", "daemon off;", "-c", "/config/nginx.conf"]