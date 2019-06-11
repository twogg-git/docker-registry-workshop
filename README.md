# docker-registry-workshop
A quick workshop for Docker Registry

## Docker Example 
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

## DockerHub Registry

We are going to use DockerHUb free registry, so you need to create an account here https://hub.docker.com/   
<img height="450" width="400" src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-sign_in.png">

<img height="5" width="5" src="IMG">

## Pushing Images with CLI

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

```sh
docker build -t twoggtest/golang:1.0 .
```

```sh
docker run --name twoggtest-golang --rm -p 8081:8080 twoggtest/golang:1.0
```





