# Go Date-Time Web App

A minimal Go web service that returns the current date and time, containerized with Docker, and deployed to Kubernetes with two replicas and a public Service.

## Overview
- HTTP server listens on port 8080 and serves the current timestamp at the root path.
- Container image is published to Docker Hub.
- Kubernetes Deployment runs 2 replicas.
- Service type LoadBalancer exposes the app publicly on a cloud or lab cluster.

## Repository structure
- go-datetime-app/main.go
- go-datetime-app/go.mod
- go-datetime-app/Dockerfile
- go-datetime-app/deployment.yaml
- go-datetime-app/service.yaml

## Prerequisites
- Docker Desktop running and logged in to Docker Hub
- kubectl configured to a Kubernetes cluster context
- For public access, use a lab or cloud cluster that supports LoadBalancer Services, or install MetalLB on local clusters

## Run locally
From the repository root:
cd go-datetime-app
go run main.go

In another terminal
curl http://localhost:8080/

text

## Build and push Docker image
Replace DOCKERUSER if different; the image is already pushed as endeavor6969/go-datetime-app:latest.
cd go-datetime-app
docker login
docker build -t endeavor6969/go-datetime-app:latest .
docker push endeavor6969/go-datetime-app:latest

text

## Kubernetes manifests
Deployment with 2 replicas and a simple readiness probe on path “/”. Service is type LoadBalancer and maps port 80 to targetPort 8080.

Files:
- go-datetime-app/deployment.yaml
- go-datetime-app/service.yaml

If using a different image name, edit deployment.yaml:
- spec.template.spec.containers[0].image: endeavor6969/go-datetime-app:latest

## Deploy to Kubernetes
Apply manifests and verify rollout:
kubectl apply -f go-datetime-app/deployment.yaml
kubectl apply -f go-datetime-app/service.yaml
kubectl get deploy,rs,pods -l app=datetime
kubectl rollout status deployment/datetime-deploy

text

## Expose and test
- Cloud or lab cluster:
kubectl get svc datetime-service -o wide

Open http://EXTERNAL-IP/
text

- Local kind or minikube quick test:
kubectl port-forward deployment/datetime-deploy 8080:8080

Open http://localhost:8080/
text

## Troubleshooting
- ImagePullBackOff or ErrImagePull
  - Ensure the image exists and is public: docker pull endeavor6969/go-datetime-app:latest
  - Confirm the image name in deployment.yaml

- Readiness probe failing
  - Ensure the app listens on port 8080 and responds at path “/”
  - Increase probe delays if needed in deployment.yaml

- No EXTERNAL-IP on local clusters
  - Use port-forward for local testing
  - Install a local load balancer like MetalLB to assign external IPs, or deploy to a lab or cloud cluster

## Clean up
kubectl delete -f go-datetime-app/service.yaml
kubectl delete -f go-datetime-app/deployment.yaml

text

## Submission checklist
- Code committed under go-datetime-app/
- Docker image pushed and public at Docker Hub
- Kubernetes Deployment running 2 replicas
- Service reachable via EXTERNAL-IP on lab or cloud cluster, or verified loca