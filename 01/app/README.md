Сборка
```
docker build -t anfilat/otussa01:0.3 .
```
Публикация
```
docker push anfilat/otussa01:0.3
```
Minikube
```
minikube start
kubectl create namespace myapp
kubectl get namespaces
kubectl config set-context --current --namespace=myapp
```
Получить адрес
```
minikube ip
```
Добавить его в `/etc/hosts`
```
192.168.49.2 arch.homework
```
Применить манифесты
```
kubectl apply -f .
```
Посмотреть что запущено
```
kubectl get pods
```
Проверить работу
```
curl http://arch.homework/otusapp/anfilat/health
curl http://arch.homework/otusapp/anfilat/healthz
```
