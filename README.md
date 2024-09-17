minikube start
docker build -t service-a:latest .
docker build -t service-b:latest .

eval $(minikube docker-env)
docker build -t service-a:latest ./service-a
docker build -t service-b:latest ./service-b

# Deploy
kubectl apply -f ./deployments/service-a-deployment.yaml 
kubectl apply -f ./deployments/service-a-service.yaml  
kubectl apply -f ./deployments/service-b-deployment.yaml
kubectl apply -f ./deployments/service-b-service.yaml 
kubectl apply -f ./config/configmap.yaml  
kubectl apply -f ./config/secret.yaml   

minikube service service-a --url
minikube service service-b --url
