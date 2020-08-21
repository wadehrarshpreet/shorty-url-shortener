import React from 'react';
import Loadable from 'react-loadable';

const Loading = () => <>Loading...</>;

export default [
  {
    exact: true,
    path: '/',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../component/Home'),
      loading: Loading
    })
  },
  {
    path: '/signup',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../component/SignUp/SignUp'),
      loading: Loading
    })
  },
  {
    exact: true,
    path: '/login',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "home", webpackMode: "lazy" */ '../component/Login/Login'),
      loading: Loading
    })
  }
];
