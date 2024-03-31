
# install ingress-nginx
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx \
    -n ingress-nginx --create-namespace 

# add prometheus-stack
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus-stack prometheus-community/kube-prometheus-stack \
    -n ops --create-namespace \
    --set alertmanager.enabled=true \
    --set alertmanager.ingress.enabled=true \
    --set alertmanager.ingress.ingressClassName=nginx \
    --set alertmanager.ingress.hosts={alertmanager.ops.kilii.cc} \
    --set grafana.enabled=true \
    --set grafana.ingress.enabled=true \
    --set grafana.ingress.ingressClassName=nginx \
    --set grafana.ingress.hosts={grafana.ops.kilii.cc} \
    --set grafana.adminPassword=helloworld \
    --set prometheus.enabled=true \
    --set prometheus.ingress.enabled=true \
    --set prometheus.ingress.ingressClassName=nginx \
    --set prometheus.ingress.hosts={prome.ops.kilii.cc} \
    --wait

# add helm repo
helm repo add bitnami https://charts.bitnami.com/bitnami

# # add elastic
helm install -f ./elasticsearch/values.yaml \
    --name-template elasticsearch ./elasticsearch -n ops --create-namespace \
    --wait
# helm install -f es  elastic/elasticsearch \
#     -n ops --create-namespace \
#     --set replicas=1 \
#     --set minimumMasterNodes=1 \
#     --set persistence.enabled=false \
#     --set ingress.enabled=true \
#     --set ingress.className=nginx \
#     --set "ingress.hosts[0].host=es.ops.kilii.cc" \
#     --set resources.requests.cpu="200m" \
#     --set resources.requests.memory="512Mi" \
#     --set resources.limits.cpu="200m" \
#     --set resources.limits.memory="512Mi" \
#     --wait


# add filebeat
helm install filebeat elastic/filebeat \
    -n ops --create-namespace \
    --wait
helm install -f ./filebeat/values.yaml \
    --name-template filebeat ./filebeat -n ops --create-namespace \
    --wait


# add kibana
helm install kibana elastic/kibana \
    -n ops --create-namespace \
    --set ingress.enabled=true \
    --set "ingress.hosts[0].host=kibana.ops.kilii.cc" \
    --set "ingress.hosts[0].paths[0].path=/" \
    --set resources.requests.cpu="200m" \
    --set resources.requests.memory="512Mi" \
    --set resources.limits.cpu="200m" \
    --set resources.limits.memory="512Mi" \
    --wait
    

# add kafka with none auth
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install kafka bitnami/kafka \
  --namespace ops --create-namespace \
  --set listeners.client.protocol=PLAINTEXT \
  --set resources.requests.cpu="200m" \
  --set resources.requests.memory="256Mi" \
  --set resources.limits.cpu="200m" \
  --set resources.limits.memory="256Mi" \
  --wait
