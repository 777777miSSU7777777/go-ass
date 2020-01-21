import React from 'react';
import '../../styles/audio-form/AudioFormTextField.css';

class AudioFormTextField extends React.Component {
    constructor(props) {
        super(props);
    }    
    
    render() {
        return (
            <div className='audio-form-field'>
                {this.props.error && <p className='validation-error'>{this.props.error}</p>}
                <label 
                htmlFor={this.props.fieldId}
                >
                {this.props.label}
                </label>
                <input 
                type='text'
                id={this.props.fieldId}
                value={this.props.value}
                onChange={this.props.onChange}
                />
            </div>
        )
    }
}

export default AudioFormTextField;