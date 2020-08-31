import React from 'react';
import Loadable from 'react-loadable';

const Loading = () => <>Loading...</>;

export default [
  {
    exact: true,
    path: '/signup',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../pages/SignUp/SignUp'),
      loading: Loading
    })
  },
  {
    exact: true,
    path: '/login',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../pages/Login/Login'),
      loading: Loading
    })
  },
  {
    path: '/',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../pages/Home/Home'),
      loading: Loading
    })
  }
];
