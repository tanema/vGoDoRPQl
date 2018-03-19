import React, { PropTypes, Component } from 'react'
import Todo from './Todo'
import gql from 'graphql-tag'
import { graphql } from 'react-apollo'

const toggleAllMutation = graphql(gql`
  mutation toggleAllTodo($ids: [Int!], $complete: Boolean!) {
    updateTodos(ids: $ids, changes: {done: $complete}) { id }
  }`,
  {
    name: 'toggleAllTodos'
  }
);

class List extends Component {
  static propTypes = {
    todos: PropTypes.object.isRequired,
    toggleAllTodos: PropTypes.func.isRequired,
  }

  get todos() {
    return (this.props.todos.todos || []);
  }

  get todo_ids() {
    return this.todos.map((todo) => todo.id);
  }

  get activeTodoCount() {
    return this.todos.reduce(function (accum, todo) {
        return todo.done ? accum : accum + 1;
    }, 0);
  }

  toggleAll (e) {
    var checked = e.target.checked;
    this.props.toggleAllTodos({variables: {ids: this.todo_ids, complete: checked}})
      .then(this.props.todos.refetch.bind(this));
  }

  render () {
    return (
      <section className='main'>
        <input
          className="toggle-all"
          type="checkbox"
          onChange={this.toggleAll.bind(this)}
          checked={this.activeTodoCount === 0}
        />
        <ul className='todo-list'>
          {this.todos.map((todo) =>
            <Todo
              key={todo.id}
              todo={todo}
              refetch={this.props.todos.refetch.bind(this)}
            />
          )}
        </ul>
      </section>
    )
  }
}

export default toggleAllMutation(List);
