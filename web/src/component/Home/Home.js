import React from 'react';
import './Home.scss';

const validateURL = (url) => {
  // eslint-disable-next-line
  const regex = /^(http[s]?:\/\/(www\.)?|ftp:\/\/(www\.)?|www\.|mailto:|gopher:\/\/){1}([0-9A-Za-z-\.@:%_\+~#=]+)+((\.[a-zA-Z]{2,3})+)(\/(.)*)?(\?(.)*)?/gi;
  return regex.test(url.toLowerCase());
};

const Home = () => {
  const [inputUrl, setURL] = React.useState({ value: '' });

  const handleSubmit = React.useCallback(
    (event) => {
      event.preventDefault();
      const { value } = inputUrl;
      // validate url
      if (!validateURL(value)) {
        setURL((x) => ({ ...x, error: 'Enter a valid URL' }));
        return null;
      }
      // API call
      return null;
    },
    [inputUrl]
  );

  const isError = Boolean(inputUrl?.error);

  return (
    <div className='main'>
      <div className='main-logo' />
      <div className='main-logo-title'>
        Shorty
        <span className='main-logo-subtitle'>URL Shortener</span>
      </div>
      <form onSubmit={handleSubmit} className={`main-form ${isError ? 'err' : ''}`}>
        {/* eslint-disable-next-line jsx-a11y/no-autofocus */}
        <input value={inputUrl?.value || ''} onChange={({ target: { value } }) => setURL({ value })} autoFocus type='url' placeholder='Very Long URL' className='main-input' />
        {isError && <div className='err mt-1'>{inputUrl?.error}</div>}
        <button className='submit' type='submit'>
          Shorten{' '}
          <span role='img' aria-label='ok'>
            👌
          </span>
        </button>
      </form>
    </div>
  );
};

export default Home;
