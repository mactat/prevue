# Database
k8s_yaml('./db/pv.yaml')
k8s_resource('database', labels=["core-module"])
k8s_yaml('./db/kubernetes.yaml')
k8s_resource('database', labels=["core-module"])

# Backend
k8s_yaml('./backend/kubernetes.yaml')
k8s_resource('prevue-backend', port_forwards=8080, labels=["core-module"])
docker_build('mactat/prevue-backend', './backend', dockerfile='./backend/dockerfile')