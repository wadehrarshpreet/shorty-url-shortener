import {
  INIT_SHORTEN_URL,
  SHORTEN_URL_SUCCESS,
  SHORTEN_URL_FAILED
} from './short.constant';
import api from '../utils/api';
import { logoutUser } from './auth';

export const shortenURL = ({ url, custom = '' }) => async (
  dispatch,
  getState
) => {
  const requestURL = `${url}${custom ? `--${custom}` : ''}`;
  const {
    shortUrl: { data: urlMap }
  } = getState();

  let cache = false;

  if (urlMap.has(requestURL)) {
    cache = true;
  }

  dispatch({
    type: INIT_SHORTEN_URL,
    data: requestURL,
    cache
  });

  if (cache) return;

  // console.log(short.data);

  const { error, response } = await api.post('/api/v1/short', { url, custom });
  if (response) {
    dispatch({ type: SHORTEN_URL_SUCCESS, data: response?.data });
  } else {
    if (error?.data?.errorCode === 10010) {
      dispatch(logoutUser());
      dispatch({
        type: SHORTEN_URL_FAILED,
        error: {
          error: true,
          message: 'Session expired please login again to make custom url.'
        }
      });
      return;
    }
    dispatch({ type: SHORTEN_URL_FAILED, error: error?.data || error });
  }
};
