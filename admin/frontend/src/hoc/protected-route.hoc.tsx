import React, { Component } from 'react';
import { useSelector } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';
import { AppState } from '../store/root-reducer';

interface Props {
    exact?: boolean;
    path: string;
    component: React.FC;
}
export const ProtectedRoute = (props: Props) => {
    const { isAuthenticated } = useSelector((state: AppState) => state.auth);
    const { component, path } = props;
    return (
        isAuthenticated ? <Route path={path} component={component} /> : <Redirect to='/auth/signin' />
    )
}

export default ProtectedRoute;