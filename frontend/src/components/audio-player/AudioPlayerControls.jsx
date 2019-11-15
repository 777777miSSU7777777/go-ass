import React from 'react';
import PrevAudioButton from './PrevAudioButton.jsx';
import NextAudioButton from './NextAudioButton.jsx';
import '../../styles/audio-player/AudioPlayerControls.css';
import Hls from 'hls.js';

class AudioPlayerControls extends React.Component {
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

        this.props.setPlayer(this.playerRef.current);;
    }

    componentDidUpdate(prevProps, prevState){
        if (prevProps.playingId != this.props.playingId && this.props.playingId !== undefined) {
            this.initPlayer()
        }
    }

    initPlayer(){
        if (this.hls){
            this.hls.destroy();
        }

        this.hls = new Hls();
        this.hls.loadSource("/media/" + this.props.playingId + "/stream/");
        this.hls.attachMedia(this.playerRef.current);
        this.hls.on(Hls.Events.MEDIA_ATTACHED, () => {
            // this.hls.on(Hls.Events.MANIFEST_PARSED, () => {
            // })
            this.playerRef.current.play();
        })
    }

    componentWillUnmount(){
        if (this.hls){
            this.hls.destroy();
        }
    }

    
    render(){
        return (
            <div id="audio-player-controls">
                <PrevAudioButton onClick={this.props.prevTrack}/>
                <NextAudioButton onClick={this.props.nextTrack}/>
                <audio controls id="player" ref={this.playerRef}/>
            </div>
        )
    }
}

export default AudioPlayerControls;