## minikube

### 安装

minikube 必须包含 docker，kubelet 才可安装

- kubectl 安装命令: [参考](https://kubernetes.io/docs/tasks/tools/)

``` sh
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```

- minikube 安装: [参考](https://minikube.sigs.k8s.io/docs/start/)

``` sh
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

初次启动使用

``` sh
minikube start
# ... 
# 初次启动程序需要拉取一些镜像，需要等待一段时间
```

### dashboard

- 访问后台面板

``` sh
# 本机可访问
minikube dashboard

# 若想外部访问该面板，则需要开启一个代理
kubectl proxy --port=8888 --address='0.0.0.0' --accept-hosts='^.*'
```
