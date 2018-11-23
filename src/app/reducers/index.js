import {combineReducers} from 'redux';
import todos from './todos.js';

var rootReducer = combineReducers({
  todos: todos
});

export default rootReducer;
