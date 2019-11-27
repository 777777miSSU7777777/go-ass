import React from 'react';
import '../../styles/audio-form/AudioFormFileField.css';

class AudioFormFileField extends React.Component {
    constructor(props){
        super(props);
        this.fileInput = React.createRef();
    }

    componentDidMount(){
        this.props.setFileInput(this.fileInput.current);
    }
    
    render(){
        return (
            <div className="audio-form-field">
                {this.props.error && <p class="validation-error">{this.props.error}</p>}
                <label 
                htmlFor={this.props.fieldId}
                >
                {this.props.label}
                </label>
                <input
                ref={this.fileInput}
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