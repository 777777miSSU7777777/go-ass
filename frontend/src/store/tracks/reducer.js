import * as types from './actionTypes';

const initialState = {
    tracks: []
}

export default function reduce(state = initialState, action = {}) {
    switch(action.type) {
        case types.FETCH_TRACKS:
            return { ...state,  tracks: action.tracks };
        case types.NEW_TRACK:
            return { ...state, tracks: [...state.tracks, action.track] };
        case types.DELETE_TRACK:
            return { ...state, tracks: state.tracks.filter(v => v.id != action.id) };
        default:
            return state;
    }
}