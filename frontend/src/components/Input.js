import React, { PropTypes, Component } from 'react'
import ReactDOM from 'react-dom'

const ENTER_KEY_CODE = 13
const ESC_KEY_CODE = 27

export default class Input extends Component {
  static propTypes = {
    className: PropTypes.string,
    initialValue: PropTypes.string,
    onCancel: PropTypes.func,
    onDelete: PropTypes.func,
    onSave: PropTypes.func.isRequired,
    placeholder: PropTypes.string,
  }

  state = {
    text: this.props.initialValue || '',
  }

  componentDidMount () {
    ReactDOM.findDOMNode(this).focus()
  }

  commitChanges() {
    var newText = this.state.text.trim()
    if (this.props.onDelete && newText === '') {
      this.props.onDelete()
    } else if (this.props.onCancel && newText === this.props.initialValue) {
      this.props.onCancel()
    } else if (newText !== '') {
      this.props.onSave(newText)
      this.setState({text: ''})
    }
  }

  handleChange(e) {
    this.setState({text: e.target.value})
  }

  handleKeyDown(e) {
    if (this.props.onCancel && e.keyCode === ESC_KEY_CODE) {
      this.props.onCancel()
    } else if (e.keyCode === ENTER_KEY_CODE) {
      this.commitChanges()
    }
  }

  render () {
    return (
      <input
        className={this.props.className}
        onChange={this.handleChange.bind(this)}
        onKeyDown={this.handleKeyDown.bind(this)}
        placeholder={this.props.placeholder}
        value={this.state.text}
      />
    )
  }
}
