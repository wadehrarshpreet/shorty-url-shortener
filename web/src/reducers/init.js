import { INIT } from '../actions/init';

const INIT_STATE = {
  loading: true
};
export default function (state = INIT_STATE, action) {
  switch (action.type) {
    case INIT: {
      return {
        ...state,
        loading: false
      };
    }
    default:
      return state;
  }
}
