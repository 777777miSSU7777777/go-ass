# Front build stage
FROM node:lts-buster AS builder

EXPOSE 3000

ARG REACT_APP_API_URL

ADD . /src

RUN cd /src && npm install && npm run-script build

# Front serve stage
FROM nginx:stable as runtime

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /src/build .

CMD ["nginx", "-g", "daemon off;"]