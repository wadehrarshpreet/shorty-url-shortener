import React from 'react';
import { Link, withRouter } from 'react-router-dom';
import { connect } from 'react-redux';

import './Header.scss';
import { logoutUser } from '../../actions/auth';

const AppHeader = ({ history, userData, logout }) => {
  const isLoggedIn = userData?.token;

  React.useEffect(() => {
    // logged in & page = signin/signup redirect
    if (userData && ['/login', '/signup'].indexOf(history?.location?.pathname) !== -1) {
      history.push('/');
    }
  }, [userData, history.location]);

  return (
    <header>
      <div className='content'>
        <div className='left-section'>
          <div className='logo-container' onClick={() => history.push('/')}>
            <img alt='logo' className='logo' src='/assets/logo.png' /> <span className='title'>Shorty</span>
          </div>
        </div>
        <div className='right-section'>
          {isLoggedIn ? (
            <>
              <div className='nav-item'>Hi, {userData?.username}</div>
              <div className='nav-item'>
                <a
                  href='#'
                  onClick={() => {
                    logout();
                  }}
                >
                  Logout
                </a>
              </div>
            </>
          ) : (
            <>
              <div className='nav-item'>
                <Link to='/login'>Login</Link>
              </div>
              <div className='nav-item'>
                <Link to='/signup'>Sign up</Link>
              </div>
            </>
          )}
        </div>
      </div>
    </header>
  );
};

const mapStateToProps = ({ auth }) => ({ userData: auth?.data });

export default withRouter(connect(mapStateToProps, { logout: logoutUser })(AppHeader));
