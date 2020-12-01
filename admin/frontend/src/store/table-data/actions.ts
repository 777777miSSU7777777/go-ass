import { Dispatch } from "redux";
import { TableDataAPI } from "../../api";
import { DataTables } from "../../enums/data-tables";
import { TableData } from "../../model/db";
import { FETCH_TABLE_DATA_FAILURE, FETCH_TABLE_DATA_IN_PROGRESS, FETCH_TABLE_DATA_SUCCESS, SAVE_TABLE_DATA_FAILURE, SAVE_TABLE_DATA_IN_PROGRESS, SAVE_TABLE_DATA_SUCCESS, SELECT_DATA_TABLE } from "./action-types";

const fetchTableDataInProgress = () => ({
    type: FETCH_TABLE_DATA_IN_PROGRESS,
});

const fetchTableDataSuccess = (data: TableData[]) => ({
    type: FETCH_TABLE_DATA_SUCCESS,
    data,
});

const fetchTableDataFailure = () => ({
    type: FETCH_TABLE_DATA_FAILURE,
});

const saveTableDataInProgress = () => ({
    type: SAVE_TABLE_DATA_IN_PROGRESS,
});

const saveTableDataSuccess = (data: TableData[]) => ({
    type: SAVE_TABLE_DATA_SUCCESS,
    data,
});

const saveTableDataFailure = () => ({
    type: SAVE_TABLE_DATA_FAILURE,
});

const selectDataTableAction = (table: DataTables | null) => ({
    type: SELECT_DATA_TABLE,
    table
});

export const fetchTableData = (dataRoute: string) => async(dispatch: Dispatch) => {
    dispatch(fetchTableDataInProgress());
    
    try {
        const { data } = await TableDataAPI.getData(dataRoute);

        dispatch(fetchTableDataSuccess(data));
    } catch (e) {
        dispatch(fetchTableDataFailure());

        console.error(`Fetch Table Data Error: ${e}`);
    }
}

export const saveTableData = (dataRoute: string, newData: TableData[], updatedData: TableData[], deletedData: TableData[]) => async(dispatch: Dispatch) => {
    dispatch(saveTableDataInProgress());

    try {
        await Promise.all(
            [
                newData.length > 0 ? TableDataAPI.newData(dataRoute, newData) : null,
                updatedData.length > 0 ? TableDataAPI.updateData(dataRoute, updatedData) : null,
                deletedData.length > 0 ? TableDataAPI.deleteData(dataRoute, deletedData) : null
            ]
        )

        const { data } = await TableDataAPI.getData(dataRoute);

        dispatch(saveTableDataSuccess(data));
    } catch(e) {
        dispatch(saveTableDataFailure());

        console.error(`Save Table Data Error: ${e}`);
    }
}

export const selectDataTable = (table: DataTables | null) => async(dispatch: Dispatch) => {
    dispatch(selectDataTableAction(table));
}