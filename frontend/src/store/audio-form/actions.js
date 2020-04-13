import * as types from './actionTypes';

export const updateAuthor = author => {
    return ({ type: types.UPDATE_AUTHOR, author});
}

export const updateTitle = title => {
    return ({ type: types.UPDATE_TITLE, title});
}

export const updateFile = file => {
    return ({ type: types.UPDATE_FILE, file});
}

export const validateForm = () => {
    return (dispatch, getState) => {
        const errors = {};
        const { author, title, file } = getState().audioForm.form;
        if (author.length == 0) {
            errors.author = "Author is empty";
        } else if (author.length > 50){
            errors.author = "Author length is more than 50 symbols"
        }
    
        if (title.length == 0) {
            errors.title = "Title is empty";
        } else if (title.length > 50) {
            errors.title = "Title length is more than 50 symbols";
        }
    
        if (!file){
            errors.file = "No file";
        } else if (file.type != "audio/mp3") {
            errors.file = "File is not mp3";
        }
    
        dispatch({ type: types.VALIDATE_FORM, errors });
    }
}

export const openForm = () => {
    return ({ type: types.OPEN_FORM });
}

export const closeForm = () => {
    return ({ type: types.CLOSE_FORM });
}