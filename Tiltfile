# Database
k8s_yaml('./db/pv.yaml')
k8s_resource('database', labels=["core-module"])
k8s_yaml('./db/kubernetes.yaml')
k8s_resource('database', labels=["core-module"], port_forwards=5432)
k8s_yaml('./db/pg_admin.yaml')
k8s_resource('pgadmin', labels=["debug-module"], port_forwards='8889:80')

# Backend
backend = read_yaml_stream('./backend/kubernetes.yaml')
backend[0]['spec']['template']['spec']['containers'][0]['env'][5]['value'] = 'debug'
k8s_yaml(encode_yaml_stream(backend))
k8s_resource('prevue-backend', port_forwards=8080, labels=["core-module"], resource_deps=['database'])
docker_build('mactat/prevue-backend', './backend', dockerfile='./backend/dockerfile')