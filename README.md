# hello-world

A basic hello world app that runs in kubernetes on port 8080, has a few extra endpoints for debugging or experimentation.

## Kubernetes manifests

Quick deploy using 

```bash
kubectl apply -f https://raw.githubusercontent.com/gabrie30/hello-world/main/kubernetes-manifests.yaml
```

## Dockerhub

https://hub.docker.com/r/gabrie30/hello-world

```
Supported Endpoints (see repo for complete list in main.go)

* `/ => "hello-world"`

* `/healthz => "healthz"`

* `/timeout?sleep=3 => will sleep for 3 seconds then respond with "Slept for 3 seconds"`
​```
