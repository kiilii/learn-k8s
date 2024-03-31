INFRA_DIR=$(shell a=`basename $$PWD` && echo $$a/infra)

# add helm repo
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add elastic https://helm.elastic.co

# install ingress-nginx
cd $INFRA_DIR/ingress-nginx && \
helm install ingress-nginx ingress-nginx -n ingress-nginx --create-namespace 

# add prometheus-stack
cd $INFRA_DIR/kube-prometheus-stack && \
helm install -f values.yaml prometheus-stack prometheus-community -n ops --create-namespace 

# add elastic
cd $INFRA_DIR/elasticsearch && \
helm install --name-template elasticsearch -f values.yaml . -n ops --create-namespace

# add kafka 
helm install kafka oci://registry-1.docker.io/bitnamicharts/kafka -n ops --create-namespace