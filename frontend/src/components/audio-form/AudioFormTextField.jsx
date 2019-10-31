import React from 'react';
import '../../styles/audio-form/AudioFormTextField.css';

class AudioFormTextField extends React.Component {
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
                type="text" 
                id={this.props.fieldId}
                onChange={this.props.onChange}
                />
            </div>
        )
    }
}

export default AudioFormTextField;