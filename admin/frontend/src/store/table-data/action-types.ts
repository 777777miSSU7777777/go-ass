import { DataTables } from "../../enums/data-tables";
import { TableData } from "../../model/db";

export const FETCH_TABLE_DATA_IN_PROGRESS = 'FETCH_TABLE_DATA_IN_PROGRESS';
export const FETCH_TABLE_DATA_SUCCESS = 'FETCH_TABLE_DATA_SUCCESS';
export const FETCH_TABLE_DATA_FAILURE = 'FETCH_TABLE_DATA_FAILURE';

export const SAVE_TABLE_DATA_IN_PROGRESS = 'SAVE_TABLE_DATA_IN_PROGRESS';
export const SAVE_TABLE_DATA_SUCCESS = 'SAVE_TABLE_DATA_SUCCESS';
export const SAVE_TABLE_DATA_FAILURE = 'SAVE_TABLE_DATA_FAILURE';

export const SELECT_DATA_TABLE = 'SELECT_DATA_TABLE';

interface FetchTableDataInProgress {
    type: 'FETCH_TABLE_DATA_IN_PROGRESS';
}

interface FetchTableDataSuccess {
    type: 'FETCH_TABLE_DATA_SUCCESS';
    data: TableData[];
}

interface FetchTableDataFailure {
    type: 'FETCH_TABLE_DATA_FAILURE';
}

interface SaveTableDataInProgress {
    type: 'SAVE_TABLE_DATA_IN_PROGRESS';
}

interface SaveTableDataSuccess {
    type: 'SAVE_TABLE_DATA_SUCCESS';
    data: TableData[];
}

interface SaveTableDataFailure {
    type: 'SAVE_TABLE_DATA_FAILURE';
}

interface SelectDataTable {
    type: 'SELECT_DATA_TABLE';
    table: DataTables | null;
}

export type TableDataAction =
    | FetchTableDataInProgress
    | FetchTableDataSuccess
    | FetchTableDataFailure
    | SaveTableDataInProgress
    | SaveTableDataSuccess
    | SaveTableDataFailure
    | SelectDataTable;