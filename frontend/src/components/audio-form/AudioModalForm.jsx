import React from 'react';
import AudioFormTextField from '../audio-form/AudioFormTextField.jsx';
import AudioFormFileField from '../audio-form/AudioFormFileField.jsx';
import AudioFormSubmit from '../audio-form/AudioFormSubmit.jsx';
import '../../styles/audio-form/AudioModalForm.css';


class AudioModalForm extends React.Component {
    constructor(props){
        super(props);
    }

    render(){
        let style = {};
        if (this.props.state == "opened"){
            style.display = "block"
        } else if (this.props.state == "closed"){
            style.display = "none"
        }
        
        return (
            <div id="audio-form-modal" className="modal" style={style}>
                <div className="modal-content">
                    <div className="modal-header">
                        <span 
                        className="close-audio-form-modal" 
                        id="close-audio-form-modal" 
                        onClick={this.props.closeModal}
                        >
                        &times;
                        </span>
                        <h2>New Audio</h2>
                    </div>
                    <div className="modal-body">
                        <div id="add-audio-form">
                            <AudioFormTextField 
                            fieldId="audio-author-field" 
                            label="Author"
                            onChange={this.props.updateNewAuthor}
                            />
                            <AudioFormTextField 
                            fieldId="audio-title-field" 
                            label="Title" 
                            onChange={this.props.updateNewTitle}
                            />
                            <AudioFormFileField 
                            fieldId="audio-file-field/" 
                            label="File" 
                            onChange={this.props.updateNewFile}
                            />
                            <AudioFormSubmit onClick={this.props.newAudio}/>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default AudioModalForm;