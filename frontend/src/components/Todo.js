import React, { PropTypes, Component } from 'react'
import classnames from 'classnames'
import gql from 'graphql-tag'
import { graphql, compose } from 'react-apollo'

import Input from '../components/Input'

const todoMutations = compose(
  graphql(gql`mutation renameTodo($id: Int!, $text: String!) {
    updateTodos(ids: [$id], changes: {text: $text}) { id }
  }`, { name: 'renameTodo' }),
  graphql(gql`mutation deleteTodo($id: Int!) {
    deleteTodos(ids: [$id]) { id }
  }`, { name: 'deleteTodo' }),
  graphql(gql`mutation toggleTodo($id: Int!, $complete: Boolean!) {
    updateTodos(ids: [$id], changes: {done: $complete}) { id }
  }`, { name: 'toggleTodo' }),
);

class Todo extends Component {
  static propTypes ={
    todo: PropTypes.object.isRequired,
    refetch: PropTypes.func.isRequired,
    renameTodo: PropTypes.func.isRequired,
    deleteTodo: PropTypes.func.isRequired,
    toggleTodo: PropTypes.func.isRequired,
  }

  state = {
    isEditing: false,
  }

  toggleDone(e) {
    this.props.toggleTodo({
      variables: {
        id: this.props.todo.id,
        complete: e.target.checked
      }
    }).then(this.props.refetch);
  }

  removeTodo () {
    this.toggleEditing(false)
    this.props.deleteTodo({variables: {id: this.props.todo.id}})
      .then(this.props.refetch);
  }

  toggleEditing(isEditing) {
    this.setState({isEditing})
  }

  renameTodo(newText) {
    this.toggleEditing(false)
    this.props.renameTodo({
      variables: {
        id: this.props.todo.id,
        text: newText
      }
    }).then(this.props.refetch);
  }

  render () {
    return (
      <li
        className={classnames({
          completed: this.props.todo.done,
          editing: this.state.isEditing,
        })}>
        <div className='view'>
          <input
            checked={this.props.todo.done}
            className='toggle'
            onChange={this.toggleDone.bind(this)}
            type='checkbox'
          />
          <label onDoubleClick={this.toggleEditing.bind(this, true)}>
            {this.props.todo.text}
          </label>
          <button
            className='destroy'
            onClick={this.removeTodo.bind(this)}
          />
        </div>
        {!this.state.isEditing ? null :
          <Input
            className='edit'
            initialValue={this.props.todo.text}
            onCancel={this.toggleEditing.bind(this, false)}
            onDelete={this.removeTodo.bind(this)}
            onSave={this.renameTodo.bind(this)}
          />
        }
      </li>
    )
  }
}

export default todoMutations(Todo);
