# Packform Interview Test in Golang/VueJS

Basically we have 3 services

 1. **Backend**
 2. **Frontend**
 3. **Importer**

 **Backend** is a go web service that will provide customer data
 **Frontend** is a VueJS application that will pull data from backend
 **Importer** is a simple script that will export data to postgresql from CSVs

# Technologies used

 - [**Fiber**](https://gofiber.io/)
> Express like Web Micro framework for Go
 - [**postgresql**](https://www.postgresql.org/)
> Data storage used to store and index data
 - [**Docker**](https://www.docker.com/)
> Used for local development setup and well as running all the services combined. Can also help in production deployment and autoscaling
# Concepts Implemented
 - Micro Service Architecture
 - Go scripting
 - Signal Trapping
 - Scheduled workers
 - Go Coroutines/async programming
 - HTTP streaming (Huge json responses)
 - GIN(General Inverted Indices) for text searching
 - BRIN (Block Range Indices) for date range quering
 - Hash indices for exact matching
 - Background workers(Sneaker rabbitmq client)
 - Batch async http calls (I/O wait minimized by using multiple http calls in parallel)
 - SQL batch inserts/updates for better performance

# Local setup and running
Install docker on your system
 - **MacOS**
> Install homebrew package manager
> ```sh
> /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
> ```
Install docker desktop and git with brew cask
> ```sh
> brew install git
> brew cask install docker
> ```

- **Windows**
> Install/Update windows package manager winget (Update app installer package from windows store)

Install docker desktop and gitusing winget
> ```sh
> winget install -e --id Git.Git
> winget install -e --id Docker.DockerDesktop
> ```

- **Ubuntu**

Install docker and git using apt package manager
> ```sh
> curl -fsSL https://get.docker.com -o get-docker.sh
> sudo sh get-docker.sh
> sudo groupadd docker
> sudo usermod -aG docker $USER
> newgrp docker
> rm get-docker.sh
> ```

Fetch code using git
```sh
git clone --recurse-submodules -j8 https://github.com/mzaidannas/packform-test.git
```
Move to project directory
```sh
cd packform-test
```
Create env file with required environment variables
```sh
tee .env << ENV
GO_ENV='development'
LOG_LEVEL='debug'
DB_HOST='postgres'
REDIS_URL='redis://redis:6379/0'
SECRET='EfuuQyxFeOlppV4t5Z5gRQ'
ENV
```
Run project
```sh
docker-compose up scheduler crawler api
```

Check products through browser on url
`http://localhost:3000/products`

# Fiber with Auth

[Postman collection](https://www.getpostman.com/collections/c862d012d5dcf50326f7)

## Endpoints

- GET /api - _Say hello_
    - POST /auth/login - _Login user_
    - GET /user/:id - _Get user_
    - POST /user - _Create user_
    - PATCH /user/:id - _Update user_
    - DELETE /user/:id - _Delete user_
    - GET /product - _Get all products_
    - GET /product/:id - _Get product_
    - POST /product - _Create product_
    - DELETE /product/:id - _Delete product_
