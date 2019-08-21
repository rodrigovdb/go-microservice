This is a simple microservice written in Go that uses Docker and Docker Compose.

# Running with Docker Compose

First, copy `.env_sample` to `.env`. After, open new `.env` file and fill with propper credentials. So, run everything with

```
$ docker-compose up
```

Access
`http://localhost:8080/?to=rodrigovdb@gmail.com&from=rodrigo@devlandia.net&subject=Test%20subject&body=My%20email%20body%20message`

And that's it.

# Running with Docker

Build
```
$ docker build --tag=golang-microservice .
```

Create a container
```
$ docker container create --name=golang-microservice -p 8080:8080 --env-file=.env golang-microservice
```

Start it
```
$ docker start golang-microservice
```

Watch logs
```
$ docker logs golang-microservice -f
```

Access
`http://localhost:8080/?to=rodrigovdb@gmail.com&from=rodrigo@devlandia.net&subject=Test%20subject&body=My%20email%20body%20message`

And that's it.

# References

* https://medium.com/trainingcenter/golang-e-docker-d2d9dedd82c0
* https://github.com/evandroferreiras/golang-webscrapper-docker-tutorial
* https://github.com/joho/godotenv
