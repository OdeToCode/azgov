import {createStore} from 'redux';
import rootReducer from '../reducers/index.js';

function configureStore(initialState) {
  var store = createStore(rootReducer, initialState);
  return store;
}

export default configureStore;
