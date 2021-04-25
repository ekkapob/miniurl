import axios from 'axios';
import { useState } from 'react';
import { Link } from 'react-router-dom'

import BigURL from './BigURL';

import './styles/Home.scss';

function Home() {
  const [url, setURL] = useState('');
  const [shortURL, setShortURL] = useState();
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);

  const onKeyDown = (e) => {
    if (e === 'Enter') createURL();
  };

  const onCreateClick = (e) => {
    e.preventDefault();
    createURL();
    e.target.blur();
  };

  const createURL = () => {
    setLoading(true);
    setError();
    setShortURL();

    axios.post('/api/v1/urls', { url })
      .then(res => {
        const { short_url } = res.data;
        setShortURL(short_url);
      })
      .catch(err => {
        const { error } = err.response.data;
        setError(error);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <div className="home">
      <div className="container mt-5 position-relative">
        <div className="logo">
          <h1><span>big</span>URL</h1>
        </div>

        <Link to="/admin/dashboard">
          <div id="admin-dashboard-link" className="d-flex">
            <i className="bi bi-file-person me-1"/>
            <div>admin</div>
          </div>
        </Link>

        <form>
          <div className="mb-3">
            <label htmlFor="urlInput"></label>
            <input type="text" id="urlInput" className="form-control" value={url}
              onChange={(e) => setURL(e.target.value)}
              onKeyDown={onKeyDown}
              placeholder="Enter full domain name e.g. https://google.com"/>
            {
              error &&
                <div className="error mt-2">{error}</div>
            }
          </div>
          <div className="d-grid">
            <button type="submit" className="btn btn-primary"
              onClick={onCreateClick}>Create big URL</button>
          </div>
        </form>
      </div>
      {
        loading &&
          <div className="text-center mt-5">
            <img className="loading"
              src={`${process.env.PUBLIC_URL}/images/loading.gif`} alt="loading" />
          </div>
      }
      {
        shortURL &&
          <BigURL url={shortURL}/>
      }
    </div>
  );
}

export default Home;
