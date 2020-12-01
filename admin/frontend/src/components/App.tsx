import React from 'react';
import styles from './App.module.scss';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Header from './common/header/header';
import SignInPage from './auth/sign-in-page';
import AdminWorkspace from './admin-workspace/admin-workspace';
import { ProtectedRoute } from '../hoc';

interface Props {}
const App = (props: Props) => {
  return (
    <div className={styles.app}>
      <BrowserRouter>
        <Header />
        <Switch>
          <Route path='/auth/signin' component={SignInPage}/>
          <ProtectedRoute path='/admin' component={AdminWorkspace}/>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
