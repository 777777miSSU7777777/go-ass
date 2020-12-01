import { TableDataAPI } from "../../api";
import { DataTables } from "../../enums/data-tables";
import { TableData } from "../../model/db";
import { FETCH_TABLE_DATA_FAILURE, FETCH_TABLE_DATA_IN_PROGRESS, FETCH_TABLE_DATA_SUCCESS, SAVE_TABLE_DATA_FAILURE, SAVE_TABLE_DATA_IN_PROGRESS, SAVE_TABLE_DATA_SUCCESS, SELECT_DATA_TABLE, TableDataAction } from "./action-types";

export interface TableDataState {
    selectedTable: DataTables | null;
    data: TableData[];
    isFetchingInProgress: boolean;
    isSavingInProgress: boolean;
}

const initialState: TableDataState = {
    selectedTable: DataTables.track,
    data: [],
    isFetchingInProgress: false,
    isSavingInProgress: false,
}

export default (state: TableDataState = initialState, action: TableDataAction) => {
    switch(action.type) {
        case FETCH_TABLE_DATA_IN_PROGRESS:
            return {
                ...state,
                isFetchingInProgress: true,
            }
        
        case FETCH_TABLE_DATA_SUCCESS:
            return {
                ...state,
                data: action.data,
                isFetchingInProgress: false,
            }

        case FETCH_TABLE_DATA_FAILURE:
            return {
                ...state,
                isFetchingInProgress: false,
            }

        case SAVE_TABLE_DATA_IN_PROGRESS:
            return {
                ...state,
                isSavingInProgress: true,
            }

        case SAVE_TABLE_DATA_SUCCESS:
            return {
                ...state,
                data: action.data,
                isSavingInProgress: false,
            }

        case SAVE_TABLE_DATA_FAILURE:
            return {
                ...state,
                isSavingInProgress: false,
            }

        case SELECT_DATA_TABLE:
            return {
                ...state,
                selectedTable: action.table,
            }

        default:
            return state;
    }
}
