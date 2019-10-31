import React from 'react';
import AudioElement from './AudioElement.jsx';
import '../../styles/audio-list/AudioList.css';

class AudioList extends React.Component {
    constructor(props){
        super(props);
    }

    render() {
        return (
            <div id="audio-list">
                {this.props.tracks.map(track => {
                    return <AudioElement 
                            key={track.id} 
                            id={track.id} 
                            author={track.author}
                            title={track.title}
                            playingId={this.props.playingId}
                            paused={!this.props.isPlaying}
                            playAudio={this.props.playAudio}
                            resumeAudio={this.props.resumeAudio}
                            pauseAudio={this.props.pauseAudio}
                            downloadAudio={this.props.downloadAudio}
                            deleteAudio={this.props.deleteAudio}
                            />
                    })
                }
            </div>
        )
    }
    
}

export default AudioList;