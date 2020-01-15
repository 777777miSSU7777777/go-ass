import React from 'react';
import { isUndefined, isEmpty } from 'lodash-es';
import '../../styles/audio-player/AudioPlayer.css';
import Hls from 'hls.js';
var Config = require('Config');

class AudioPlayer extends React.Component {
    constructor(props){
        super(props);
        this.playerRef = React.createRef();
        this.hls = null;
    }

    componentDidMount(){
        this.playerRef.current.onplay = () => {
            this.props.setPlaying(true);
        }
        this.playerRef.current.onpause = () => {
            this.props.setPlaying(false);
        }
        this.playerRef.current.onended = () => {
            this.props.nextTrack();
        }

        this.props.setPlayer(this.playerRef.current);
    }

    componentDidUpdate(prevProps, prevState){
        if (!isUndefined(this.props.track) || !isEmpty(this.props.track)){
            document.title = this.props.track.author + " - " + this.props.track.title;
        }

        if (isUndefined(prevProps.track) && !isUndefined(this.props.track)){
            this._initPlayer();
        } else if (!isUndefined(prevProps.track) && !isUndefined(this.props.track)) {
            if (prevProps.track.id != this.props.track.id) {
                this._initPlayer();
            }
        }
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
            this.hls.loadSource(Config.serverUrl + "/media/" + this.props.track.id + "/stream/");
            // this.hls.on(Hls.Events.MANIFEST_PARSED, () => {
            // });
            this.playerRef.current.play();
        });
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
                    <div id="prev-audio-button" onClick={this.props.prevTrack}></div>
                    <div id="next-audio-button" onClick={this.props.nextTrack}></div>
                    <audio controls id="player" ref={this.playerRef}/>
                </div>
            </div>
        )
    }
}

export default AudioPlayer;