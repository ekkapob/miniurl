import axios from 'axios';
import { useContext, useEffect, useRef, useState } from 'react';
import { formatDate } from './../../date';

import { LoginContext } from './../../context/login';

function BlacklistURLs() {
  const modal = useRef(null);
  const [urls, setURLs] = useState([]);
  const [displayURLs, setDisplayURLs] = useState([]);
  const [addURLText, setAddURLText] = useState('');
  const [addURLError, setAddURLError] = useState(false);
  const [filterText, setFilterText] = useState('');
  const [deletingURL, setDeletingURL] = useState();
  const [login] = useContext(LoginContext);

  useEffect(() => {
    fetchURLs();
  }, []);

  const fetchURLs = () => {
    axios.get('/api/v1/blacklist_urls', {
      headers: {
        Authorization: `Basic ${login.basicAuth}`,
      },
    })
      .then(res => {
        const { urls } = res.data;
        setURLs(urls);
        setDisplayURLs(urls);
      })
      .catch(err => {});
  };

  const onDeleteURLClick = (url) => {
    return (e) => {
      setDeletingURL(url);
      showModal();
    };
  };

  const showModal = () => {
    const deleteModal = new window.bootstrap.Modal(modal.current);
    deleteModal.show();
  };

  const onConfirmDeleteURL = (url) => {
    return (e) => {
      axios.delete(`/api/v1/blacklist_urls/${url.id}`, {
        headers: {
          Authorization: `Basic ${login.basicAuth}`,
        },
      })
        .then(res => {
          setDeletingURL();
          fetchURLs();
        })
        .catch(err => {})
        .finally(() => setDeletingURL());
    };
  };

  const onFilterChange = (e) => {
    const value = e.target.value;
    setFilterText(value);
    const regExp = new RegExp(value, 'g');
    const u = [...urls];
    const matchedURLs = u.filter(url => (
      regExp.test(url.url)
    ));
    setDisplayURLs(matchedURLs);
  };

  const onAddURLChange = (e) => {
    setAddURLText(e.target.value);
  };

  const onBlacklistURLCreate = () => {
    setAddURLError(false);
    axios.post('/api/v1/blacklist_urls', {
      url: addURLText,
    }, {
      headers: {
        Authorization: `Basic ${login.basicAuth}`,
      },
    })
      .then(res => {
        setAddURLText('');
        setAddURLError(false);
        fetchURLs();
      })
      .catch(err => {
        setAddURLError(true);
      });
  };

  return (
    <div className="blacklist_urls">
      <div className="input-group my-3">
        <label className="input-group-text" htmlFor="addURLInput">
          Add Blacklist URL
        </label>
        <input id="addURLInput" className="form-control"
          type="text" onChange={onAddURLChange} value={addURLText}/>
        <button type="submit" className="btn btn-primary"
          onClick={onBlacklistURLCreate}>Create</button>
      </div>
      {
        addURLError &&
          <div className="error">Enter valid full URL e.g. https://www.google.com</div>
      }

      <div className="input-group my-3">
        <label className="input-group-text" htmlFor="filterURLInput">
          Filter URLs
        </label>
        <input id="filterURLInput" className="form-control"
          type="text" onChange={onFilterChange} value={filterText}/>
      </div>

      <table className="table mt-2">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">URL</th>
            <th scope="col">Created At</th>
            <th scope="col">Actions</th>
          </tr>
        </thead>
        <tbody>
          {
            displayURLs.map((v,k) => (
              <tr key={k}>
                <th scope="row">{v.id}</th>
                <td>{v.url}</td>
                <td>{formatDate(v.created_at)}</td>
                <td>
                  <button className="btn btn-danger btn-sm"
                    onClick={onDeleteURLClick(v)}>
                    delete
                  </button>
                </td>
              </tr>
            ))
          }
        </tbody>
      </table>

      <div ref={modal} id="deleteURLModal" className="modal fade" tabIndex="-1"
        aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              <h5 className="modal-title" id="exampleModalLabel">Are you sure to delete a URL?</h5>
              <button type="button" className="btn-close" data-bs-dismiss="modal"
                aria-label="Close"></button>
            </div>
            { deletingURL &&
                <div className="modal-body">
                  Confirm to delete URL "{deletingURL.url}"
                </div>
            }
            <div className="modal-footer">
              <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">
                Close
              </button>
              <button type="button" className="btn btn-danger"
                data-bs-dismiss="modal"
                onClick={onConfirmDeleteURL(deletingURL)}>
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>

    </div>
  )
}

export default BlacklistURLs;
