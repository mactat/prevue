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

# Example
k8s_yaml('./examples/kubernetes.yaml')
docker_build('mactat/prevue-example', '.', dockerfile="./examples/dockerfile")
k8s_resource('prevue-example', labels=["debug-module"])

# button

pod_exec_script = '''
set -eu
# get k8s pod name from tilt resource name
POD_NAME="$(tilt get kubernetesdiscovery prevue-example -ojsonpath='{.status.pods[0].name}')"
kubectl exec "$POD_NAME" -- python3 /examples/Irys_dataset.py
'''

load('ext://uibutton', 'cmd_button')
cmd_button('podexec',
        argv=['sh', '-c', pod_exec_script],
        resource='prevue-example',
        icon_name='rocket_launch',
        text='start',
)

