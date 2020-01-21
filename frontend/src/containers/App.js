import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as audioFormActions from '../store/audio-form/actions';
import App from '../components/App.jsx';

const mapDispatchToProps = dispatch => ({
    dispatchOpenAudioForm: bindActionCreators(audioFormActions.openForm, dispatch),
});

export default connect(null, mapDispatchToProps)(App);