import { INIT_SHORTEN_URL, SHORTEN_URL_SUCCESS, SHORTEN_URL_FAILED } from '../actions/short.constant';

const INIT_STATE = {
  fetching: false,
  data: null
};

export default function(action, state = INIT_STATE) {
  switch (action.type) {
    case INIT_SHORTEN_URL:
      return { ...state, data: null, fetching: true, error: false };
    case SHORTEN_URL_SUCCESS:
      return { ...state, error: false, fetching: false };
    case SHORTEN_URL_FAILED:
      return { ...state, error: action.error, fetching: false };
    default:
      return state;
  }
}
