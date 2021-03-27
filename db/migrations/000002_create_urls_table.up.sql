CREATE TABLE urls (
  id                serial primary key,
  short_url         text,
  full_url          text,
  hits              int,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON urls
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
