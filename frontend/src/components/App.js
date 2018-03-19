import React, { Component, PropTypes } from 'react';
import gql from 'graphql-tag'
import { graphql } from 'react-apollo'
import { SHOW_ALL } from '../constants/FilterTypes';

import AddTodo from './AddTodo';
import List from './List';
import Footer from './Footer';

const todosQuery = graphql(gql`
  query getTodos($status: TodoStatus){
    todos(status: $status) {
      id,
      text,
      done
    }
  }`,
  {
    name: 'todos',
    options: (props) => {
      return {
        variables: {
          status: SHOW_ALL
        }
      };
    }
  }
);

class App extends Component {
  static propTypes = {
    todos: PropTypes.object.isRequired,
  }

  render() {
    return (
      <section className="todoapp">
        <AddTodo todos={this.props.todos} />
        <List todos={this.props.todos} />
        <Footer todos={this.props.todos} />
      </section>
    );
  }
}

export default todosQuery(App);
