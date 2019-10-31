import React from 'react';
import SearchAudioField from './SearchAudioField.jsx';
import SearchAudioButton from './SearchAudioButton.jsx';
import '../../styles/search-audio-form/SearchAudioForm.css';

class SearchAudioForm extends React.Component {
    constructor(props){
        super(props);
    }

    render(){
        return (
            <div id="search-audio-form">
                <SearchAudioField onChange={this.props.updateSearchKey}/>
                <SearchAudioButton onClick={this.props.searchAudio}/>
            </div>
        )
    }
}

export default SearchAudioForm;