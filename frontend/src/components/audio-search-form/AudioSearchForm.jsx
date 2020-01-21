import React from 'react';
import SearchAudioField from './SearchAudioField.jsx';
import SearchAudioButton from './SearchAudioButton.jsx';
import '../../styles/search-audio-form/SearchAudioForm.css';
import autoBind from 'react-autobind';

class AudioSearchForm extends React.Component {
    constructor(props) {
        super(props);
        autoBind(this);
    }

    updateSearchKey(e) {
        const searchKey = e.currentTarget.value;
        const { dispatchUpdateSearchKey } = this.props;
        
        dispatchUpdateSearchKey(searchKey);
    }

    onClick() {
        const { dispatchFetchTracksByKey } = this.props;

        dispatchFetchTracksByKey();
    }

    onKeyUp(e) {
        if (e.keyCode == 13) {
            const { dispatchFetchTracksByKey } = this.props;

            dispatchFetchTracksByKey();
        }
    }

    render() {
        return (
            <div id='search-audio-form'>
                <SearchAudioField onChange={this.updateSearchKey} onKeyUp={this.onKeyUp}/>
                <SearchAudioButton onClick={this.onClick}/>
            </div>
        )
    }
}

export default AudioSearchForm;