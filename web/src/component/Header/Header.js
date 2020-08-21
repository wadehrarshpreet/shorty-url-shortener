import React from 'react';
import { Link, withRouter } from 'react-router-dom';
import './Header.scss';

const AppHeader = ({ history }) => (
  <header>
    <div className='content'>
      <div className='left-section'>
        <div className='logo-container' onClick={() => history.push('/')}>
          <img alt='logo' className='logo' src='/assets/logo.png' /> <span className='title'>Shorty</span>
        </div>
      </div>
      <div className='right-section'>
        <div className='nav-item'>
          <Link to='/login'>Login</Link>
        </div>
        <div className='nav-item'>
          <Link to='/signup'>Sign up</Link>
        </div>
      </div>
    </div>
  </header>
);

export default withRouter(AppHeader);
