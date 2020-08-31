import React from 'react';
import './Loader.scss';

const AppLoader = ({ message = 'Shortening...' }) => (
  <div className='loader'>
    <img alt='logo' className='logo' src='/assets/logo.png' />
    <div>{message}</div>
  </div>
);

export default AppLoader;
