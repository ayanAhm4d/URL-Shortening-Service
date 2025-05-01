
## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)](https://linkedin.com/in/www.linkedin.com/in/ayanahmad15) [![X](https://img.shields.io/badge/X-black.svg?logo=X&logoColor=white)](https://x.com/ayanAhm4d) 

# 💻 Tech Stack:

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- [UUID](https://github.com/google/uuid)
- [Govalidator](https://github.com/asaskevich/govalidator)
- [Godotenv](https://github.com/joho/godotenv)

# 📊 GitHub Stats:
![](https://github-readme-stats.vercel.app/api?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false)<br/>
![](https://github-readme-streak-stats.herokuapp.com/?user=ayanAhm4d&theme=dark&hide_border=false)<br/>
![](https://github-readme-stats.vercel.app/api/top-langs/?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false&layout=compact)

---
[![](https://visitcount.itsvg.in/api?id=ayanAhm4d&icon=0&color=0)](https://visitcount.itsvg.in)


# URL Shortener

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)


## Features

- Shorten URLs: Generate shortened links for long URLs.

- Custom Short URLs: Users can create their custom short links.

- Automatic Expiry: URLs expire after a specified duration (default: 24 hours).

- Rate Limiting: Prevent abuse by limiting requests to 10 per user every 30 minutes.

- HTTPS Enforcement: Ensures all URLs are served with HTTP or HTTPS.



## Project Structure


    ```
    url-shortener/
    ├── main.go
    ├── .env
    ├── go.mod
    ├── go.sum
    ├── config/
    │   └── config.go
    ├── handlers/
    │   └── url.go
    ├── middleware/
    │   └── rate_limiter.go
    ├── models/
    │   └── request.go
    ├── utils/
    │   └── helpers.go
    └── redis/
        └── client.go
    
    ```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ayanAhm4d/URL-shortener.git
   ```
2. Set up .env file:
    ```
   
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    API_RATE_LIMIT=10
    API_RATE_DURATION=30m
    DEFAULT_EXPIRY=24h
    BASE_URL=http://localhost:8080


    ```
   
3. Install dependencies:
   ```
   go mod tidy
   ```


1. Start the server:
   ```
   go run main.go
   ```



2. Access the application at http://localhost:8080.

## Usage

API Endpoints

POST /shorten: Shorten a URL.

Request Body:
    ```
    {
      "url": "https://example.com/very-long-url",
      "custom_short": "myshort",
      "expiry": 12
    }
    
    ```
Response:
    ```
    {
      "short_url": "http://localhost:8080/myshort",
      "expires_in": "12h0m0s"
    }
    
    ```
GET /:short: Redirect to the original URL.

curl http://localhost:8080/myshort.


You can use tools like curl, Postman, or your browser to interact with the API.

Contributing
If you'd like to contribute, feel free to open an issue or submit a pull request.

Contact

For any queries, reach out at www.ayan007ahmad@gmail.com.
