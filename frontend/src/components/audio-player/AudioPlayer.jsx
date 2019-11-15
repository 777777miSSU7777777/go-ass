import React from 'react';
import AudioPlayerHeader from './AudioPlayerHeader.jsx';
import AudioPlayerControls from './AudioPlayerControls.jsx';
import '../../styles/audio-player/AudioPlayer.css';

class AudioPlayer extends React.Component {
    constructor(props){
        super(props);
    }

    componentDidUpdate(){
        if (this.props.track !== undefined){
            document.title = this.props.track.author + " - " + this.props.track.title;
        }
    }

    render(){
        let id, author, title;
        if (this.props.track !== undefined){
            id = this.props.track.id;
            author = this.props.track.author;
            title = this.props.track.title;
        } else {
            id = 0;
            author = "Author";
            title = "Title"
        }
        
        return(
            <div id="audio-player">
                <AudioPlayerHeader 
                author={author} 
                title={title} />
                <AudioPlayerControls 
                setPlayer={this.props.setPlayer}
                playingId={id}
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