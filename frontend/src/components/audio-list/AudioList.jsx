import React from 'react';
import AudioElement from './AudioElement.jsx';
import '../../styles/audio-list/AudioList.css';
import autoBind from 'react-autobind';


class AudioList extends React.Component {
    constructor(props) {
        super(props);
        autoBind(this);
    }
    
    componentDidMount() { 
        const { dispatchFetchTracks } = this.props;

        dispatchFetchTracks();
    }

    playAudio(e) {
        const id = e.currentTarget.parentElement.id;
        const { dispatchPlay } = this.props;

        dispatchPlay(id);
    }

    resumeAudio() {
        const { dispatchResume } = this.props;

        dispatchResume();

        this.props.player.play();
    }

    pauseAudio() {
        const { dispatchPause } = this.props;

        dispatchPause();

        this.props.player.pause();
    }

    deleteAudio(e) {
        const id = e.currentTarget.parentElement.parentElement.id;
        const { dispatchDeleteTrack } = this.props;

        dispatchDeleteTrack(id);
    }

    render() {
        return (
            <div className='audio-list'>
                {this.props.tracks.map(track => {
                    return <AudioElement 
                            key={track.id} 
                            id={track.id} 
                            author={track.author}
                            title={track.title}
                            playingId={this.props.playingId}
                            paused={!this.props.isPlaying}
                            playAudio={this.playAudio}
                            resumeAudio={this.resumeAudio}
                            pauseAudio={this.pauseAudio}
                            deleteAudio={this.deleteAudio}
                            />
                    })
                }
            </div>
        )
    }
    
}

export default AudioList;