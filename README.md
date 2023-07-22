# Prevue

Tool for visualization of training and testing machine learning models.

## Design

### Architecture

```mermaid
flowchart LR
    subgraph connectors[Connectors]
        keras[Keras]
        pytorch[Pytorch]
        sklearn[Scikit-Learn]
    end
    subgraph core_module[Core Module]
        db[Database]
        backend[Backend]
        frontend[Frontend]

    end
    subgraph debug_module[Debug Module]
        pg_admin[Postgres Admin]
        jupyter[Jupyter Notebook]
    end
    connectors --> backend
    backend --> db
    backend --> frontend

    pg_admin --> db
    jupyter --> backend

```

### Database design

```mermaid
erDiagram
    CITY {
        int city_id
        string name
        string state_abbreviation
    }
    STATE {
        string state_abbreviation
        string name
        int country_id
    }
    COUNTRY {
        int country_id
        string name
    }
    COUNTRY ||--|{ STATE : "Has"
    STATE ||--|{ CITY : "Has"
    CITY ||--o| STATE : "Is capital of"
    CITY ||--o| COUNTRY : "Is capital of"
```

### Deployment

```mermaid
flowchart TD
    subgraph db[Database]
        deployment_db[Deployment]
        service_db[Service]
        persistent_volume_db[Persistent Volume]
        deployment_db o--o persistent_volume_db
        deployment_db o--o service_db
    end
    subgraph backend[Backend]
        deployment_backend[Deployment]
        service_backend[Service]
        ingress_backend[Ingress]
        deployment_backend o--o service_db
        deployment_backend o--o service_backend
        service_backend o--o ingress_backend
    end
    subgraph frontend[Frontend]
        deployment_frontend[Deployment]
        service_frontend[Service]
        ingress_frontend[Ingress]

        deployment_frontend o--o service_frontend
        service_frontend o--o ingress_frontend

    end
    subgraph k8s[Kubernetes]
        db
        backend
        frontend
    end

    prevue_client[Prevue Client] --> ingress_backend
```

## Development

Prerequisites:

- tilt
- docker
- kind

### Start development environment

```bash
tilt up
```

### Stop development environment

```bash
tilt down
```

### Connect to backend using python client

```python
from prevue import PrevueKerasCallback

callback = PrevueKerasCallback(
    connector_name="keras",
    project_name="test",
    uid="test",
    url="http://localhost:8080"
)

model.fit(x_train, y_train, epochs=5, callbacks=[callback])
```

### Connect to backend using docker

Build backend:

```bash
 docker build . -t backend 
```

Run docker build of backend 

```bash 
docker run -p 8080:8080 backend:latest
```