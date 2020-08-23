import { INIT_LOGIN, LOGIN_FAILED, LOGIN_SUCCESS, LOGOUT, INIT_SIGNUP, SIGNUP_SUCCESS, SIGNUP_FAILED } from './auth.constant';
import api from '../utils/api';

export function loginUser(loginData) {
  return async (dispatch) => {
    dispatch({ type: INIT_LOGIN });

    const { error, response } = await api.post('/login', loginData);
    if (response) {
      localStorage.setItem('userData', JSON.stringify(response?.data));
      api.setAuthHeaders(response?.data?.token);
      dispatch({ type: LOGIN_SUCCESS, data: response?.data });
    } else {
      dispatch({ type: LOGIN_FAILED, error: error?.data || error });
    }
  };
}

export function logoutUser() {
  return (dispatch) => {
    localStorage.removeItem('userData');
    dispatch({ type: LOGOUT });
  };
}

export function signUpUser(userData) {
  return async (dispatch) => {
    dispatch({ type: INIT_SIGNUP });

    const { error, response } = await api.post('/signup', userData);
    if (response) {
      localStorage.setItem('userData', JSON.stringify(response?.data));
      api.setAuthHeaders(response?.data?.token);
      dispatch({ type: SIGNUP_SUCCESS, data: response?.data });
    } else {
      dispatch({ type: SIGNUP_FAILED, error: error?.data || error });
    }
  };
}
