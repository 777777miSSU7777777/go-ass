import React from 'react';
import '../../styles/search-audio-form/SearchAudioField.css';

class SearchAudioField extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
                <input type='text' id='search-audio-field' onChange={this.props.onChange} onKeyUp={this.props.onKeyUp}/>
            )
    }
}

export default SearchAudioField;