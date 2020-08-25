import React from 'react';
import { renderRoutes } from 'react-router-config';
import Header from './component/Header/Header';
import routes from './routes/routes';

const App = () => (
  <div className='app-wrapper'>
    <Header />
    <div className='app-container'>{renderRoutes(routes)}</div>
    <footer>&copy; Shorty {new Date().getFullYear()}</footer>
  </div>
);

export default App;
