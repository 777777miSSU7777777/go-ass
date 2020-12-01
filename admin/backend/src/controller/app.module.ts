import { JWTMiddleware } from '@middleware';
import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { AuthModule } from './auth/auth.module';
import { TableDataModule } from './table-data/table-data.module';

@Module({
  imports: [TableDataModule, AuthModule],
  controllers: [],
  providers: [],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer
      .apply(JWTMiddleware)
      .exclude('/auth/(.*)')
      .forRoutes('*');
  }
}
