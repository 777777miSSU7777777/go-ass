
import React, { Component } from "react";
import AudioPlayer from './audio-player/AudioPlayer.jsx';
import SearchAudioForm from './search-audio-form/SearchAudioForm.jsx';
import AudioList from './audio-list/AudioList.jsx';
import OpenAudioFormButton from './audio-form/OpenAudioFormButton.jsx';
import AudioModalForm from './audio-form/AudioModalForm.jsx';
import '../styles/App.css';

class App extends React.Component {
    constructor(props){
        super(props);
        this.openAudioFormModal = this.openAudioFormModal.bind(this);
        this.closeAudioFormModal = this.closeAudioFormModal.bind(this);
        this.playAudio = this.playAudio.bind(this);
        this.resumeAudio = this.resumeAudio.bind(this);
        this.pauseAudio = this.pauseAudio.bind(this);
        this.prevTrack = this.prevTrack.bind(this);
        this.nextTrack = this.nextTrack.bind(this);
        this.setPlayer = this.setPlayer.bind(this);
        this.setPlaying = this.setPlaying.bind(this);
        this.searchAudio = this.searchAudio.bind(this);
        this.newAudio = this.newAudio.bind(this);
        this.deleteAudio = this.deleteAudio.bind(this);
        this.state = {tracks: [], audioModalState: "closed", isPlaying: false, player: null};
    }

    componentWillMount(){
        fetch("/api/audio", {method: "GET"})
        .then(resp => resp.json())
        .then(data => this.setState({tracks: data["audio"]}))
        .catch(error => console.error(error));
    }

    openAudioFormModal(e){
        this.setState({audioModalState: "opened"});
    }

    closeAudioFormModal(e){
        this.setState({audioModalState: "closed"});
    }

    playAudio(e){
        this.setState({playingId: e.currentTarget.parentElement.id});
    }

    resumeAudio(e){
        this.state.player.play();
    }

    pauseAudio(e){
        this.state.player.pause();
    }

    setPlaying(playing){
        this.setState({isPlaying: playing});
    }

    setPlayer(player){
        this.setState({player: player});
    }

    prevTrack(){
        let arrId = this.state.tracks.findIndex(v => v.id == this.state.playingId);
        arrId = arrId == 0 ? this.state.tracks.length - 1: arrId - 1; 
        this.setState({playingId: this.state.tracks[arrId].id});
        this.state.player.play();
    }
    nextTrack(){
        let arrId = this.state.tracks.findIndex(v => v.id == this.state.playingId);
        arrId = arrId == this.state.tracks.length - 1 ? 0 : arrId + 1; 
        this.setState({playingId: this.state.tracks[arrId].id});
        this.state.player.play();
    }

    searchAudio(searchKey){
        fetch("/api/audio?key=" + searchKey, {method: "GET"})
        .then(resp => resp.json())
        .then(data => {
            this.setState({tracks: data["audio"]});
        })
        .catch(error => console.error(error));
    }

    newAudio(form){
        let formData = new FormData();
        formData.append("author", form.author);
        formData.append("title", form.title);
        formData.append("audiofile", form.file);
        fetch("/api/audio", {method: "POST", body: formData})
        .then(resp => resp.json())
        .then(data => this.setState({tracks: [...this.state.tracks, data]}))
        .catch(error => console.error(error));
    }

    deleteAudio(e){
        let id = e.currentTarget.parentElement.parentElement.id;
        fetch("/api/audio/" + id, {method: "DELETE"})
        .then(resp => {
            if (resp.status == 200){
                this.setState({tracks: this.state.tracks.filter(v => v.id != id)});
            }
        });
    }

    render() {
        return (
            <div id="go-ass">
                <AudioPlayer 
                setPlayer={this.setPlayer}
                track={this.state.tracks.find(e => e.id == this.state.playingId)} 
                isPlaying={this.state.isPlaying}
                setPlaying={this.setPlaying}
                prevTrack={this.prevTrack}
                nextTrack={this.nextTrack}
                />
                <SearchAudioForm
                updateSearchKey={this.updateSearchKey}
                searchAudio={this.searchAudio}
                />
                <OpenAudioFormButton 
                openModal={this.openAudioFormModal} 
                />
                <AudioList 
                tracks={this.state.tracks} 
                playingId={this.state.playingId} 
                isPlaying={this.state.isPlaying} 
                playAudio={this.playAudio}
                resumeAudio={this.resumeAudio}
                pauseAudio={this.pauseAudio}
                deleteAudio={this.deleteAudio}
                />
                <AudioModalForm 
                state={this.state.audioModalState} 
                closeModal={this.closeAudioFormModal} 
                newAudio={this.newAudio}
                />
            </div>
        );
    }
}

export default App;