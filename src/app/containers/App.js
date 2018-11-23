import React from 'react';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import Header from '../components/Header.js';
import MainSection from '../components/MainSection.js';
import TodoActions from '../actions/index.js';

var App = React.createClass({
  propTypes: {
    todos: React.PropTypes.array.isRequired,
    actions: React.PropTypes.object.isRequired
  },
  render: function () {
    var todos = this.props.todos;
    var actions = this.props.actions;
    return (
      <div>
        <Header
          addTodo={actions.addTodo}
          />
        <MainSection
          todos={todos}
          actions={actions}
          />
      </div>
    );
  }

});

function mapStateToProps(state) {
  return {
    todos: state.todos
  };
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(TodoActions, dispatch)
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App);
