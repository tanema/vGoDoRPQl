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

  _handleCompleteChange = (e) => {
    this.props.toggleTodo({variables: {id: this.props.todo.id, complete: e.target.checked}})
      .then(this.props.refetch);
  }

  _handleDestroyClick = () => {
    this._removeTodo()
  }

  _handleLabelDoubleClick = () => {
    this._setEditMode(true)
  }

  _handleTextInputCancel = () => {
    this._setEditMode(false)
  }

  _handleTextInputDelete = () => {
    this._setEditMode(false)
    this._removeTodo()
  }

  _handleTextInputSave = (newText) => {
    this._setEditMode(false)
    this.props.renameTodo({variables: {id: this.props.todo.id, text: newText}})
      .then(this.props.refetch);
  }

  _removeTodo () {
    this.props.deleteTodo({variables: {id: this.props.todo.id}})
      .then(this.props.refetch);
  }

  _setEditMode = (shouldEdit) => {
    this.setState({isEditing: shouldEdit})
  }

  renderTextInput () {
    return (
      <Input
        className='edit'
        initialValue={this.props.todo.text}
        onCancel={this._handleTextInputCancel}
        onDelete={this._handleTextInputDelete}
        onSave={this._handleTextInputSave}
      />
    )
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
            onChange={this._handleCompleteChange}
            type='checkbox'
          />
          <label onDoubleClick={this._handleLabelDoubleClick}>
            {this.props.todo.text}
          </label>
          <button
            className='destroy'
            onClick={this._handleDestroyClick}
          />
        </div>
        {this.state.isEditing && this.renderTextInput()}
      </li>
    )
  }
}

export default todoMutations(Todo);
