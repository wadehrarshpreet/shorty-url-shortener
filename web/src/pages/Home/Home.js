/* eslint-disable operator-linebreak */
import React from 'react';
import { connect } from 'react-redux';
import QRCode from 'qrcode.react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import './Home.scss';
import AppLoader from '../../component/Loader/Loader.js';
import { shortenURL } from '../../actions/short';
import Modal from '../../component/Modal/Modal.js';

const validateURL = (url) => {
  // eslint-disable-next-line
  const regex = /^(http[s]?:\/\/(www\.)?|ftp:\/\/(www\.)?|www\.|mailto:|gopher:\/\/){1}([0-9A-Za-z-\.@:%_\+~#=]+)+((\.[a-zA-Z]{2,3})+)(\/(.)*)?(\?(.)*)?/gi;
  return regex.test(url.toLowerCase());
};

const Home = ({ userData, shortLongURL, shortUrl }) => {
  const baseURL =
    API_BASE || `${window?.location?.protocol}//${window?.location?.host}`;

  const isLoggedIn = userData?.token;
  const [showCustomURLInput, setCustomURLInputState] = React.useState(false);
  const [inputUrl, setURL] = React.useState({ value: '' });
  const [customInputUrl, setCustomURL] = React.useState({ value: '' });
  const [result, setResult] = React.useState(false);
  const [showLoader, setLoader] = React.useState(false);
  const loaderRef = React.useRef(null);
  const inputRef = React.useRef(null);
  const shortURLRef = React.useRef(null);

  const resetForm = React.useCallback(() => {
    setURL({ value: '' });
    setResult(false);
    if (typeof inputRef?.current?.focus === 'function') {
      inputRef.current.focus();
    }
  });

  React.useEffect(() => {
    // wait 500ms before starting loader
    if (shortUrl.fetching) {
      loaderRef.current = setTimeout(() => {
        setLoader(true);
      }, 500);
    } else {
      clearTimeout(loaderRef.current);
      setLoader(false);
    }
  }, [shortUrl.fetching]);

  React.useEffect(() => {
    const url = inputUrl?.value;
    const custom = customInputUrl?.value || '';
    const requestURL = `${url}${custom ? `--${custom}` : ''}`;
    if (
      shortUrl?.requestURL === requestURL &&
      shortUrl?.data?.constructor === Map &&
      shortUrl.data.has(requestURL)
    ) {
      const shortCode = shortUrl.data.get(requestURL);
      setResult({
        long: url,
        short: `${baseURL}/${shortCode}`
      });
    }
  }, [shortUrl]);

  const copyShortURL = React.useCallback(() => {
    try {
      // Select Text
      if (document.selection) {
        // IE
        const range = document.body.createTextRange();
        range.moveToElementText(shortURLRef.current);
        range.select();
      } else if (window.getSelection) {
        const range = document.createRange();
        range.selectNode(shortURLRef.current);
        window.getSelection().removeAllRanges();
        window.getSelection().addRange(range);
      }
      // Copy Text IE 11 + support
      document.execCommand('copy');
    } catch (e) {
      // e
    }
  }, []);

  const handleSubmit = React.useCallback(
    (event) => {
      event.preventDefault();
      const { value } = inputUrl;
      // validate url
      if (!validateURL(value)) {
        setURL((x) => ({ ...x, error: 'Enter a valid URL' }));
        return null;
      }
      // validate custom URL alphanumeric&_ & min 4 length
      if (
        isLoggedIn &&
        customInputUrl?.value &&
        !/[a-zA-Z0-9_]{4,}/gi.test(customInputUrl?.value || '')
      ) {
        setCustomURL((x) => ({
          ...x,
          error:
            'ShortURL must be at least 4 characters long and contain only alphanumeric & underscore'
        }));
        return null;
      }
      // API call
      shortLongURL({ url: value, custom: customInputUrl?.value });
      return null;
    },
    [inputUrl, customInputUrl]
  );

  const isError = Boolean(inputUrl?.error);

  return (
    <>
      {showLoader && <AppLoader />}
      {
        <Modal
          className='result-modal'
          showState={result}
          onClose={() => setResult(false)}
        >
          <div className='qr-code'>
            {result?.short && (
              <QRCode
                size={150}
                bgColor={'#ffffff'}
                fgColor={'#000'}
                value={result?.short}
                img={{
                  src: '/assets/logo.png',
                  top: 50,
                  left: 50,
                  width: 20,
                  height: 20
                }}
              />
            )}
          </div>

          <div onClick={copyShortURL} className='result-box'>
            <div ref={shortURLRef}>{result?.short}</div>
            <div className='copy-icon'>
              <img alt='copy' src='/assets/copy.svg' />
            </div>
          </div>
          <div className='long-url'>
            Redirect to:{' '}
            <span>
              <a href={result?.long} target='_blank' rel='noreferrer'>
                {result?.long}
              </a>
            </span>
          </div>
          <div className='short-again'>
            <div onClick={resetForm} className='main-logo m-auto' />
            <div onClick={resetForm}>Short another URL</div>
          </div>
        </Modal>
      }
      <div className='main'>
        <div className='main-logo' />
        <div className='main-logo-title'>
          Shorty
          <span className='main-logo-subtitle'>URL Shortener</span>
        </div>

        <form onSubmit={handleSubmit} className={'main-form'}>
          {isError && <div className='err mt-1'>{inputUrl?.error}</div>}
          <div className={`${isError ? 'err' : ''}`}>
            <input
              autoComplete='url'
              value={inputUrl?.value || ''}
              onChange={({ target: { value } }) => setURL({ value })}
              // eslint-disable-next-line jsx-a11y/no-autofocus
              autoFocus
              type='url'
              placeholder='Very Long URL'
              className='main-input'
              ref={inputRef}
            />
          </div>
          {isLoggedIn && (
            <>
              <div
                className={`${
                  showCustomURLInput ? 'expand' : ''
                } custom-url-trigger`}
                onClick={() => setCustomURLInputState((x) => !x)}
              >
                Custom URL <ExpandMoreIcon />{' '}
              </div>
              <div
                className={`${showCustomURLInput ? 'expand' : ''} ${
                  customInputUrl?.error ? 'err' : ''
                } custom-container`}
              >
                {baseURL}/
                <input
                  value={customInputUrl.value}
                  onChange={({ target: { value } }) => setCustomURL({ value })}
                  maxLength={32}
                  minLength={4}
                  className='main-input'
                  type='text'
                ></input>
              </div>
              {customInputUrl?.error && (
                <div className='err'>{customInputUrl?.error}</div>
              )}
            </>
          )}

          <button className='submit' type='submit'>
            Shorten{' '}
            <span role='img' aria-label='ok'>
              👌
            </span>
          </button>
          {shortUrl?.error?.error && inputUrl?.value && (
            <div className='err mt-1'>{shortUrl?.error?.message}</div>
          )}
        </form>
      </div>
    </>
  );
};

const mapStateToProps = ({ auth, shortUrl }) => ({
  userData: auth?.data,
  shortUrl
});

export default connect(mapStateToProps, { shortLongURL: shortenURL })(Home);
