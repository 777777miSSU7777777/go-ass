import * as types from './actionTypes';

const initialState = {
    searchKey: ""
}

export default function reduce(state = initialState, action = {}) {
    switch(action.type) {
        case types.UPDATE_SEARCH_KEY:
            return { ...state, searchKey: action.searchKey };
        default:
            return state;
    }
}