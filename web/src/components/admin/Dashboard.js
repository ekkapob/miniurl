import { Link } from "react-router-dom";
import URLs from "./URLs";

import './styles/Dashboard.scss';

function Dashboard() {
  return (
    <div id="admin-dashboard" className="container mt-3">
      <Link to="/">Home</Link>
      <h3 className="mt-3">Dashboard</h3>
      <hr />
      <div className="d-flex">
        <div className="menu-sidebar d-flex">
          <ul className="menus me-5 pt-3">
            <li>
              <Link to="/admin/dashboard" className="active">
                URLs
              </Link>
            </li>
            {/*
              <li>
                <Link to="/admin/blacklist">
                  Blacklist
                </Link>
              </li>
            */}
          </ul>
        </div>
        <div className="content flex-fill pt-3 ps-3">
          <URLs/>
        </div>

      </div>
    </div>
  );
}

export default Dashboard;
