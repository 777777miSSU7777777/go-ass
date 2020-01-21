import * as types from './actionTypes';

export const setPlayer = player => {
    return ({ type: types.SET_PLAYER, player });
}

export const play = id => {
    return ({ type: types.PLAY, id});
}

export const resume = () => {
    return ({ type: types.RESUME });
}

export const pause = () => {
    return ({ type: types.PAUSE });
}

export const prevTrack = () => {
    return (dispatch, getState) => {
        const tracks = getState().tracks.tracks;
        let arrId = tracks.findIndex(v => v.id == getState().audioPlayer.playingId);
        arrId = arrId == 0 ? tracks.length - 1: arrId - 1; 
        dispatch({ type: types.PREV_TRACK, id: tracks[arrId].id})
    }
}

export const nextTrack = () => {
    return (dispatch, getState) => {
        const tracks = getState().tracks.tracks;
        let arrId = tracks.findIndex(v => v.id == getState().audioPlayer.playingId);
        arrId = arrId == tracks.length - 1 ? 0 : arrId + 1; 
        dispatch({ type: types.NEXT_TRACK, id: tracks[arrId].id})
    }
}