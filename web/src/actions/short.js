import { INIT_SHORTEN_URL, SHORTEN_URL_SUCCESS, SHORTEN_URL_FAILED } from './short.constant';
import api from '../utils/api';

export const shortenURL = (url) => async (dispatch) => {
  dispatch({ type: INIT_SHORTEN_URL });

  const { error, response } = await api.post('/shorten', { url });
  if (response) {
    dispatch({ type: SHORTEN_URL_SUCCESS, data: response?.data });
  } else {
    dispatch({ type: SHORTEN_URL_FAILED, error: error?.data || error });
  }
};
