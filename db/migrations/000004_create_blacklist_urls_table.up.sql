CREATE TABLE blacklist_urls (
  id                  serial PRIMARY KEY,
  url                 text NOT NULL,
  created_at          timestamp WITH time zone DEFAULT now()
);

CREATE UNIQUE INDEX blacklist_url
ON blacklist_urls (url, id);
