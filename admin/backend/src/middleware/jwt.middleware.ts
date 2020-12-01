import { HttpStatus, Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response } from 'express';
import * as jwt from 'jsonwebtoken';
import { SECRET_KEY, UserRoles } from '@helper';
import { JWTErrorsNames } from '@app/enums';

@Injectable()
export class JWTMiddleware implements NestMiddleware {
    use(req: Request, res: Response, next: Function) {
        try {
            const accessToken: string = req.headers.authorization;

            const { role } = jwt.verify(accessToken, SECRET_KEY);

            if (role !== UserRoles.admin) {
                throw new Error('User is not admin');
            }
            
            next();
        } catch(e) {
            let errStatus: number;
            switch(e.name) {
                case JWTErrorsNames.jsonWebTokenError:
                    errStatus = HttpStatus.UNAUTHORIZED;
                    break;
                case JWTErrorsNames.tokenExpiredError:
                    errStatus = HttpStatus.FORBIDDEN;
                    break;
                default:
                    errStatus = HttpStatus.BAD_REQUEST;
            }

            res.status(errStatus).json({
                'statusCode': errStatus,
                'ok': false,
                'data': null,
                'error': e,
            });

            console.error(`JWT Auth Error: ${e}`);
        }
    }
}