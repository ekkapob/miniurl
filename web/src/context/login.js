import { useReducer, createContext } from 'react';

export const LoginContext = createContext();

const initialState = {
  status: false,
};

const UPDATE_LOGIN = 'UPDATE_LOGIN';
const RESET = 'RESET';

const reducer = (state, action) => {
  switch (action.type) {
    case UPDATE_LOGIN:
      return { ...state, ...action.payload };
    case RESET:
      return { ...initialState };
    default:
      return state;
  }
};

export const LoginContextProvider = (props) => {
  const [state, dispatch] = useReducer(reducer, initialState);
  const contextValue = [state, dispatch];

  return (
    <LoginContext.Provider value={ contextValue }>
      { props.children }
    </LoginContext.Provider>
  );
};

