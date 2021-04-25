import { Link } from 'react-router-dom';

import DashboardMenu from './DashboardMenu';
import BlacklistURLs from './BlacklistURLs';

function Blacklist() {
  return(
    <div id="admin-dashboard" className="container mt-3">
      <Link to="/">Home</Link>
      <h3 className="mt-3">Dashboard</h3>
      <hr />
      <div className="d-flex">
        <DashboardMenu active="Blacklist"/>
        <div className="content flex-fill pt-3 ps-3">
          <BlacklistURLs/>
        </div>
      </div>
    </div>
  );
}

export default Blacklist;
