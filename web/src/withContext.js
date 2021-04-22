import { LoginContextProvider } from './context/login';

function withContext(WrappedComponent) {
  return (props) => {
    return (
      <LoginContextProvider>
        <WrappedComponent {...props}/>
      </LoginContextProvider>
    );
  };
}

export default withContext;
