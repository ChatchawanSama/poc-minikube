Intro
Hello everyone, it’s me, Chatchawan Sama. Today, I will demonstrate Minikube. But first, what is Kubernetes (Kube)?

Great! Let me tell you a bit about the story of Kube.

Before Kubernetes, we faced challenges with managing and maintaining containers. Imagine you have 10 containers—you could probably manage them by yourself. But what if you had 100 or more? That would be much harder, right?

This is where "Container Orchestration" comes in to help us manage and maintain our containers, and Kubernetes is one of the leading container orchestration platforms.


Flow
1. Call GET /config Service A
2. Service A will call Service B
3. Service B will read secret and config map value
4. Service B return values to Service A
5. Service A return secret and config map value 



Demo Time


Install Minikube 

brew cask install minikube

Install kubect

https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/


Start Minikube (Don’t forget to open docker desktop)

minikube start

After run this command, it will create docker container for minikube



Create API 

Do Go API for call service A, and service A will call API from service B

Do Go API for service B, and service will read secret and config map value and return them to service A

Create Dockerfile for service A and service B


Build Image in local
In service A path

docker build -t service-a:latest .

In service B path

docker build -t service-b:latest .



Build Image in Minikube Environment

eval $(minikube docker-env)
docker build -t service-a:latest ./service-a
docker build -t service-b:latest ./service-b



Deploy Pod 

kubectl apply -f ./deployments/service-a-deployment.yaml
kubectl apply -f ./deployments/service-a-service.yaml  
kubectl apply -f ./deployments/service-b-deployment.yaml
kubectl apply -f ./deployments/service-b-service.yaml
kubectl apply -f ./config/configmap.yaml  
kubectl apply -f ./config/secret.yaml   



Check service url in Minikube

minikube service service-a --url
minikube service service-b --url

when we run the commands, we will get result like this.

a677235@AR8040677235N1 poc-minikube % minikube service service-a --url
http://127.0.0.1:51015
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.

a677235@AR8040677235N1 poc-minikube % minikube service service-b --url
http://127.0.0.1:55730
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.


Test CURL

Curl Service A

curl http://127.0.0.1:55751/config

We will got response like this.

{"response":"{\"secretData\":\"c2VjcmV0X3ZhbHVl\",\"configMapVal\":\"config_value\"}\n"}

