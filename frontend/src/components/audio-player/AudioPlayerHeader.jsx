import React from 'react';
import '../../styles/audio-player/AudioPlayerHeader.css';

class AudioPlayerHeader extends React.Component {
    constructor(props){
        super(props);
    }
    render(){
        return (
            <div id="audio-player-header">
                {this.props.author} - {this.props.title}
            </div>
        )
    }
}

export default AudioPlayerHeader;