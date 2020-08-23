import { INIT_LOGIN, LOGIN_FAILED, LOGIN_SUCCESS, LOGOUT, INIT_SIGNUP, SIGNUP_SUCCESS, SIGNUP_FAILED } from '../actions/auth.constant';
import api from '../utils/api';

let userData = null;
try {
  userData = JSON.parse(localStorage.getItem('userData'));
  api.setAuthHeaders(userData?.token);
} catch (e) {
  // e
}

const INIT_STATE = {
  login: {
    loading: false,
    error: false
  },
  data: userData
};

export default function(state = INIT_STATE, action) {
  switch (action.type) {
    case INIT_LOGIN:
    case INIT_SIGNUP: {
      return { ...state, signup: { loading: false, error: false }, login: { loading: false, error: false }, userData: null };
    }
    case LOGIN_FAILED: {
      return {
        ...state,
        login: {
          loading: false,
          error: action.error
        }
      };
    }
    case LOGIN_SUCCESS: {
      return { ...state, login: { loading: false, error: false }, data: action.data };
    }
    case LOGOUT: {
      return { ...state, data: null };
    }
    case SIGNUP_SUCCESS: {
      return { ...state, signup: { loading: false, error: false }, data: action.data };
    }
    case SIGNUP_FAILED: {
      return {
        ...state,
        signup: {
          loading: false,
          error: action.error
        }
      };
    }
    default:
      return state;
  }
}
