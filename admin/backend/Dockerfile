FROM node:lts-buster

EXPOSE 8080

ADD . /src

RUN cd /src && npm install && npm run-script build

WORKDIR /src/dist

CMD ["node", "./main.js"]
