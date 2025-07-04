# Todo App Deployment

A full-stack Todo application with a modern Svelte frontend, Go backend, containerization, and Kubernetes-native deployment. This repo demonstrates real-world DevOps with GitHub Actions CI and ArgoCD GitOps CD—all ready for local and cloud-native workflows.

---

## 🚀 Features

- **Frontend:** Svelte for a responsive UI ([go-todo-frontend/](https://hub.docker.com/repository/docker/docker0y4sh/todo-frontend/general))
- **Backend:** Go REST API ([backend/](https://hub.docker.com/repository/docker/docker0y4sh/todo-backend/general)), containerized for K8s
- **Manifests:** YAML & Jinja templates for flexible Kubernetes deployment
- **CI/CD:** Automated with GitHub Actions (CI) and ArgoCD (CD)
- **No Vendor Lock-In:** Deploy on any Kubernetes cluster (local/cloud)
- **DevOps Best Practices:** Full pipeline and manifest templating

---

## Repository Structure

```
go-todo-frontend/     # Svelte frontend code
backend/              # Go backend API
manifests/            # K8s manifests & Jinja templates
.github/workflows/    # GitHub Actions CI/CD pipelines
Dockerfile            # Container build files
README.md             # Project overview 
```

---

## 🏁 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yashpal2104/todo-app-deployment.git
cd todo-app-deployment
```

---

### 2. Frontend: Svelte

```bash
cd go-todo-frontend
npm install
npm run dev
```
Visit `http://localhost:5173` (or as printed in terminal) to view the app.

---

### 3. Backend: Go API

```bash
cd backend
go build -o app
./app
```

---

### 4. Kubernetes Deployment

1. Ensure you have a Kubernetes cluster running (local: Minikube, Kind, K3d, etc.)
2. Render any Jinja templates if applicable:
    ```bash
    jinja2 manifests/todo-deployment.yaml.j2 > manifests/todo-deployment.yaml
    ```
3. Apply manifests:
    ```bash
    kubectl apply -f manifests/todo-deployment.yaml
    ```

---

### 5. CI/CD

- **CI:** All pushes trigger GitHub Actions for lint, build, and test
- **CD:** ArgoCD (if set up) automatically syncs manifests from this repo to your K8s cluster

---

## 💡 Notes

- No cloud service is required; works on local Kubernetes.
- Manifests are ready for cloud if you want to deploy later.
- You can tweak the manifests and templates for your needs.

---

## 📸 Screenshots
# App
![todo-app](https://github.com/user-attachments/assets/bc8d3d32-6bea-47fc-8d15-3d28f29c6910)

# Github Actions CI
![github com_yashpal2104_todo-app-deployment_actions](https://github.com/user-attachments/assets/01591c33-e9c9-47de-b56e-b1a87c180554)

# ArgoCD Deployment
![argocd](https://github.com/user-attachments/assets/465c29b9-57c5-41c5-b8c7-c11347e6143b)

---

## 📝 License

This project is licensed under the [MIT License](LICENSE).

---

## 🤝 Connect

If you like this project, connect with me on [LinkedIn](https://www.linkedin.com/in/yash-pal-88621224b/).
---

![Made with Svelte](https://img.shields.io/badge/Svelte-frontend-orange)
![Go Backend](https://img.shields.io/badge/Go-backend-blue)
