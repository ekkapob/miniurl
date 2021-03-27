CREATE TABLE urls (
  id                  serial PRIMARY KEY,
  short_url           text NOT NULL,
  full_url            text NOT NULL,
  hits                int DEFAULT 0,
  expires_in_seconds  int,
  created_at          timestamp WITH time zone DEFAULT now(),
  last_modified_at    timestamp WITH time zone DEFAULT now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON urls
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
