import ApiClient from '../../api/ApiClient';
import * as types from './actionTypes';

export const fetchTracks = () => {
    return async(dispatch, getState) => {
        try {
            const tracks = await ApiClient.instance().getAllAudio();
            dispatch({ type: types.FETCH_TRACKS, tracks});
        } catch (error) {
            console.error(error);
        }
    }
}

export const fetchTracksByKey = () => {
    return async(dispatch, getState) => {
        try {
            const key = getState().audioSearchForm.searchKey;
            const tracks = await ApiClient.instance().searchAudioByKey(key);
            dispatch({ type: types.FETCH_TRACKS, tracks})
        } catch (error) {
            console.error(error);
        }
    }
}

export const newTrack = () => {
    return async(dispatch, getState) => {
        try {
            const form = getState().audioForm.form;
            const track = await ApiClient.instance().newAudio(form);
            dispatch({ type: types.NEW_TRACK, track});
        } catch (error) {
            console.error(error);
        }
    }
}

export const deleteTrack = id => {
    return async(dispatch, getState) => {
        try {
            const resp = await ApiClient.instance().deleteAudioById(id);
            if (resp.status == 200) {
                dispatch({ type: types.DELETE_TRACK, id});
            }
        } catch (error) {
            console.error(error);
        }
    }
}