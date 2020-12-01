import { Dispatch } from "redux";
import { AuthAPI } from "../../api";
import { LocalStorageKeys } from "../../enums";
import { REFRESH_TOKEN_FAILURE, REFRESH_TOKEN_IN_PROGRESS, REFRESH_TOKEN_SUCCESS, SIGN_IN_FAILURE, SIGN_IN_IN_PROGRESS, SIGN_IN_SUCCESS, SIGN_OUT_FAILURE, SIGN_OUT_IN_PROGRESS, SIGN_OUT_SUCCESS } from "./action-types";

const signInInProgress = () => ({
    type: SIGN_IN_IN_PROGRESS,
});

const signInSuccess = (accessToken: string, refreshToken: string) => ({
    type: SIGN_IN_SUCCESS,
    accessToken,
    refreshToken,
});

const signInFailure = () => ({
    type: SIGN_IN_FAILURE,
})

const refreshTokenInProgress = () => ({
    type: REFRESH_TOKEN_IN_PROGRESS,
});

const refreshTokenSuccess = (accessToken: string, refreshToken: string) => ({
    type: REFRESH_TOKEN_SUCCESS,
    accessToken,
    refreshToken,
});

const refreshTokenFailure = () => ({
    type: REFRESH_TOKEN_FAILURE,
});

const signOutInProgress = () => ({
    type: SIGN_OUT_IN_PROGRESS,
});

const signOutSuccess = () => ({
    type: SIGN_OUT_SUCCESS,
});

const signOutFailure = () => ({
    type: SIGN_OUT_FAILURE,
});

export const signIn = (email: string, password: string) => async(dispatch: Dispatch) => {
    dispatch(signInInProgress());

    try {
        const { accessToken, refreshToken } = await AuthAPI.signIn(email, password);
        
        localStorage.setItem(LocalStorageKeys.accessToken, accessToken);
        localStorage.setItem(LocalStorageKeys.refreshToken, refreshToken);
        dispatch(signInSuccess(accessToken, refreshToken));
    } catch(e) {
        dispatch(signInFailure());

        console.error(`Sign In Error: ${e}`);
    }
}

export const refreshToken = (token: string) => async(dispatch: Dispatch) => {
    dispatch(refreshTokenInProgress());

    try {
        const { accessToken, refreshToken } = await AuthAPI.refreshToken(token);

        localStorage.setItem(LocalStorageKeys.accessToken, accessToken);
        localStorage.setItem(LocalStorageKeys.refreshToken, refreshToken);
        dispatch(refreshTokenSuccess(accessToken, refreshToken));
    } catch(e) {
        localStorage.setItem(LocalStorageKeys.accessToken, '');
        localStorage.setItem(LocalStorageKeys.refreshToken, '');
        dispatch(refreshTokenFailure());

        console.error(`Refresh Token Error: ${e}`);
    }
}

export const signOut = (token: string) => async(dispatch: Dispatch) => {
    dispatch(signOutInProgress());

    try {
        const _ = await AuthAPI.signOut(token);

        localStorage.setItem(LocalStorageKeys.accessToken, '');
        localStorage.setItem(LocalStorageKeys.refreshToken, '');
        dispatch(signOutSuccess());
    } catch(e) {
        localStorage.setItem(LocalStorageKeys.accessToken, '');
        localStorage.setItem(LocalStorageKeys.refreshToken, '');
        dispatch(signOutFailure());

        console.error(`Sign Out Erorr: ${e}`);
    }
}