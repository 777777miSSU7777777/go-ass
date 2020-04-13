import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as audioPlayerActions from '../store/audio-player/actions';
import AudioPlayer from '../components/audio-player/AudioPlayer.jsx';


const mapStateToProps = state => ({
    player: state.audioPlayer.player,
    isPlaying: state.audioPlayer.isPlaying,
    playingId: state.audioPlayer.playingId,
    track: state.tracks.tracks.find(v => v.id == state.audioPlayer.playingId),
});

const mapDispatchToProps = dispatch => ({
    dispatchSetPlayer: bindActionCreators(audioPlayerActions.setPlayer, dispatch),
    dispatchResume: bindActionCreators(audioPlayerActions.resume, dispatch),
    dispatchPause: bindActionCreators(audioPlayerActions.pause, dispatch),
    dispatchPrevTrack: bindActionCreators(audioPlayerActions.prevTrack, dispatch),
    dispatchNextTrack: bindActionCreators(audioPlayerActions.nextTrack, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(AudioPlayer);