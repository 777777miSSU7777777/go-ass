import { History } from 'history';
import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';
import history from './browserhistory';
import { AuthState } from './auth/reducer';
import AuthReducer from './auth/reducer';
import { TableDataState } from './table-data/reducer';
import TableDataReducer from './table-data/reducer';

export interface AppState {
    router: History;
    auth: AuthState;
    tableData: TableDataState;
}

const rootReducer = combineReducers({
    router: connectRouter(history),
    auth: AuthReducer,
    tableData: TableDataReducer,
});

export default rootReducer;
