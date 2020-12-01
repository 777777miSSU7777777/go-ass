import React from 'react';
import { Redirect, Route, useRouteMatch } from 'react-router-dom';
import AdminHeader from './admin-header/admin-header';
import styles from './admin-workspace.module.scss';
import TableDataEditor from './table-data-editor/table-data-editor';

interface Props {}

const AdminWorkspace = (props: Props) => {
    const { url } = useRouteMatch();
    return (
        <div className={styles.adminWorkspace}>
            <AdminHeader />
            <div className={styles.workspaceWrapper}>
                <Route path={`${url}/table-editor`} component={TableDataEditor}/>
                <Redirect from={url} to={`${url}/table-editor`} />
            </div>
        </div>
    )
}

export default AdminWorkspace;