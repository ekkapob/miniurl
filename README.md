# BigURL [http://bigurl.tk](http://bigurl.tk)
Shortest URL Ever!

For Admin, use
- Account: `admin`
- Password: `admin`

Notes: Success login is saved to local storage.

## To Run
```sh
$ make docker-up

# Only first time running the service
# 1) Wait unitl Postgres is ready
# 2) Run migrations

$ make docker-migrate
```


## Redirect with Short URL

```sh
http://bigurl.tk/urls/{short_url}
```

Results:
 1. Redirect to registered URL with status 302
 2. Return 401 Status when cannot find the registered URL

## Postman Collection
[Postman Collection](https://raw.githubusercontent.com/ekkapob/miniurl/main/MiniURL.postman_collection.json)

Environments
- `host`: `http://bigurl.tk`

## API Usage
### 1) Create a short URL
**[POST] http://bigurl.tk/api/v1/urls**

JSON Request:
```sh
{
    "url": "https://www.cnn.com"
}
```
Response Statuses:
  - [201] Successfully created

```sh
{
    "short_url": "b",
    "full_url": "https://www.cnn.com"
}
```

  - [400] on bad requests: invalid URL, blacklisted URL

```sh
{
    "error": "URL is not valid. Please provide a full URL e.g. https://google.com"
}
```

```sh
{
    "error": "This URL is in blacklist."
}
```

### 2) Get a short URL
**[GET] http://bigurl.tk/api/v1/urls/{short_url}**

JSON Response:
```sh
{
    "short_url": "b",
    "full_url": "https://www.cnn.com"
}
```
### 3) Get URLs with sorting and pagination
**[GET] http://bigurl.tk/api/v1/urls?{options}**

Options
| names | values | descriptions |
| ----- | ------ | ------ |
| page  | _integer_ e.g. 1, 2, 3 | page number |
| limit | _integer_ e.g. 10, 20, 20 | no. of items per request
| orderBy  | _column names_ e.g. expired_date | column to be sorted |
| orderDirection | _asc_ or _desc_ | sorting directions |

Examples:
- /api/v1/urls?page=5&limit=5
- /api/v1/urls?orderBy=expired_date&orderDirection=desc

JSON Response:
```sh
{
    "page": 5,
    "total_pages": 5,
    "urls": [
        {
            "id": 5,
            "short_url": "f",
            "full_url": "https://www.cnn.com",
            "hits": 0,
            "created_at": "2021-03-30T06:47:03.5104Z",
            "expires_in_seconds": 604800,
            "last_modified_at": "2021-03-30T06:47:03.5104Z"
        },
        {
            "id": 4,
            "short_url": "e",
            "full_url": "https://www.cnn.com",
            "hits": 0,
            "created_at": "2021-03-30T06:47:03.051716Z",
            "expires_in_seconds": 604800,
            "last_modified_at": "2021-03-30T06:47:03.051716Z"
        }
    ]
}
```

### 4) Delete a URL
**[DELETE] http://bigurl.tk/api/v1/urls/{url_id}**

Response Statuses:
  - [200] OK
  - [400] When cannot find the URL to be deleted

### 5) Create a blacklist URL
**[POST] http://bigurl.tk/api/v1/blacklist_urls**

JSON Request:
```sh
{
    "url": "https://www.cnn.com"
}
```

Response Statuses:
  - [201] OK
  - [400] on bad request

```sh
{
    "error": "URL is not valid. Please provide a full URL e.g. https://google.com"
}
```

## TODOs
1. Secured all APIs with OAuth
