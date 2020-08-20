import React from 'react';

export default class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className='container'>
        <h2>This is React Redux BoilerPlace!</h2>
      </div>
    );
  }
}
