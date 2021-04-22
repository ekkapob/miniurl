import axios from 'axios';
import { useEffect, useState, useRef } from 'react';
import { addSeconds } from 'date-fns'

import Pagination from './Pagination';

import classNames from 'classnames';

import './styles/URLs.scss';

const PAGE_URLS = 20;

function URLs() {
  const [page, setPage] = useState(1);
  const [pageURLs] = useState(PAGE_URLS);
  const [totalPages, setTotalPages] = useState(0);
  const [urls, setURLs] = useState([]);
  const [displayURLs, setDisplayURLs] = useState([]);
  const [orderBy, setOrderBy] = useState('expired_date');
  const [filterText, setFilterText] = useState('');
  const [filterSelect, setFilterSelect] = useState('full_url');

  const [deletingURL, setDeletingURL] = useState()

  const modal = useRef(null);

  useEffect(() => {
    fetchURLs();
  }, [page, orderBy]);

  const fetchURLs = () => {
    setFilterText('');
    setURLs([]);
    setDisplayURLs([]);
    axios.get(`/api/v1/urls?page=${page}&limit=${pageURLs}&orderBy=${orderBy}&orderDirection=desc`)
      .then(res => {
        const { urls, total_pages } = res.data;
        setURLs(urls);
        setDisplayURLs(urls);
        setTotalPages(total_pages);
      })
      .catch(err => {});
  };

  const formatDate = (date) => {
    const d = new Date(date);
    return new Intl.DateTimeFormat('en-GB',{
      year: 'numeric', month: 'numeric', day: 'numeric',
      hour: 'numeric', minute: 'numeric', second: 'numeric',
      hour12: false,
      timeZone: 'Asia/Bangkok',
    }).format(d);
  };

  const expiredDateTime = (createdDateStr, expireInSeconds) => {
    const date = new Date(createdDateStr);
    return formatDate(addSeconds(date, expireInSeconds));
  };

  const onFilterChange = (e) => {
    const value = e.target.value;
    setFilterText(value);
    const regExp = new RegExp(value, 'g');
    const u = [...urls];
    const matchedURLs = u.filter(url => (
      regExp.test(url[filterSelect])
    ));
    setDisplayURLs(matchedURLs);
  };

  const onDeleteURLClick = (url) => {
    return (e) => {
      setDeletingURL(url);
      showModal();
    };
  };

  const onConfirmDeleteURL = (deletingUrl) => {
    return (e) => {
      axios.delete(`/api/v1/urls/${deletingUrl.id}`)
        .then(res => {
          setDeletingURL();
          fetchURLs();
        })
        .catch(err => {})
        .finally(() => setDeletingURL());
    };
  };

  const showModal = () => {
    const deleteModal = new window.bootstrap.Modal(modal.current);
    deleteModal.show();
  };

  const onFilterSelectChange = (e) => {
    setFilterSelect(e.target.value);
  };

  const onPageLinkClick = (page) => {
    setPage(page);
  };

  return (
    <div className="admin-urls">
      <div className="d-flex justify-content-between">
        <div>
          Ordered by:
          <span className={
            classNames('badge mx-2 bg-secondary',
              {'bg-danger': orderBy === 'expired_date'})}
              onClick={() => {setOrderBy('expired_date')}}>
            Expired date
          </span>
          <span className={
            classNames('badge bg-secondary',
              {'bg-danger': orderBy === 'hits'})}
              onClick={() => {setOrderBy('hits')}}>
            Hits
          </span>
        </div>
        <div className="ms-2">
          URLs per page: <span className="badge bg-light">{pageURLs}</span>
        </div>
      </div>
      <div className="input-group my-3">
        <label className="input-group-text" htmlFor="filterURLInput">
          Filter
        </label>
      <select className="form-select" aria-label="Default select example"
        onChange={onFilterSelectChange} value={filterSelect}>
        <option value="full_url">Full URL</option>
        <option value="short_url">Short URL</option>
      </select>
        <input id="filterURLInput" className="form-control"
          type="text" onChange={onFilterChange} value={filterText}/>
      </div>


      <table className="table mt-2">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Short URL</th>
            <th scope="col">Full URL</th>
            <th scope="col">Hits</th>
            <th scope="col">Expired At</th>
            <th scope="col">Created At</th>
            <th scope="col">Actions</th>
          </tr>
        </thead>
        <tbody>
          {
            displayURLs.map((v,k) => (
              <tr key={k}>
                <th scope="row">{v.id}</th>
                <td>{v.short_url}</td>
                <td>{v.full_url}</td>
                <td>{v.hits}</td>
                <td>{formatDate(v.created_at)}</td>
                <td>
                  {expiredDateTime(v.created_at, v.expires_in_seconds)}
                </td>
                <td>
                  <button className="btn btn-danger btn-sm"
                    onClick={onDeleteURLClick(v)}>
                    Delete
                  </button>
                </td>
              </tr>
            ))
          }
        </tbody>
      </table>

      <div className="my-5">
        <Pagination currentPage={page} totalPages={totalPages}
          onPageLinkClick={onPageLinkClick}/>
      </div>


      <div ref={modal} id="deleteURLModal" className="modal fade" tabIndex="-1"
        aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              <h5 className="modal-title" id="exampleModalLabel">Are you sure to delete a URL</h5>
              <button type="button" className="btn-close" data-bs-dismiss="modal"
                aria-label="Close"></button>
            </div>
            { deletingURL &&
                <div className="modal-body">
                  Short URL: <span className="badge bg-danger">
                    {deletingURL.short_url}</span> ({deletingURL.full_url})
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
  );
}

export default URLs;
