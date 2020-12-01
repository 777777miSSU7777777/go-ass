import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { signIn } from '../../store/auth/actions';
import { AppState } from '../../store/root-reducer';
import styles from './sign-in-page.module.scss';

interface Props {}

const SignInPage = (props: Props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { isAuthenticated } = useSelector((state: AppState) => state.auth);
    const dispatch = useDispatch();
    const history = useHistory();

    const onEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    }

    const onPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    }

    const onSignIn = async(e: any) => {
        e.preventDefault();
        dispatch(signIn(email, password));
    }

    useEffect(() => {
        if (isAuthenticated) {
            history.push('/admin');
        }
    }, [isAuthenticated]);
    
    return (
        <div className={styles.signInPage}>
            <div className={styles.signInFormWrapper}>
                <form className={styles.signInForm}>
                    {/* <div className={styles.header}>Sign In</div> */}
                    <div className={styles.group}>
                        <label
                            htmlFor="email"
                            className={styles.label}>Email</label>
                        <input
                            type="text"
                            name="email"
                            placeholder="Your email..."
                            className={styles.field}
                            onChange={onEmailChange} value={email}
                        />
                    </div>
                    <div className={styles.group}>
                        <label
                            htmlFor="password"
                            className={styles.label}>Password</label>
                        <input
                            type="password"
                            name="password"
                            placeholder="Your password..."
                            className={styles.field}
                            onChange={onPasswordChange}
                            value={password}
                        />
                    </div>
                    <button
                        className={styles.submit}
                        onClick={onSignIn}
                    >Sign In</button>
                </form>
            </div>
        </div>
    )
}

export default SignInPage;