import React from 'react';
import '../../styles/audio-player/PrevAudioButton.css';

class PrevAudioButton extends React.Component {
    constructor(props){
        super(props);
    }
    render(){
        return (
            <div id="prev-audio-button" onClick={this.props.onClick}></div>
        )
    }
}

export default PrevAudioButton;