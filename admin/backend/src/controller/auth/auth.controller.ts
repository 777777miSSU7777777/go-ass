import { Body, Controller, Delete, Get, HttpStatus, Post, Put, Res } from '@nestjs/common';
import { Response } from 'express';
import AuthService from '@app/service/auth/auth.service';

@Controller('/auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('/signin')
  async signIn(@Body() body: { email: string, password: string}, @Res() res: Response) {
    try {
        const tokens: { accessToken: string, refreshToken: string } = await this.authService.signIn(body.email, body.password);

        res.status(HttpStatus.OK).json({
            'statusCode': HttpStatus.OK,
            'ok': true,
            'data': tokens,
            'error': null,
        });
    } catch(e) {
        res.status(HttpStatus.BAD_REQUEST).json({
            'statusCode': HttpStatus.BAD_REQUEST,
            'ok': false,
            'data': null,
            'error': e,
        });

        console.error(`Sign In Error: ${e}`);
    }
  }

  @Post('/refresh-token')
  async refreshToken(@Body() body: { token: string }, @Res() res: Response) {
    try {
        const tokens: { accessToken: string, refreshToken: string } = await this.authService.refreshToken(body.token);

        res.status(HttpStatus.OK).json({
            'statusCode': HttpStatus.OK,
            'ok': true,
            'data': tokens,
            'error': null,
        });
    } catch(e) {
        res.status(HttpStatus.BAD_REQUEST).json({
            'statusCode': HttpStatus.BAD_REQUEST,
            'ok': false,
            'data': null,
            'error': e,
        });

        console.error(`Refresh Token Error: ${e}`);
    }
  }

  @Post('/signout')
  async signOut(@Body() body: { token: string }, @Res() res: Response) {
    try {
        await this.authService.signOut(body.token);

        res.status(HttpStatus.OK).json({
            'statusCode': HttpStatus.OK,
            'ok': true,
            'data': null,
            'error': null,
        });
    } catch(e) {
        res.status(HttpStatus.BAD_REQUEST).json({
            'statusCode': HttpStatus.BAD_REQUEST,
            'ok': false,
            'data': null,
            'error': e,
        });

        console.error(`Sign Out Error: ${e}`);
    }
  }
}

export default AuthController;
