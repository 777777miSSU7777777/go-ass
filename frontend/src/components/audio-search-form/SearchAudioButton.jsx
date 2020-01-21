import React from 'react';
import '../../styles/search-audio-form/SearchAudioButton.css';

class SearchAudioButton extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <input type='button' id='search-audio-button' value='Search' onClick={this.props.onClick}/>
        )
    }
}

export default SearchAudioButton;