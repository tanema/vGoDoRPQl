import React, { Component, PropTypes } from 'react';
import Input from '../components/Input'
import gql from 'graphql-tag'
import { graphql } from 'react-apollo'

const addTodoMutation = graphql(gql`mutation addTodo($text: String!) {
  createTodo(done: false, text: $text) { id }
}`, {
  name: 'addTodo'
});

class AddTodo extends Component {
  static propTypes = {
    todos: PropTypes.object.isRequired,
    addTodo: PropTypes.func.isRequired,
  }

  handleSave (text) {
    const refetch = this.props.todos.refetch;
    this.props.addTodo({variables: {text: text}}).then((e) => {
      refetch();
    })
  }

  render () {
    return (
      <Input
        className="new-todo"
        placeholder="What needs to be done?"
        onSave={this.handleSave.bind(this)}
      />
    )
  }
}

export default addTodoMutation(AddTodo);
