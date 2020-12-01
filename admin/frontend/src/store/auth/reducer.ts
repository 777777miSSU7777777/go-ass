import { AuthAction, REFRESH_TOKEN_FAILURE, REFRESH_TOKEN_IN_PROGRESS, REFRESH_TOKEN_SUCCESS, SIGN_IN_IN_PROGRESS, SIGN_IN_SUCCESS, SIGN_OUT_FAILURE, SIGN_OUT_IN_PROGRESS, SIGN_OUT_SUCCESS } from "./action-types";
import { LocalStorageKeys } from "../../enums";
import { getTokenPayload } from "../../helper";

export interface AuthState {
    isAuthenticated: boolean;
    userId: string;
    userRole: string;
    signInInProgress: boolean;
}

const initialState: AuthState = {
    isAuthenticated: !!localStorage.getItem(LocalStorageKeys.refreshToken),
    userId: getTokenPayload(localStorage.getItem(LocalStorageKeys.accessToken)).userId || '',
    userRole: getTokenPayload(localStorage.getItem(LocalStorageKeys.accessToken)).role || '',
    signInInProgress: false,
}

export default (state: AuthState = initialState, action: AuthAction) => {
    switch(action.type) {
        case SIGN_IN_IN_PROGRESS:
            return {
                ...state,
                signInProgress: true,
            };

        case SIGN_IN_SUCCESS:
            return {
                ...state,
                isAuthenticated: true,
                userId: getTokenPayload(localStorage.getItem(LocalStorageKeys.accessToken)).userId,
                userRole: getTokenPayload(localStorage.getItem(LocalStorageKeys.accessToken)).role,
                signInInProgress: false,
            };

        case SIGN_OUT_FAILURE:
            return {
                ...state,
                signInInProgress: false,
            };

        case REFRESH_TOKEN_IN_PROGRESS:
            return state;

        case REFRESH_TOKEN_SUCCESS:
            return state;

        case REFRESH_TOKEN_FAILURE:
            return {
                ...state,
                isAuthenticated: false,
                userId: '',
                userRole: '',
            };

        case SIGN_OUT_IN_PROGRESS:
            return state;

        case SIGN_OUT_SUCCESS:
            return {
                ...state,
                isAuthenticated: false,
                userId: '',
                userRole: '',
            };

        case SIGN_OUT_FAILURE:
            return state;

        default:
            return state;
    }
}