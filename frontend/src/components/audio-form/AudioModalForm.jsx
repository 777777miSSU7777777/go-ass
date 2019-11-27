import React from 'react';
import AudioFormTextField from '../audio-form/AudioFormTextField.jsx';
import AudioFormFileField from '../audio-form/AudioFormFileField.jsx';
import AudioFormSubmit from '../audio-form/AudioFormSubmit.jsx';
import { isEmpty } from 'lodash-es';
import '../../styles/audio-form/AudioModalForm.css';


class AudioModalForm extends React.Component {
    constructor(props){
        super(props);
        this.state = {form: {author: "", title: "", file: null}, errors: {}, fileInput: null};
        this.updateAuthor = this.updateAuthor.bind(this);
        this.updateTitle = this.updateTitle.bind(this);
        this.updateFile = this.updateFile.bind(this);
        this.onFormSubmit = this.onFormSubmit.bind(this);
        this.validate = this.validate.bind(this);
        this.reset = this.reset.bind(this);
        this.setFileInput = this.setFileInput.bind(this);
        this.resetFile = this.resetFile.bind(this);
        this.resetErrors = this.resetErrors.bind(this);
        this.close = this.close.bind(this);
        this.fileInput = null;
    }

    onFormSubmit(){
        const errors = this.validate();
        this.setState({errors: errors});
        if (isEmpty(errors)){
            this.props.newAudio(this.state.form);
            this.close();
        }
    }

    updateAuthor(e){
        const author = e.currentTarget.value;
        this.setState(prevState => {
            const form = { ...prevState.form };
            form.author = author;
            return { form };
        });
    }

    updateTitle(e){
        const title = e.currentTarget.value;
        this.setState(prevState => {
            const form = { ...prevState.form };
            form.title = title;
            return { form };
        });
    }

    updateFile(e){
        const file = e.currentTarget.files[0];
        this.setState(prevState => {
            const form = { ...prevState.form };
            form.file = file;
            return { form };
        });
    }

    validate(){
        const errors = {};
        const {author, title, file} = this.state.form;
        if (author.length == 0){
            errors.author = "Author is empty";
        } else if (author.length > 50){
            errors.author = "Author length is more than 50 symbols"
        }

        if (title.length == 0){
            errors.title = "Title is empty";
        } else if (title.length > 50){
            errors.title = "Title length is more than 50 symbols";
        }

        if (!file){
            errors.file = "No file";
        } else if (file.type != "audio/mp3"){
            errors.file = "File is not mp3";
        }

        return errors;
    }

    reset(){
        this.setState({form: {author: "", title: "", file: null}});
    }

    resetFile(){
        this.fileInput.value = null;
    }

    setFileInput(fileInput){
        this.fileInput = fileInput;
    }

    resetErrors(){
        this.setState({errors: {}});
    }

    close(){
        this.reset();
        this.resetFile();
        this.resetErrors();
        this.props.closeModal();
    }

    render(){
        const style = {};
        if (this.props.state == "opened"){
            style.display = "block";
        } else if (this.props.state == "closed"){
            style.display = "none";
        }
        
        return (
            <div id="audio-form-modal" className="modal" style={style}>
                <div className="modal-content">
                    <div className="modal-header">
                        <span 
                        className="close-audio-form-modal" 
                        id="close-audio-form-modal" 
                        onClick={this.close}
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
                            value={this.state.form.author}
                            onChange={this.updateAuthor}
                            error={this.state.errors.author}
                            />
                            <AudioFormTextField 
                            fieldId="audio-title-field" 
                            label="Title" 
                            value={this.state.form.title}
                            onChange={this.updateTitle}
                            error={this.state.errors.title}
                            />
                            <AudioFormFileField 
                            fieldId="audio-file-field" 
                            label="File" 
                            onChange={this.updateFile}
                            error={this.state.errors.file}
                            setFileInput={this.setFileInput}
                            />
                            <AudioFormSubmit onClick={this.onFormSubmit}/>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default AudioModalForm;