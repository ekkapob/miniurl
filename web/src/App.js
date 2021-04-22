import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './components/Home';
import withContext from './withContext';

import AdminDashboard from './components/admin/Dashboard';
import Login from './components/Login';
import MustLogin from './components/MustLogin';

import './App.scss';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/admin/dashboard">
          <MustLogin>
            <AdminDashboard></AdminDashboard>
          </MustLogin>
        </Route>

        <Route path="/login">
          <Login/>
        </Route>

        <Route path="/">
          <Home></Home>
        </Route>
      </Switch>
    </Router>
  );
}

export default withContext(App);
