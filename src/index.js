import 'babel-polyfill';

import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import App from './app/containers/App.js';
import configureStore from './app/store/configureStore.js';

import 'todomvc-app-css/index.css!';

var store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <App/>
  </Provider>,
  document.getElementById('root')
);
