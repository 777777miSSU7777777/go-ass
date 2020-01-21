import * as types from './actionTypes';
import { isEmpty } from 'lodash-es/isEmpty';

const initialState = {
    playerRef: null,
    isPlaying: false,
    playingId: 0
}

export default function reduce(state = initialState, action = {}) {
    switch(action.type) {
        case types.SET_PLAYER:
            return { ...state, player: action.player }
        case types.PLAY:
            return { ...state, isPlaying: true, playingId: action.id };
        case types.RESUME:
            return { ...state, isPlaying: true };
        case types.PAUSE:
            return { ...state, isPlaying: false};
        case types.PREV_TRACK:
            return { ...state, playingId: action.id };
        case types.NEXT_TRACK:
            return { ...state, playingId: action.id };
        default:
            return state;
    }
}

export const getPlayingTrack = (state) => {
    return state.tracks.tracks.find( v => v.id == state.audioPlayer.playingId );
}