import { combineReducers } from 'redux';
import auth from './auth';
import shortUrl from './short';

const rootReducer = combineReducers({
  auth,
  shortUrl
});

export default rootReducer;
