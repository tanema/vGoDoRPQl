import React, { Component, PropTypes } from 'react';
import classNames from 'classnames';
import gql from 'graphql-tag'
import { graphql } from 'react-apollo'
import { SHOW_ALL, SHOW_COMPLETED, SHOW_ACTIVE } from '../constants/FilterTypes';

const clearCompletedMutation = graphql(gql`mutation deleteTodos($ids: [Int!]) {
  deleteTodos(ids: $ids) { id }
}`, {
  name: 'clearTodos'
});

class Footer extends Component {
  static propTypes = {
    todos: PropTypes.object.isRequired,
    clearTodos: PropTypes.func.isRequired,
  }

  get todos() {
    return (this.props.todos.todos || []);
  }

  get completedTodoIDs() {
    return this.todos.reduce(function (accum, todo) {
      if (todo.done) {
        accum.push(todo.id);
      }
      return accum;
    }, []);
  }

  get activeTodoCount() {
    return this.todos.reduce(function (accum, todo) {
        return todo.done ? accum : accum + 1;
    }, 0);
  }

  clearCompleted() {
    this.props.todos.refetch({status: SHOW_ALL})
      .then(() => this.props.clearTodos({variables: {ids: this.completedTodoIDs}}))
      .then(() => this.props.todos.refetch());
  }

  setStatus(status) {
    this.props.todos.refetch({status});
  }

  render () {
    const { status } = this.props.todos.variables;
    return (
      <footer className="footer">
        <span className="todo-count">
          <strong>{ this.activeTodoCount }</strong> left
        </span>
        <ul className="filters">
          <li>
            <a
              onClick={this.setStatus.bind(this, SHOW_ALL)}
              className={classNames({selected: status === SHOW_ALL})}
            > All </a>
          </li>
          <li>
            <a
              onClick={this.setStatus.bind(this, SHOW_ACTIVE)}
              className={classNames({selected: status === SHOW_ACTIVE})}
            > Active </a>
          </li>
          <li>
            <a
              onClick={this.setStatus.bind(this, SHOW_COMPLETED)}
              className={classNames({selected: status === SHOW_COMPLETED})}
            > Completed </a>
          </li>
        </ul>
        <button
          className="clear-completed"
          onClick={this.clearCompleted.bind(this)}>
          Clear completed
        </button>
      </footer>

    )
  }
}

export default clearCompletedMutation(Footer);
