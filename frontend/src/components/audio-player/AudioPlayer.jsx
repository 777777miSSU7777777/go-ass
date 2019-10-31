import React from 'react';
import AudioPlayerHeader from './AudioPlayerHeader.jsx';
import AudioPlayerControls from './AudioPlayerControls.jsx';
import '../../styles/audio-player/AudioPlayer.css';

class AudioPlayer extends React.Component {
    constructor(props){
        super(props);
        this.state = {author: "Author", title: "Title", currentTime: 0};
    }

    componentDidUpdate(){
        if (this.props.playingId) {
            fetch("/api/audio/" + this.props.playingId, {method: "GET"})
            .then(resp => resp.json())
            .then(data => {
                this.setState({author: data.author, title: data.title})
                document.title = this.state.author + " - " + this.state.title;
            });
        };
    }

    render(){
        return(
            <div id="audio-player">
                <AudioPlayerHeader 
                author={this.state.author} 
                title={this.state.title} />
                <AudioPlayerControls 
                setPlayer={this.props.setPlayer}
                playingId={this.props.playingId}
                isPlaying={this.props.isPlaying}
                setPlaying={this.props.setPlaying}
                prevTrack={this.props.prevTrack}
                nextTrack={this.props.nextTrack}
                />
            </div>
        )
    }
}

export default AudioPlayer;