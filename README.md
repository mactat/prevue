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

This Mermaid diagram represents the database schema with the tables `users`, `projects`, `models`, and `metrics`. The primary keys are denoted with (PK) and foreign keys with (FK). The relationships between the tables are represented by lines connecting them. For example, users can own multiple projects, and each project is associated with a single user through the `user_id` foreign key. Similarly, models are used in multiple projects, and each project can use a single model, connected through the `model_id` foreign key. Users and models can also have interactions with metrics, which are represented by lines connecting them.

Note that the timestamp is defined in UTC. 

```mermaid
erDiagram
    users {
        user_id(PK) int
        username  varchar
        email  varchar
        password  varchar
    }
  
    projects {
        project_id(PK)  int
        project_name  varchar
        user_id(FK)  int
        model_id(FK)  int
    }

    models {
        model_id(PK)  int
        model_name varchar
        connector_name varchar
        architecture text
    }
  
    metrics {
        metric_id(PK) int
        model_id(FK) int
        epoch int
        batch int
        loss_name  varchar
        loss_value  float
        metric_name  varchar
        metric_value  float
        timestamp  date
    }
    
    users ||--o{ projects : "own"
    users ||--|{ metrics : "interact with"
    models ||--o{ projects : "used in"
    models ||--|{ metrics : "used in"
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