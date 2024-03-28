# Ingress

## Install

ingress [安装文档](https://kubernetes.github.io/ingress-nginx/deploy/)

```sh
# 
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

# 检查是否已安装
kubectl get pods --namespace=ingress-nginx

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=120s

# 设置 ingress rules ...
# ...

# 开启代理 监听本机 8080 端口
kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 8080:80
```
