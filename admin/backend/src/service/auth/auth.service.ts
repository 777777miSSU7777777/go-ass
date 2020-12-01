import { User, UserTokens } from '@model';
import { Injectable } from '@nestjs/common';
import * as bcrypt from 'bcrypt';
import * as jwt from 'jsonwebtoken';
import { TokensDuration, UserRoles, SECRET_KEY } from '@helper';

@Injectable()
export class AuthService {
    async signIn(email: string, password: string): Promise<{ accessToken: string, refreshToken: string }> {
        const user = await User.query().findOne({ email: email });

        if (!user) {
            throw new Error(`User with email '${email}' is not found`);
        }

        if (bcrypt.compareSync(password, user.password)) {
            const accessToken: string = jwt.sign({
                'userId': user.userId,
                'role': user.role,
            }, SECRET_KEY, { expiresIn: TokensDuration.access });

            const refreshToken: string = jwt.sign({
                'userId': user.userId,
                'role': user.role,
            }, SECRET_KEY, { expiresIn: TokensDuration.refresh });

            await UserTokens.query().insert({ userId: user.userId, token: refreshToken });

            return { accessToken, refreshToken };
        } else {
            throw new Error('User password is incorrect');
        }
    }

    async refreshToken(token: string): Promise<{ accessToken: string, refreshToken: string }> {
        jwt.verify(token, SECRET_KEY);
    
        const { userId, role } = jwt.decode(token);

        if (role === UserRoles.admin) {            
            const accessToken: string = jwt.sign({
                'userId': userId,
                'role': role,
            }, SECRET_KEY, { expiresIn: TokensDuration.access });

            const refreshToken: string = jwt.sign({
                'userId': userId,
                'role': role,
            }, SECRET_KEY, { expiresIn: TokensDuration.refresh });

            await UserTokens.query()
                .findOne({ user_id: userId, token: token })
                .update({ userId: userId, token: refreshToken })
                .throwIfNotFound(new Error('User Token is not found'));

            return { accessToken, refreshToken };
        } else {
            throw new Error('User is not admin');
        }
    }

    async signOut(token: string) {
        await UserTokens.query().findOne({ token: token }).delete();
    }
}

export default AuthService;
