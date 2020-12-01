export const SIGN_IN_IN_PROGRESS = 'SIGN_IN_IN_PROGRESS';
export const SIGN_IN_SUCCESS = 'SIGN_IN_SUCCESS';
export const SIGN_IN_FAILURE = 'SIGN_IN_FAILURE';

export const REFRESH_TOKEN_IN_PROGRESS = 'REFRESH_TOKEN_IN_PROGRESS';
export const REFRESH_TOKEN_SUCCESS = 'REFRESH_TOKEN_SUCCESS';
export const REFRESH_TOKEN_FAILURE = 'REFRESH_TOKEN_FAILURE';

export const SIGN_OUT_IN_PROGRESS = 'SIGN_OUT_IN_PROGRESS';
export const SIGN_OUT_SUCCESS = 'SIGN_OUT_SUCCESS';
export const SIGN_OUT_FAILURE = 'SIGN_OUT_FAILURE';

interface SignInInProgress {
    type: 'SIGN_IN_IN_PROGRESS';
}

interface SignInSuccess {
    type: 'SIGN_IN_SUCCESS';
    acessToken: string;
    refreshToken: string;
}

interface SignInFailure {
    type: 'SIGN_IN_FAILURE';
}

interface RefreshTokenInProgress {
    type: 'REFRESH_TOKEN_IN_PROGRESS';
}

interface RefreshTokenSucces {
    type: 'REFRESH_TOKEN_SUCCESS';
    accessToken: string;
    refreshToken: string;
}

interface RefreshTokenFailure {
    type: 'REFRESH_TOKEN_FAILURE';
}

interface SignOutInProgress {
    type: 'SIGN_OUT_IN_PROGRESS';
}

interface SignOutSuccess {
    type: 'SIGN_OUT_SUCCESS';
}

interface SignOutFailure {
    type: 'SIGN_OUT_FAILURE';
}

export type AuthAction =
    | SignInInProgress
    | SignInSuccess
    | SignInFailure
    | RefreshTokenInProgress
    | RefreshTokenSucces
    | RefreshTokenFailure
    | SignOutInProgress
    | SignOutSuccess
    | SignOutFailure;