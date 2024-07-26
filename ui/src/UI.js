import React from 'react';
import './UI.css';

import PersonsList from "./PersonsList";

class UI extends React.Component {
  constructor() {
    super()
    this.state = {
      persons: []
    };
    this.getData();
  };

  getData() {
    var xhr = new XMLHttpRequest()

    xhr.addEventListener('load', () => {
      this.setState({
        persons: JSON.parse(xhr.responseText)
      });
    })

    xhr.open('GET', 'http://localhost:8080/app/people?orderBy=email')
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send()
  };

  render() {
    return (
      <div className="UI">
        <header className="UI-header">
          <h1 className="UI-title">Persons Manager</h1>
          <PersonsList persons={this.state.persons} />
        </header>
      </div>
    );
  };
}

export default UI;
