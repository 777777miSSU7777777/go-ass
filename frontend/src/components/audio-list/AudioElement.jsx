import React from 'react';
import '../../styles/audio-list/AudioElement.css';

class AudioElement extends React.Component {
    constructor(props){
        super(props);
        this.downloadAudio = this.downloadAudio.bind(this);
    }

    downloadAudio(e){
        let a = document.createElement("a");
        let id = e.currentTarget.parentElement.parentElement.id;
        let {author , title} = this.props;
        a.href = "/media/" +  id + "/download";   
        a.setAttribute("download", author + " - " + title);
        a.click();
    }
    
    render(){
        let playButton = null;
        if (this.props.playingId == this.props.id && !this.props.paused){
            playButton = <div className="pause-button" onClick={this.props.pauseAudio}></div>
        } else if (this.props.playingId == this.props.id && this.props.paused){
            playButton = <div className="play-button" onClick={this.props.resumeAudio}></div>
        } else if(this.props.playingId != this.props.id){
            playButton = <div className="play-button" onClick={this.props.playAudio}></div>
        }
        return (
            <div className="audio-element" id={this.props.id}>
                {playButton}
                <p className="audio-info">{this.props.author} - {this.props.title}</p>
            <div className="right-controls">
                <div className="download-button" onClick={this.downloadAudio}></div>
                <div className="delete-button" onClick={this.props.deleteAudio}></div>
            </div>
            </div> 

        )
    }
}

export default AudioElement;