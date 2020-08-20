import React from 'react';
import { Link } from 'react-router-dom';
import { Menu, Layout } from 'antd';

const { Header } = Layout;

const LEFT_NAV_ITEMS = [
  {
    id: 'home',
    label: 'Home',
    path: '/'
  },
  {
    id: 'about',
    label: 'About',
    path: '/about'
  },
  {
    id: 'contact',
    label: 'Contact',
    path: '/contact'
  }
];

const AppHeader = () => (
  <Header className='header'>
    <div className='logo' />
    <Menu theme='dark' mode='horizontal' defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }}>
      {LEFT_NAV_ITEMS.map((navItem) => (
        <Menu.Item key={navItem.id}>
          <Link to={navItem.path}>{navItem.label}</Link>
        </Menu.Item>
      ))}
    </Menu>
  </Header>
);

export default AppHeader;
