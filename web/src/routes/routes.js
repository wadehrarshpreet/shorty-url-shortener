import React from 'react';
import Loadable from 'react-loadable';
import { Spin } from 'antd';

const Loading = () => <Spin />;

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
    exact: true,
    path: '/about',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "contact", webpackMode: "lazy" */ '../component/About'),
      loading: Loading
    })
  },
  {
    exact: true,
    path: '/contact',
    component: Loadable({
      loader: () => import(/* webpackChunkName: "about", webpackMode: "lazy" */ '../component/Contact'),
      loading: Loading
    })
  }
];
