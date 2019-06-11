<img height="5" width="5" src="IMG">

<img src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/">

<img src="https://raw.githubusercontent.com/twogg-git/go-docker-hub/master/docker-hub.png">

# Docker Registry Workshop
A quick workshop for Docker Registry.

Tools to use:
- https://docs.docker.com/
- https://hub.docker.com/

## Working locally

We are going to start with an Apache httpd deploy. Here is our Dockerfile and a simple index.html file.

```sh
FROM httpd:2.4.39-alpine 
ADD index.html /usr/local/apache2/htdocs/
EXPOSE 80
```

```sh
<!DOCTYPE html>
<html>
<head>
  <title>Docker</title>
</head>
<body><center>
  <img src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-index-httpd.png">
  <h1 style="color:purple">Baby steps with Docker Registry + HTTPD!</h1>   
</center></body>
</html>
```

### Then build, test and commit locally
```sh
docker build -t httpd-test .
```

```sh
docker images
```

```sh
docker run --name httpd-test --rm -p 81:80 httpd-test
```

```sh
http://localhost:81/
```

```sh
docker container ls
```

```sh
docker commit httpd-test twogghub/httpd-test:version1
```

## DockerHub Registry

We are going to use DockerHUb free registry, so you need to create an account here https://hub.docker.com/   

<img height="450" width="400" src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-sign_in.png">


## Push a image into DockerHub 

Now, push the images from our terminal following the dockerhub instructions.
<img height="170" width="420" src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-cli_push.png">

```sh
docker login --username=twoggtest
```

```sh
docker tag httpd-test twoggtest/httpd:1.0
```

```sh
docker push twoggtest/httpd:1.0
```



## DockerHub Automated Builds

<img src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-create_repository.png">

<img src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-build_settings.png">


```sh
docker build -t twoggtest/golang:1.0 .
```

```sh
docker run --name twoggtest-golang --rm -p 8081:8080 twoggtest/golang:1.0
```





