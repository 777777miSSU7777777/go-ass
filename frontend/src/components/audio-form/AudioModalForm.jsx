import React from 'react';
import AudioFormTextField from './AudioFormTextField.jsx';
import AudioFormFileField from './AudioFormFileField.jsx';
import AudioFormSubmit from './AudioFormSubmit.jsx';
import '../../styles/audio-form/AudioModalForm.css';
import autoBind from 'react-autobind';


class AudioModalForm extends React.Component {
    constructor(props) {
        super(props);
        autoBind(this);
        this.fileInput = null;
    }

    async onFormSubmit() {
        const { dispatchValidateForm } = this.props;
        const { dispatchNewTrack } = this.props;

        await dispatchValidateForm();

        if (this.props.isFormValid) {
            dispatchNewTrack();
            this.close();
        }
    }

    updateAuthor(e) {
        const author = e.currentTarget.value;
        const { dispatchUpdateAuthor } = this.props;
        
        dispatchUpdateAuthor(author);
    }

    updateTitle(e) {
        const title = e.currentTarget.value;
        const { dispatchUpdateTitle } = this.props;

        dispatchUpdateTitle(title)
    }

    updateFile(e) {
        const file = e.currentTarget.files[0];
        const { dispatchUpdateFile } = this.props;
        
        dispatchUpdateFile(file);
    }

    resetFile() {
        this.fileInput.value = null;
    }

    setFileInput(fileInput) {
        this.fileInput = fileInput;
    }

    close() {
        const { dispatchCloseForm } = this.props;
        
        dispatchCloseForm();
        this.resetFile();
    }

    render() {
        const style = {};
        if (this.props.opened){
            style.display = 'block';
        } else {
            style.display = 'none';
        }
        
        return (
            <div className='audio-form-modal' className='modal' style={style}>
                <div className='modal-content'>
                    <div className='modal-header'>
                        <span 
                        className='close-audio-form-modal' 
                        onClick={this.close}
                        >
                        &times;
                        </span>
                        <h2>New Audio</h2>
                    </div>
                    <div className='modal-body'>
                        <div className='add-audio-form'>
                            <AudioFormTextField 
                            fieldClass='audio-author-field' 
                            label='Author'
                            value={this.props.form.author}
                            onChange={this.updateAuthor}
                            error={this.props.errors.author}
                            />
                            <AudioFormTextField 
                            fieldClass='audio-title-field' 
                            label='Title'
                            value={this.props.form.title}
                            onChange={this.updateTitle}
                            error={this.props.errors.title}
                            />
                            <AudioFormFileField 
                            className='audio-file-field' 
                            label='File'
                            onChange={this.updateFile}
                            error={this.props.errors.file}
                            setFileInput={this.setFileInput}
                            />
                            <AudioFormSubmit onClick={this.onFormSubmit} />
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default (AudioModalForm);