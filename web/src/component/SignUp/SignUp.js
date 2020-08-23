import React from 'react';
import { Link } from 'react-router-dom';

import './SignUp.scss';
import { signUpUser } from '../../actions/auth';
import { connect } from 'react-redux';

function validateEmail(email) {
  // eslint-disable-next-line
  const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(String(email).toLowerCase());
}

function validatePassword(pwd) {
  return pwd.length > 8 && /(?=.*?[0-9])(?=.*?[A-Za-z]).+/gi.test(pwd);
}

const SignUp = ({ signup, signupState }) => {
  const [form, setForm] = React.useState({ username: { value: '', error: '' } });

  return (
    <div className='signup'>
      <div className='label-text'>
        <i>Shorty</i> is free to use for as long as you want.
      </div>
      <div className='label-text'>and with an unlimited number of URL for short.</div>
      <div className='signup-form'>
        <div className='title'>Create an account</div>
        <div className='sub-title'>
          Already have an account? <Link to='/login'>Log In</Link>
        </div>
        {signupState?.error?.message && <div className='err mtb-1 fs-18'>{signupState?.error?.message}</div>}
        <form
          onSubmit={(e) => {
            e.preventDefault();
            let isError = false;
            const username = form?.username?.value || '';
            const email = form?.email?.value || '';
            const password = form?.password?.value || '';
            if (username.length < 4 || username.length > 32) {
              setForm((x) => ({ ...x, username: { ...(x?.username || {}), error: 'Letter & numbers only. 4–32 characters.' } }));
              isError = true;
            }
            if (!validateEmail(email)) {
              isError = true;
              setForm((x) => ({ ...x, email: { ...(x?.email || {}), error: 'Enter a valid email address.' } }));
            }
            if (!validatePassword(password)) {
              isError = true;
              setForm((x) => ({ ...x, password: { ...(x?.password || {}), error: 'Password must be 8 character long with at least one number & one character.' } }));
            }
            if (isError) return null;
            signup({ username, email, password });
            return null;
          }}
        >
          <div className={`input-group ${form?.username?.error ? 'err' : ''}`}>
            <label htmlFor='username'>
              <div>Username</div>
              <input value={form?.username?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, username: { value, error: '' } }))} maxLength={32} minLength={4} type='text' id='username' />
              {form?.username?.error && <span className='err text-left'>{form?.username?.error}</span>}
            </label>
          </div>
          <div className={`input-group ${form?.email?.error ? 'err' : ''}`}>
            <label htmlFor='email'>
              <div>Email</div>
              <input value={form?.email?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, email: { value, error: '' } }))} maxLength={64} minLength={4} type='email' id='email' />
              {form?.email?.error && <span className='err text-left'>{form?.email?.error}</span>}
            </label>
          </div>
          <div className={`input-group ${form?.password?.error ? 'err' : ''}`}>
            <label htmlFor='password'>
              <div>Password</div>
              <input value={form?.password?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, password: { value, error: '' } }))} maxLength={64} minLength={8} type='password' id='password' />
              {form?.password?.error && <span className='err text-left'>{form?.password?.error}</span>}
            </label>
          </div>
          <button type='submit' className='btn-primary'>
            Sign up
          </button>

          <div className='disclaimer'>By creating an account, you agree to </div>
          <div className='disclaimer'>Shorty&apos;s Terms of Service and Privacy Policy.</div>
        </form>
      </div>
    </div>
  );
};

const mapStateToProps = ({ auth }) => ({ signupState: auth?.signup });

export default connect(mapStateToProps, {
  signup: signUpUser
})(SignUp);
