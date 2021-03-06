import axios from 'axios';

import { useContext, useRef, useState } from 'react';
import { withRouter, Redirect } from 'react-router-dom';
import { encrypt } from './../crypto';

import qs from 'qs';

import { LoginContext } from './../context/login';

function Login(props) {
  const account = useRef();
  const password = useRef();
  const [error, setError] = useState(false);
  const [login, dispatchLogin] = useContext(LoginContext);

  const search = qs.parse(
    props.location.search,
    { ignoreQueryPrefix: true },
  );
  const redirectURL = search.redirect || '/';

  const doLogin = () => {
    setError(false);

    axios.post('/api/v1/auth', {
      account: account.current.value,
      password: password.current.value,
    })
      .then(res => {
        const { basic_auth } = res.data;

        const auth = encrypt(JSON.stringify({
          login: true,
          basicAuth: basic_auth,
        }));
        localStorage.setItem('auth', auth);

        dispatchLogin({
          type: 'UPDATE_LOGIN',
          payload: {
            status: true,
            basicAuth: basic_auth,
          },
        });
      })
      .catch(err => {
        setError(true);
      });
  };

  const onSubmit = () => {
    doLogin();
  };

  const onPasswordKeyDown = (e) => {
    if (e.key === 'Enter') {
      onSubmit();
    }
  };

  return (
    <>
    {
      login.status
        ? <Redirect to={redirectURL}/>
        : <div className="container-fluid mt-5">
            <div className="mb-3">
              <label htmlFor="accountInput"
                className="form-label">Account</label>
              <input ref={account} type="text"
                className="form-control" id="accountInput"/>
            </div>
            <div className="mb-3">
              <label htmlFor="passwordInput"
                className="form-label">Password</label>
              <input ref={password} type="password"
                className="form-control" id="passwordInput"
                onKeyDown={onPasswordKeyDown}/>
            </div>
            <button className="btn btn-primary"
              onClick={onSubmit}>Submit</button>
            {
              error &&
                <div className="mt-2" style={{color: 'red'}}>
                  Incorrect Account or Password
                </div>
            }
          </div>
    }
    </>
  )
}
export default withRouter(Login);

