import * as types from './actionTypes';

export const updateSearchKey = searchKey => {
    return ({ type: types.UPDATE_SEARCH_KEY, searchKey});
}
