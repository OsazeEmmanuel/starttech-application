# StartTech Enterprise CI/CD on Amazon EKS

## Project Overview

This project demonstrates the deployment of a modern full-stack application on **Amazon Elastic Kubernetes Service (EKS)** using **Terraform**, **Docker**, **GitHub Actions**, and **AWS cloud services**.

The application consists of:

- **Frontend:** React + Vite
- **Backend:** Golang REST API
- **Database:** MongoDB Atlas
- **Caching:** Redis
- **Container Orchestration:** Amazon EKS
- **Ingress:** AWS Load Balancer Controller (Application Load Balancer)
- **Container Registry:** Amazon ECR
- **Infrastructure as Code:** Terraform
- **CI/CD:** GitHub Actions

---

# Architecture

```
                    GitHub Repository
                           │
                           │
                    GitHub Actions
                           │
        ┌──────────────────┴──────────────────┐
        │                                     │
        ▼                                     ▼
 Build Backend Image                  Build Frontend Image
        │                                     │
        └──────────────┬──────────────────────┘
                       │
                       ▼
                Amazon Elastic Container Registry (ECR)
                       │
                       ▼
                Amazon Elastic Kubernetes Service (EKS)
                       │
      ┌────────────────┴─────────────────┐
      │                                  │
      ▼                                  ▼
Frontend Deployment              Backend Deployment
      │                                  │
      └──────────────┬───────────────────┘
                     ▼
          AWS Application Load Balancer
                     │
                     ▼
                  End Users

Backend communicates with:

• MongoDB Atlas
• Redis
```

---

# Technologies Used

| Technology | Purpose |
|------------|----------|
| Go | Backend API |
| React | Frontend |
| Vite | Frontend Build Tool |
| Docker | Containerization |
| Kubernetes | Container Orchestration |
| Amazon EKS | Kubernetes Cluster |
| Amazon ECR | Container Registry |
| Terraform | Infrastructure Provisioning |
| AWS ALB Controller | Ingress |
| GitHub Actions | CI/CD |
| MongoDB Atlas | Database |
| Redis | Caching |

---

# Repository Structure

```
.
├── Client/
│   ├── Dockerfile
│   └── ...
│
├── Server/
│   └── MuchToDo/
│       ├── Dockerfile
│       └── ...
│
├── k8s/
│   ├── backend-deployment.yaml
│   ├── backend-service.yaml
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   ├── ingress.yaml
│   ├── namespace.yaml
│   └── secret.yaml
│
└── .github/
    └── workflows/
        └── application-deploy.yml
```

---

# Kubernetes Resources

The application is deployed using:

- Namespace
- Secrets
- Deployments
- Services
- Ingress

Deployments include:

- Backend Deployment
- Frontend Deployment

Services include:

- Backend ClusterIP
- Frontend ClusterIP

Ingress:

- AWS Application Load Balancer

---

# CI/CD Pipeline

GitHub Actions performs the following steps automatically:

1. Checkout source code
2. Configure AWS credentials
3. Login to Amazon ECR
4. Build Backend Docker Image
5. Push Backend Image
6. Build Frontend Docker Image
7. Push Frontend Image
8. Update kubeconfig
9. Deploy to Amazon EKS
10. Verify rollout

---

# Infrastructure

Infrastructure was provisioned using Terraform.

Resources include:

- VPC
- Public Subnets
- Private Subnets
- NAT Gateway
- Internet Gateway
- Route Tables
- IAM Roles
- Amazon EKS Cluster
- Managed Node Group
- Security Groups

---

# Docker Images

Backend

```
starttech-backend-api
```

Frontend

```
starttech-frontend
```

---

# Kubernetes Deployment

Deploy resources:

```bash
kubectl apply -f k8s/
```

Check deployments:

```bash
kubectl get deployments -n starttech
```

Check pods:

```bash
kubectl get pods -n starttech
```

Check services:

```bash
kubectl get svc -n starttech
```

Check ingress:

```bash
kubectl get ingress -n starttech
```

---

# Application Endpoints

Frontend

```
http://<ALB-DNS>/
```

Backend Health

```
http://<ALB-DNS>/api/health
```

Ping

```
http://<ALB-DNS>/api/ping
```

Swagger

```
http://<ALB-DNS>/api/swagger/index.html
```

---

# Environment Variables

Backend requires:

```
PORT
MONGO_URI
DB_NAME
JWT_SECRET_KEY
JWT_EXPIRATION_HOURS
REDIS_ADDR
REDIS_PASSWORD
ENABLE_CACHE
ALLOWED_ORIGINS
```

---

# Security

Sensitive configuration is stored using:

- Kubernetes Secrets
- GitHub Actions Secrets

No secrets are committed into source control.

---

# Monitoring

Useful commands

Pods

```bash
kubectl get pods -n starttech
```

Logs

```bash
kubectl logs -f deployment/starttech-backend -n starttech
```

Rollout

```bash
kubectl rollout status deployment/starttech-backend -n starttech
```

---

# Future Improvements

- HTTPS using AWS Certificate Manager
- External DNS integration
- Horizontal Pod Autoscaler
- Prometheus Monitoring
- Grafana Dashboards
- ArgoCD GitOps
- Blue/Green Deployments

---

# Author

**Osaze Emmanuel**

DevOps Engineer

Built as part of the **StartTech Enterprise CI/CD on Amazon EKS Assessment**.
