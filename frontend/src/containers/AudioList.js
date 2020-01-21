import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as tracksActions from '../store/tracks/actions';
import * as audioPlayerActions from '../store/audio-player/actions';
import AudioList from '../components/audio-list/AudioList.jsx';


const mapStateToProps = state => ({
    tracks: state.tracks.tracks,
    player: state.audioPlayer.player,
    isPlaying: state.audioPlayer.isPlaying,
    playingId: state.audioPlayer.playingId
});

const mapDispatchToProps = dispatch => ({
    dispatchFetchTracks: bindActionCreators(tracksActions.fetchTracks, dispatch),
    dispatchDeleteTrack: bindActionCreators(tracksActions.deleteTrack, dispatch),
    dispatchPlay: bindActionCreators(audioPlayerActions.play, dispatch),
    dispatchResume: bindActionCreators(audioPlayerActions.resume, dispatch),
    dispatchPause: bindActionCreators(audioPlayerActions.pause, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(AudioList);