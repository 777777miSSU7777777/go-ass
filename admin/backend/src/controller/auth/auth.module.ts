import AuthService from '@app/service/auth/auth.service';
import { Module } from '@nestjs/common';
import AuthController from './auth.controller';

@Module({
  imports: [],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
