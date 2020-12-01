import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { LocalStorageKeys } from '../../../enums';
import { signOut } from '../../../store/auth/actions';
import { AppState } from '../../../store/root-reducer';
import styles from './header.module.scss';

interface Props {}

const Header = (props: Props) => {
    const { isAuthenticated } = useSelector((state: AppState) => state.auth);
    const history = useHistory();
    const dispatch = useDispatch();

    const goToSignInPage = () => {
        history.push('/auth/signin');
    }
    
    const onSignOut = () => {
        const refreshToken = localStorage.getItem(LocalStorageKeys.refreshToken) || '';
        dispatch(signOut(refreshToken));
    }

    return (
        <div className={styles.header}>
            <div className={styles.brandLogo}>
                Go ASS Admin
            </div>
            {
                isAuthenticated 
                ? <button className={styles.signOutButton} onClick={onSignOut}>Sign Out</button>
                : <button className={styles.signInButton} onClick={goToSignInPage}>Sign In</button>
            }
        </div>
    )
}

export default Header;