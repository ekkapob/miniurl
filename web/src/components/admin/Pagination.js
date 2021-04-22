import { useEffect, useState } from 'react';

import classNames from 'classnames';

import './styles/Pagination.scss';

const PAGE_LINKS_PADDING = 4;

function Pagination(props) {
  const {
    currentPage,
    totalPages,
    onPageLinkClick,
  } = props;
  const [pageLinks, setPageLinks] = useState([]);


  useEffect(() => {
    createPageLinks();
  }, [currentPage, totalPages]);

  const createPageLinks = () => {
    if (!totalPages || totalPages.length === 0) return;

    let links = [];
    links.push(
      <PageLink key="previous" currentPage={currentPage}
        totalPages={totalPages} text="Previous"
        disabled={ currentPage === 1 }
        onPageLinkClick={() => onPageLinkClick(currentPage - 1)}
      />
    );

    links = links.concat(leftPageLinks(currentPage));
    links = links.concat(rightPageLinks(currentPage, totalPages));

    links.push(
      <PageLink key="next" currentPage={currentPage}
        totalPages={totalPages} text="Next"
        disabled={ currentPage === totalPages }
        onPageLinkClick={() => onPageLinkClick(currentPage + 1)}
      />
    );

    setPageLinks(links);
  };

  const leftPageLinks = (currentPage) => {
    let links = [];
    let cursor = currentPage;

    while (cursor > 0 && cursor > currentPage - PAGE_LINKS_PADDING) {
      const c = cursor;
      links.push(
        <PageLink key={cursor} cursor={cursor}
          currentPage={currentPage} text={cursor}
          onPageLinkClick={() => onPageLinkClick(c)}/>
      );
      cursor -= 1;
    }
    if (cursor >= 1) {
      links.push(
        <PageLink key="left-dot" text="..."/>
      )
    }
    return links.reverse();
  };

  const rightPageLinks = (currentPage, totalPages) => {
    let links = [];
    let cursor = currentPage + 1;

    while (cursor <= totalPages && cursor < currentPage + PAGE_LINKS_PADDING) {
      const c = cursor;
      links.push(
        <PageLink key={cursor} cursor={cursor}
          currentPage={currentPage} text={cursor}
          onPageLinkClick={() => onPageLinkClick(c)}/>
      );
      cursor += 1;
    }
    if (cursor <= totalPages) {
      links.push(
        <PageLink key="right-dot" text="..."/>
      )
    }
    return links;
  };

  return (
    <div className="pagination">
      {pageLinks}
    </div>
  );
}

function PageLink(props) {
  const {
    cursor = -1,
    currentPage,
    text,
    disabled = false,
    onPageLinkClick = () => {},
  } = props;

  return (
    <li className={classNames(
      'page-item',
      { disabled },
      { active: currentPage === cursor },
    )}>
      <span className="page-link"
        onClick={onPageLinkClick}>{text}</span>
    </li>
  )
}

export default Pagination;
