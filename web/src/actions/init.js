export const INIT = 'INIT';

export function init() {
  return (dispatch) => {
    setTimeout(() => {
      dispatch({
        type: INIT
      });
    }, 2000);
  };
}
