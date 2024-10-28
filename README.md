# Music Catalog API

This is a simple REST API for showing music. It allows client to search, like/unlike, and get music recommendations.

## Authentication

All endpoints require authentication via a Bearer Token. Include the token in the Authorization header of your requests:

```
Authorization: Bearer <your_token_here>
```

## Endpoints Overview

| Method | Endpoint                | Description                                                 |
| ------ | ----------------------- | ----------------------------------------------------------- |
| POST   | `/api/register`         | Create a new user by providing required details             |
| POST   | `/api/login`            | Authenticate a user to receive an access token              |
| GET    | `/api/search`           | Search for music based on query parameters                  |
| GET    | `/api/recommendations`  | Show music recommendations based on similar music           |
| POST   | `/api/track_activities` | Like or unlike music from search results or recommendations |

## Endpoints

### User Register

<details>
<summary>Create a new user by providing required details</summary>
Request:

- Method: `POST`
- URL: `/api/register`
- Headers:
  - `Content-Type`: `application/json`
- Request Body:
  - email `string`: The user's email address, which must be valid and unique.
  - username `string`: The chosen username for the account.
  - password `string`: The user's password.
- On Success Status Code: `201`
- Request Example:

  ```json
  {
    "email": "test@gmail.com",
    "username": "test",
    "password": "rahasia"
  }
  ```

- Response Example:

  ```json
  {
    "id": "ce2ed257-4f43-4405-bf8c-f732a945fc59",
    "email": "test@gmail.com",
    "username": "test",
    "created_at": "2024-10-28T13:40:13.571283Z",
    "updated_at": "2024-10-28T13:40:13.571283Z"
  }
  ```

  </details>

### User Login

<details>
<summary>Authenticate a user to receive an access token</summary>
Request:

- Method: `POST`
- URL: `/api/login`
- Headers:
  - `Content-Type`: `application/json`
- Request Body:
  - email `string`: The user's registered email address.
  - password `string`: The user's password.
- On Success Status Code: `200`
- Request Example:

  ```json
  {
    "email": "test@gmail.com",
    "password": "rahasia"
  }
  ```

- Response Example:

  ```json
  {
    "id": "ce2ed257-4f43-4405-bf8c-f732a945fc59",
    "email": "test@gmail.com",
    "username": "test",
    "created_at": "2024-10-28T13:40:13.571283Z",
    "updated_at": "2024-10-28T13:40:13.571283Z",
    "token": "<your_token_here>"
  }
  ```

  </details>

### Music Search

<details>
<summary>Search for music based on query parameters</summary>
Request:

- Method: `GET`
- URL: `/api/search`
- Headers:
  - `Authorization`: `Bearer <your_token_here>`
- Query Parameters:
  - query `string` `required`: Search musics by title.
  - page_size `integer` `optional`: The number of musics to retrieve per page (default is 10).
  - page_index `integer` `optional`: The page number to retrieve (default is 1).
- On Success Status Code: `200`
- Request Example:

  ```bash
  curl --request GET \
  --url 'http://localhost:8081/api/search?query=sialan&page_size=1&page_index=2' \
  --header 'Authorization: Bearer <your_token_here>'
  ```

- Response Example:

  ```json
  {
    "limit": 1,
    "offset": 1,
    "total": 118,
    "items": [
      {
        "album_type": "single",
        "album_total_tracks": 1,
        "album_images": [
          "https://i.scdn.co/image/ab67616d0000b273daa30945bef3b4ace8721d30",
          "https://i.scdn.co/image/ab67616d00001e02daa30945bef3b4ace8721d30",
          "https://i.scdn.co/image/ab67616d00004851daa30945bef3b4ace8721d30"
        ],
        "album_name": "Sialan",
        "album_release_date": "2024-10-10",
        "artists_name": ["Second Day"],
        "explicit": false,
        "id": "5gniat5rhCboNlDMlA1pYY",
        "name": "Sialan",
        "is_liked": false
      }
    ]
  }
  ```

  </details>

### Music Recommendations

<details>
<summary>Show music recommendations based on similar music</summary>
Request:

- Method: `GET`
- URL: `/api/recommendations`
- Headers:
  - `Authorization`: `Bearer <your_token_here>`
- Query Parameters:
  - track_id `string` `required`: The ID of the track used as the basis for generating recommendations.
  - limit `integer` `optional`: The number of music recommendations to retrieve (default is 10).
- On Success Status Code: `200`
- Request Example:

  ```bash
  curl --request GET \
   --url 'http://localhost:8081/api/recommendations?track_id=2aDgJHhAbABvdW9NszrAPQ&limit=1' \
   --header 'Authorization: Bearer <your_token_here>'
  ```

- Response Example:

  ```json
  {
    "items": [
      {
        "album_type": "ALBUM",
        "album_total_tracks": 9,
        "album_images": [
          "https://i.scdn.co/image/ab67616d0000b27384851775712bec00a7e9ed6c",
          "https://i.scdn.co/image/ab67616d00001e0284851775712bec00a7e9ed6c",
          "https://i.scdn.co/image/ab67616d0000485184851775712bec00a7e9ed6c"
        ],
        "album_name": "HOP3",
        "album_release_date": "2011-01-01",
        "artists_name": ["RAN"],
        "explicit": false,
        "id": "7D4owue92VTPPMG0nNc4b5",
        "name": "Kulakukan Semua Untukmu",
        "is_liked": false
      }
    ]
  }
  ```

  </details>

### Like or Unlike Music

<details>
<summary>Like or unlike music from search results or recommendations</summary>
Request:

- Method: `POST`
- URL: `/api/track_activities`
- Headers:
  - `Content-Type`: `application/json`
  - `Authorization`: `Bearer <your_token_here>`
- Request Body:
  - track_id `string`: The ID of the track that the user wants to like or unlike.
  - is_liked `boolean`: A flag indicating the user's preference for the specified track. Set to true to like the track and false to unlike it.
- On Success Status Code: `200`
- Request Example:

  ```json
  {
    "track_id": "7D4owue92VTPPMG0nNc4b5",
    "is_liked": true
  }
  ```

- Response:

  - `200 OK`

  </details>
