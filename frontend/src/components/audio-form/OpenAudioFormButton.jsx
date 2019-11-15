import React from 'react';
import '../../styles/audio-form/OpenAudioFormButton.css';

class OpenAudioFormButton extends React.Component {
    constructor(props){
        super(props);
    }
    render(){
        return (
            <input type="button" value="+" align="center" id="open-audio-form-modal" onClick={this.props.openModal} />
        )
    }
}

export default OpenAudioFormButton;