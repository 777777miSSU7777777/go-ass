
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
        this.updateSearchKey = this.updateSearchKey.bind(this);
        this.searchAudio = this.searchAudio.bind(this);
        this.updateNewAuthor = this.updateNewAuthor.bind(this);
        this.updateNewTitle = this.updateNewTitle.bind(this);
        this.updateNewFile = this.updateNewFile.bind(this);
        this.newAudio = this.newAudio.bind(this);
        this.downloadAudio = this.downloadAudio.bind(this);
        this.deleteAudio = this.deleteAudio.bind(this);
        this.state = {tracks: [], audioModalState: "closed", isPlaying: false, player: null}
    }

    componentWillMount(){
        fetch("http://localhost:8080/api/audio", {method: "GET"})
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

    updateSearchKey(e){
        this.setState({searchKey: e.currentTarget.value});
    }

    searchAudio(){
        fetch("http://localhost:8080/api/audio?key=" + this.state.searchKey, {method: "GET"})
        .then(resp => resp.json())
        .then(data => {
            this.setState({tracks: data["audio"]});
            this.clearAudioForm();
            this.closeAudioFormModal();
        })
        .catch(error => console.error(error));
    }

    updateNewAuthor(e){
        this.setState({newAuthor: e.currentTarget.value});
    }

    updateNewTitle(e){
        this.setState({newTitle: e.currentTarget.value});
    }

    updateNewFile(e){
        this.setState({newFile: e.currentTarget.files[0]});
    }

    newAudio(){
        let formData = new FormData();
        formData.append("author", this.state.newAuthor);
        formData.append("title", this.state.newTitle);
        formData.append("audiofile", this.state.newFile);
        fetch("http://localhost:8080/api/audio", {method: "POST", body: formData})
        .then(resp => resp.json())
        .then(data => this.setState({tracks: [...this.state.tracks, data]}))
        .catch(error => console.error(error));
    }

    downloadAudio(e){
        let a = document.createElement("a");
        let id = e.currentTarget.parentElement.parentElement.id;
        let {author , title} = this.state.tracks.find(v => v.id == id)
        a.href = "/media/" +  id + "/download";   
        a.setAttribute("download", author + " - " + title);
        a.click();
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
                playingId={this.state.playingId} 
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
                downloadAudio={this.downloadAudio}
                deleteAudio={this.deleteAudio}
                />
                <AudioModalForm 
                state={this.state.audioModalState} 
                closeModal={this.closeAudioFormModal} 
                updateNewAuthor={this.updateNewAuthor}
                updateNewTitle={this.updateNewTitle}
                updateNewFile={this.updateNewFile}
                newAudio={this.newAudio}
                />
            </div>
        );
    }
}

export default App;