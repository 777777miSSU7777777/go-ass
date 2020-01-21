import React from 'react';
import { isUndefined, isEmpty } from 'lodash-es';
import '../../styles/audio-player/AudioPlayer.css';
import Hls from 'hls.js';
import autoBind from 'react-autobind';

var Config = require('Config');


class AudioPlayer extends React.Component {
    constructor(props){
        super(props);
        autoBind(this);
        this.playerRef = React.createRef();
        this.hls = null;
    }

    componentDidMount(){
        const { dispatchSetPlayer } = this.props;

        dispatchSetPlayer(this.playerRef.current);
    }

    shouldComponentUpdate(prevProps, prevState){
        if (this.props.playingId !== prevProps.playingId){
            return true;
        }
        return false;
    }

    componentDidUpdate(prevProps, prevState) {
        document.title = this.props.track.author + " - " + this.props.track.title;
        this._initPlayer();
    }

    componentWillUnmount(){
        if (this.hls){
            this.hls.destroy();
        }
    }

    _initPlayer(){
        if (this.hls) {
            this.hls.destroy();
        }
        
        this.hls = new Hls();
        this.hls.attachMedia(this.playerRef.current);
        this.hls.on(Hls.Events.MEDIA_ATTACHED, () => {
            this.hls.loadSource(`${Config.serverUrl}/media/${this.props.playingId}/stream/`);
            // this.hls.on(Hls.Events.MANIFEST_PARSED, () => {
            // });
            this.playerRef.current.play();
        });
    }

    onPlay() {
        const { dispatchResume } = this.props;

        dispatchResume();
    }

    onPause() {
        const { dispatchPause } = this.props;

        dispatchPause();
    }

    onEnded() {
        const { dispatchNextTrack } = this.props;

        dispatchNextTrack();
        this.props.player.play();
    }

    prevTrack() {
        const { dispatchPrevTrack } = this.props;
        
        dispatchPrevTrack();
        this.props.player.play();
    }

    nextTrack() {
        const { dispatchNextTrack } = this.props;

        dispatchNextTrack();
        this.props.player.play();
    }

    render(){
        let author, title;
        if (!isUndefined(this.props.track)){
            author = this.props.track.author;
            title = this.props.track.title;
        } else {
            author = "Author";
            title = "Title"
        }

        return(
            <div id="audio-player">
                <div id="audio-player-header">{author} - {title}</div>
                <div id="audio-player-controls">
                    <div id="prev-audio-button" onClick={this.prevTrack}></div>
                    <div id="next-audio-button" onClick={this.nextTrack}></div>
                    <audio 
                    controls 
                    id="player" 
                    ref={this.playerRef} 
                    onPlay={this.onPlay}
                    onPause={this.onPause}
                    onEnded={this.onEnded}
                    />
                </div>
            </div>
        )
    }
}

export default AudioPlayer;