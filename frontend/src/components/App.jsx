
import React, { Component } from "react";
import AudioPlayer from '../containers/AudioPlayer.js';
import SearchAudioForm from '../containers/AudioSearchForm.js';
import AudioList from '../containers/AudioList.js';
import OpenAudioFormButton from './audio-form/OpenAudioFormButton.jsx';
import AudioModalForm from '../containers/AudioModalForm.js';
import '../styles/App.css';
import autoBind from 'react-autobind';

class App extends React.Component {
    constructor(props){
        super(props);
        autoBind(this);
    }

    openAudioFormModal() {
        const { dispatchOpenAudioForm } = this.props;

        dispatchOpenAudioForm();
    }

    render() {
        return (
            <div id='go-ass'>
                <AudioPlayer />
                <SearchAudioForm />
                <OpenAudioFormButton openModal={this.openAudioFormModal} />
                <AudioList />
                <AudioModalForm />
            </div>
        );
    }
}

export default App;