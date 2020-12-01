import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import { LocalStorageKeys, DataActions } from '../enums';
import { TableData } from '../model/db';
import { AuthAPI } from './auth.api';
import store from '../store/store';
import { REFRESH_TOKEN_FAILURE, REFRESH_TOKEN_IN_PROGRESS, REFRESH_TOKEN_SUCCESS } from '../store/auth/action-types';

const tableDataInstance = axios.create({
    baseURL: `${process.env.REACT_APP_API_URL}/data`,
    headers: {
        'Content-Type': 'application/json',
    }
});

const TOKEN_EXPIRED_STATUS_CODE = 403;

tableDataInstance.interceptors.request.use((req: AxiosRequestConfig) => {
    req.headers['Authorization'] = localStorage.getItem(LocalStorageKeys.accessToken);
    return req;
});

tableDataInstance.interceptors.response.use((res: AxiosResponse) => res, (err: AxiosError) => {
    const status = err.response ? err.response.status : null

    if (status) {
        if (status === TOKEN_EXPIRED_STATUS_CODE) {
            const token: string = localStorage.getItem(LocalStorageKeys.refreshToken) || '';
            store.dispatch({
                type: REFRESH_TOKEN_IN_PROGRESS,
            });

            return AuthAPI.refreshToken(token)
                .then(data => {
                    const { accessToken, refreshToken } = data;

                    localStorage.setItem(LocalStorageKeys.accessToken, accessToken);
                    localStorage.setItem(LocalStorageKeys.refreshToken, refreshToken);
                    store.dispatch({
                        type: REFRESH_TOKEN_SUCCESS,
                        accessToken,
                        refreshToken,
                    });
                    
                    return tableDataInstance(err.config);
                })
                .catch(err => {
                    store.dispatch({
                        type: REFRESH_TOKEN_FAILURE,
                    });
                    localStorage.setItem(LocalStorageKeys.accessToken, '');
                    localStorage.setItem(LocalStorageKeys.refreshToken, '');

                    return Promise.reject(err);
                });
        }
    }

    return Promise.reject(err);
});

export const TableDataAPI = {
    getData(dataRoute: string) {
        return tableDataInstance.post(`${dataRoute}?action=${DataActions.get}`).then(res => res.data);
    },

    newData(dataRoute: string, data: TableData[]) {
        return tableDataInstance.post(`${dataRoute}?action=${DataActions.create}`, data).then(res => res.data);
    },

    updateData(dataRoute: string, data: TableData[]) {
        return tableDataInstance.post(`${dataRoute}?action=${DataActions.update}`, data).then(res => res.data);
    },

    deleteData(dataRoute: string, data: TableData[]) {
        return tableDataInstance.post(`${dataRoute}?action=${DataActions.delete}`, data).then(res => res.data);
    }
}