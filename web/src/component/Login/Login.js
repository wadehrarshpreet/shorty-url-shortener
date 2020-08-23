import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { loginUser } from '../../actions/auth';

const Login = ({ login, loginState }) => {
  const [form, setForm] = React.useState({ username: { value: '', error: '' } });
  console.log(loginState);

  return (
    <div className='signup'>
      <div className='signup-form'>
        <div className='title'>Login and start shortening</div>
        <div className='sub-title'>
          Don&apos;t have an account? <Link to='/signup'>Sign Up</Link>
        </div>
        {loginState?.error?.message && <div className='err mtb-1 fs-18'>{loginState?.error?.message}</div>}
        <form
          onSubmit={(e) => {
            e.preventDefault();
            const username = form?.username?.value || '';
            const password = form?.password?.value || '';
            if (!username) {
              setForm((x) => ({ ...x, username: { ...(x?.username || {}), error: 'Enter Username or Email' } }));
              return null;
            }
            if (!password) {
              setForm((x) => ({ ...x, password: { ...(x?.password || {}), error: 'Enter Password' } }));
              return null;
            }
            login({ username, password });
            return null;
          }}
        >
          <div className={`input-group ${form?.username?.error ? 'err' : ''}`}>
            <label htmlFor='username'>
              <div>Username</div>
              <input value={form?.username?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, username: { value, error: '' } }))} maxLength={32} type='text' id='username' />
              {form?.username?.error && <span className='err text-left'>{form?.username?.error}</span>}
            </label>
          </div>
          <div className={`input-group ${form?.password?.error ? 'err' : ''}`}>
            <label htmlFor='password'>
              <div>Password</div>
              <input value={form?.password?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, password: { value, error: '' } }))} maxLength={64} type='password' id='password' />
              {form?.password?.error && <span className='err text-left'>{form?.password?.error}</span>}
            </label>
          </div>
          <button type='submit' className='btn-primary'>
            Log In
          </button>
          {/* <div className='disclaimer'>Forgot Password?</div> */}
        </form>
      </div>
    </div>
  );
};

const mapStateToProps = ({ auth }) => ({ loginState: auth?.login });

export default connect(mapStateToProps, {
  login: loginUser
})(Login);
