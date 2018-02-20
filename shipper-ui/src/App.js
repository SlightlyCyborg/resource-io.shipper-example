import React, { Component } from 'react';
import './App.css';
import CreateConsignment from './CreateConsignment';
import Authenticate from './Authenticate';

class App extends Component {

    constructor(props){
        super(props);
        this.state = {
            err: null,
            authenticated: false,
            token: null
        };
    }

  onAuth = (token) => {
    this.setState({
        authenticated: true,
        token:token
    });
  }

  renderLogin = () => {
    return (
      <Authenticate onAuth={this.onAuth} />
    );
  }

  renderAuthenticated = () => {
    return (
            <CreateConsignment token={this.getToken()}/>
    );
  }

  getToken = () => {
      return this.state.token;
  }

  isAuthenticated = () => {
    return this.state.authenticated || this.getToken() || false;
  }

  render() {
    const authenticated = this.isAuthenticated();
    return (
      <div className="App">
        <div className="App-header">
          <h2>Shippy</h2>
        </div>
        <div className='App-intro container'>
          {(authenticated ? this.renderAuthenticated() : this.renderLogin())}
        </div>
      </div>
    );
  }
}

export default App;
