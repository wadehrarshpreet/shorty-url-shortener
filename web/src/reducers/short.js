import {
  INIT_SHORTEN_URL,
  SHORTEN_URL_SUCCESS,
  SHORTEN_URL_FAILED
} from '../actions/short.constant';

const INIT_STATE = {
  fetching: false,
  data: new Map()
};

export default function(state = INIT_STATE, action) {
  switch (action.type) {
    case INIT_SHORTEN_URL:
      return {
        ...state,
        fetching: !action.cache,
        error: false,
        requestURL: action?.data
      };
    case SHORTEN_URL_SUCCESS:
      return {
        ...state,
        error: false,
        fetching: false,
        data: state.data.set(state.requestURL, action?.data?.shortUrl)
      };
    case SHORTEN_URL_FAILED:
      return { ...state, error: action.error, fetching: false };
    default:
      return state;
  }
}
