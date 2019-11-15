import React from 'react';
import SearchAudioField from './SearchAudioField.jsx';
import SearchAudioButton from './SearchAudioButton.jsx';
import '../../styles/search-audio-form/SearchAudioForm.css';

class SearchAudioForm extends React.Component {
    constructor(props){
        super(props);
        this.state = {searchKey: ""};
        this.onClick = this.onClick.bind(this);
    }

    updateSearchKey(e){
        this.setState({searchKey: e.currentTarget.value});
    }

    onClick(){
        this.props.searchAudio(this.state.searchKey);
    }

    render(){
        return (
            <div id="search-audio-form">
                <SearchAudioField onChange={this.updateSearchKey}/>
                <SearchAudioButton onClick={this.onClick}/>
            </div>
        )
    }
}

export default SearchAudioForm;