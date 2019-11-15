import React from 'react';
import '../../styles/audio-list/AudioElement.css';

class AudioElement extends React.Component {
    constructor(props){
        super(props);
        this.downloadAudio = this.downloadAudio.bind(this);
    }

    downloadAudio(e){
        const a = document.createElement("a");
        const id = e.currentTarget.parentElement.parentElement.id;
        const {author , title} = this.props;
        a.href = "/media/" +  id + "/download";   
        a.setAttribute("download", author + " - " + title);
        a.click();
    }
    
    render(){
        let className, onClick;
        if (this.props.playingId == this.props.id && !this.props.paused){
            className = "pause-button";
            onClick = this.props.pauseAudio
        } else if (this.props.playingId == this.props.id && this.props.paused){
            className = "play-button";
            onClick = this.props.resumeAudio;
        } else if(this.props.playingId != this.props.id){
            className = "play-button";
            onClick = this.props.playAudio;
        }
        return (
            <div className="audio-element" id={this.props.id}>
                <div className={className} onClick={onClick}></div>
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