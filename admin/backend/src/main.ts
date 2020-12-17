import { NestFactory } from '@nestjs/core';
import { AppModule } from './controller/app.module';
import * as dotenv from 'dotenv';
import * as Knex from 'knex';
import { Model } from 'objection';

dotenv.config();
const { DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME } = process.env;

console.log('ENV', DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME);

const knex = Knex({
  client: 'postgresql',
  useNullAsDefault: true,
  connection: {
    user: DB_USER,
    password: DB_PASSWORD,
    host: DB_HOST,
    port: +DB_PORT,
    database: DB_NAME,
  },
});

Model.knex(knex);

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.enableCors();
  await app.listen(8080);
}
bootstrap();
