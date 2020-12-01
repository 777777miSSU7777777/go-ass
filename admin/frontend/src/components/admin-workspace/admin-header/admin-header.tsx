import React from 'react';
import { Link, useRouteMatch } from 'react-router-dom';
import styles from './admin-header.module.scss';

interface Props {}

const AdminHeader = (props: Props) => {
    const { url } = useRouteMatch();

    return (
     <div className={styles.adminHeader}>
        <Link to={`${url}/table-editor`} className={styles.link}>Table Data Editor</Link>
     </div>
 )   
}

export default AdminHeader;