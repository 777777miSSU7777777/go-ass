import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as audioSearchFormActions from '../store/audio-search-form/actions';
import * as tracksActions from '../store/tracks/actions';
import AudioSearchForm from '../components/audio-search-form/AudioSearchForm.jsx';


const mapStateToProps = state => ({
    searchKey: state.audioSearchForm.searchKey
});

const mapDispatchToProps = dispatch => ({
    dispatchUpdateSearchKey: bindActionCreators(audioSearchFormActions.updateSearchKey, dispatch),
    dispatchFetchTracksByKey: bindActionCreators(tracksActions.fetchTracksByKey, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(AudioSearchForm);