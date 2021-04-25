import { useContext, useEffect } from 'react';
import { Redirect, withRouter } from 'react-router-dom';
import { LoginContext } from './../context/login';
import { decrypt } from './../crypto';

function MustLogin(props) {
  const [login, dispatchLogin] = useContext(LoginContext);
  const redirectURL = props.location.pathname;

  useEffect(() => {
    const { login, basicAuth } = checkLocalAuth();
    if (!login) return;

    dispatchLogin({
      type: 'UPDATE_LOGIN',
      payload: { status: login, basicAuth },
    });

  }, []);

  const checkLocalAuth = () => {
    const auth = localStorage.getItem('auth');
    const decryptedAuth = decrypt(auth);
    if (!decryptedAuth) return false;
    try {
      const json = JSON.parse(decryptedAuth) || { login: false };
      const { login, basicAuth } = json;
      return { login, basicAuth };
    } catch (err) {
      return { login: false };
    }
  };

  return (
    <>
    {
      login.status
        ? <div>{props.children}</div>
        : <Redirect
            to={{
              pathname: '/login',
              search: `?redirect=${redirectURL}`,
            }}/>
    }
    </>
  )
}
export default withRouter(MustLogin);

