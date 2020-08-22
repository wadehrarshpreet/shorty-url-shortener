import React from 'react';
import { Link } from 'react-router-dom';

const Login = () => {
  const [form, setForm] = React.useState({ username: { value: '', error: '' } });
  console.log(form);

  return (
    <div className='signup'>
      <div className='signup-form'>
        <div className='title'>Login and start shortening</div>
        <div className='sub-title'>
          Don&apos;t have an account? <Link to='/signup'>Sign Up</Link>
        </div>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            const username = form?.username?.value || '';
            const password = form?.password?.value || '';
            console.log(username);
            console.log(password);
            return null;
          }}
        >
          <div className={`input-group ${form?.username?.error ? 'err' : ''}`}>
            <label htmlFor='username'>
              <div>Username or Email</div>
              <input value={form?.username?.value || ''} onChange={({ target: { value } }) => setForm((x) => ({ ...x, username: { value, error: '' } }))} maxLength={32} minLength={4} type='text' id='username' />
              {form?.username?.error && <span className='err text-left'>{form?.username?.error}</span>}
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
            Log In
          </button>
          {/* <div className='disclaimer'>Forgot Password?</div> */}
        </form>
      </div>
    </div>
  );
};

export default Login;
