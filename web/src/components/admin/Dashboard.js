import { Link } from 'react-router-dom';

import DashboardMenu from './DashboardMenu';
import URLs from "./URLs";

import './styles/Dashboard.scss';

function Dashboard() {
  return (
    <div id="admin-dashboard" className="container mt-3">
      <Link to="/">Home</Link>
      <h3 className="mt-3">Dashboard</h3>
      <hr />
      <div className="d-flex">
        <DashboardMenu active="URLs"/>
        <div className="content flex-fill pt-3 ps-3">
          <URLs/>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
