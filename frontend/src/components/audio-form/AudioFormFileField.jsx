import React from 'react';
import '../../styles/audio-form/AudioFormFileField.css';

class AudioFormFileField extends React.Component {
    constructor(props){
        super(props);
    }    
    render(){
        return (
            <div className="audio-form-field">
                <label 
                htmlFor={this.props.fieldId}
                >
                {this.props.label}
                </label>
                <input 
                type="file" 
                id={this.props.fieldId} 
                accept="audio/mp3" 
                onChange={this.props.onChange}
                />
            </div>
        )
    }
}

export default AudioFormFileField;