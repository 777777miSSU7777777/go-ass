import * as types from './actionTypes';

const initialState = {
    opened: false,
    form: { author: "", title: "", file: null },
    formErrors: {}
}

export default function reduce(state = initialState, action = {}) {
    switch(action.type) {
        case types.OPEN_FORM:
            return { ...state, opened: true };
        case types.UPDATE_AUTHOR:
            return { ...state,  form: { ...state.form, author: action.author } };
        case types.UPDATE_TITLE:
            return { ...state, form: { ...state.form, title: action.title } };
        case types.UPDATE_FILE:
            return { ...state, form: { ...state.form, file: action.file } };
        case types.VALIDATE_FORM:
            return { ...state, formErrors: action.errors };
        case types.CLOSE_FORM:
            return { ...state, opened: false, form: { author: "", title: "", file: null}, formErrors: {} };
        default:
            return state;
    }
}