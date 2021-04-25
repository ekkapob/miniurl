import { Link } from 'react-router-dom';
import classNames from 'classnames';

function DashboardMenu(props) {

  const { active = 'URLs' } = props;

  return (
    <div className="menu-sidebar d-flex">
      <ul className="menus me-5 pt-3">
        <li>
          <MenuLink to="/admin/dashboard" text="URLs" active={active}/>
        </li>
        <li>
          <MenuLink to="/admin/blacklist" text="Blacklist" active={active}/>
        </li>
      </ul>
    </div>
  );
}

function MenuLink(props) {
  const { to, text, active } = props;

  return (
    <Link to={to}
      className={classNames({active: active === text})}>
      {text}
    </Link>
  )
}

export default DashboardMenu;
