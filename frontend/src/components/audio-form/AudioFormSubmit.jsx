import React from 'react';
import '../../styles/audio-form/AudioFormSubmit.css';

class AudioFormSubmit extends React.Component {
    constructor(props){
        super(props);
    }

    render() {
        return (
            <input 
            type='submit' 
            value='Add'
            className='add-audio-button' 
            onClick={this.props.onClick} 
            />
        )
    }
}

export default AudioFormSubmit;