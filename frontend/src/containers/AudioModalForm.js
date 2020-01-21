import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as audioFormActions from '../store/audio-form/actions';
import * as tracksActions from '../store/tracks/actions';
import AudioSearchForm from '../components/audio-form/AudioModalForm.jsx';
import isEmpty from 'lodash-es/isEmpty.js';


const mapStateToProps = state => ({
    opened: state.audioForm.opened,
    form: state.audioForm.form,
    errors: state.audioForm.formErrors,
    isFormValid: isEmpty(state.audioForm.formErrors)
});

const mapDispatchToProps = dispatch => ({
    dispatchOpenForm: bindActionCreators(audioFormActions.openForm, dispatch),
    dispatchUpdateAuthor: bindActionCreators(audioFormActions.updateAuthor, dispatch),
    dispatchUpdateTitle: bindActionCreators(audioFormActions.updateTitle, dispatch),
    dispatchUpdateFile: bindActionCreators(audioFormActions.updateFile, dispatch),
    dispatchValidateForm: bindActionCreators(audioFormActions.validateForm, dispatch),
    dispatchNewTrack: bindActionCreators(tracksActions.newTrack, dispatch),
    dispatchCloseForm: bindActionCreators(audioFormActions.closeForm, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(AudioSearchForm);