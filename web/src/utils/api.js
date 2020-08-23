import axios from 'axios';
import querystring from 'querystring';

function get(location = '', body, apiBaseURL = API_BASE) {
  let url = `${apiBaseURL}${location}`;
  if (body) {
    const qs = querystring.stringify(body);
    if (qs) {
      url += (url.indexOf('?') >= 0 ? '&' : '?') + qs;
    }
  }
  return axios
    .get(`${url}`)
    .then((response) => ({ error: null, response }))
    .catch((error) => {
      if (error.response) {
        return { error: error.response };
      }
      return { error };
    });
}

function post(location, body, apiBaseURL = API_BASE) {
  const url = `${apiBaseURL}${location}`;
  return axios
    .post(url, body)
    .then((response) => ({ error: null, response }))
    .catch((error) => {
      if (error.response) {
        return { error: error.response };
      }
      return { error };
    });
}

function put(location, body, apiBaseURL = API_BASE) {
  const url = `${apiBaseURL}${location}`;
  return axios
    .put(url, body)
    .then((response) => ({ error: null, response }))
    .catch((error) => {
      if (error.response) {
        return { error: error.response };
      }
      return { error };
    });
}

function deleteMethod(location, body, apiBaseURL = API_BASE) {
  let url = `${apiBaseURL}${location}`;
  if (body) {
    const qs = querystring.stringify(body);
    if (qs) {
      url += (url.indexOf('?') >= 0 ? '&' : '?') + qs;
    }
  }

  return axios
    .delete(`${url}`)
    .then((response) => ({ error: null, response }))
    .catch((error) => {
      if (error.response) {
        return { error: error.response };
      }
      return { error };
    });
}

function setAuthHeaders(token) {
  if (!token) {
    axios.defaults.headers.common.Authorization = '';
    return;
  }
  axios.defaults.headers.common.Authorization = `Bearer ${token}`;
}

export default {
  post,
  get,
  put,
  delete: deleteMethod,
  setAuthHeaders
};
