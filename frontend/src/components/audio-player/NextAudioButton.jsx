import React from 'react';
import '../../styles/audio-player/NextAudioButton.css';

class NextAudioButton extends React.Component {
    constructor(props){
        super(props);
    }
    render(){
        return (
            <div id="next-audio-button" onClick={this.props.onClick}></div>
        )
    }
}

export default NextAudioButton;