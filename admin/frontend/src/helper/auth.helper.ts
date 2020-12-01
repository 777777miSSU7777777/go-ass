import jwt_decode from 'jwt-decode';

export const getTokenPayload = (token?: string | null): any => {
    if (token) {
        return jwt_decode(token);
    }

    return {};
}