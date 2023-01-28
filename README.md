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
 - Go scripting
 - Go Coroutines/async programming
 - HTTP streaming (Huge CSV Uploads)
 - GIN(General Inverted Indices) for text searching
 - BRIN (Block Range Indices) for date range quering
 - Hash indices for exact matching
 - SQL batch inserts/updates for better performance
 - Vue composition API
 -
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
git clone https://github.com/mzaidannas/packform-test.git
```
Move to project directory
```sh
cd packform-test
```
Create env file with required environment variables
```sh
tee .env << ENV
DB_HOST='postgres'
SECRET='7B4XAV8XOZ9BT52O'
USERNAME="username"
PASSWORD="password"
ENV
```
Run project
```sh
docker-compose up
```

Create new user through browser on url with the username/password in the .env file
`http://localhost:3000`

Import data using script
```sh
docker-compose exec -it packform-test ./import-data
```
