import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import reducer from './reducers';

export default function configureStore(initialState = {}) {
  const middlewares = [thunk];
  const enhancers = [applyMiddleware(...middlewares)];
  // If Redux DevTools Extension is installed use it, otherwise use Redux compose
  /* eslint-disable no-underscore-dangle */
  // TODO Try to remove when `react-router-redux` is out of beta, LOCATION_CHANGE should not be fired more than once after hot reloading
  // Prevent recomputing reducers for `replaceReducer`
  const composeEnhancers = process.env.NODE_ENV === 'development' && typeof window === 'object' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ? window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({ shouldHotReload: false }) : compose;
  const store = createStore(reducer, initialState, composeEnhancers(...enhancers));
  return store;
}
