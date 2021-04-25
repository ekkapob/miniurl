import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './components/Home';
import withContext from './withContext';

import AdminDashboard from './components/admin/Dashboard';
import Blacklist from './components/admin/Blacklist';
import MustLogin from './components/MustLogin';
import Login from './components/Login';

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

        <Route path="/admin/blacklist">
          <MustLogin>
            <Blacklist></Blacklist>
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
