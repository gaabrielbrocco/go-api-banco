# go-api-banco

## 🇧🇷 PT-BR

Este repositório contém uma API desenvolvida em Go (Golang) com foco na aplicação da **Clean Architecture** (Arquitetura Hexagonal).

---

### 📌 Objetivo

Este projeto serve como base de estudo para organização e desenvolvimento de APIs com Go, utilizando uma estrutura desacoplada que facilita a manutenção e evolução do código. A aplicação implementa um CRUD simples para usuários, bancos e contas bancárias.

### 📁 Estrutura do Projeto

- `cmd/` — Ponto de entrada da aplicação.
- `internal/` — Lógica de negócio e domínios internos.
- `pkg/di/` — Responsável pela injeção de dependências.
- `Dockerfile` e `docker-compose.yml` — Para containerizar e facilitar o deploy da aplicação.

### 🚀 Tecnologias Utilizadas

- Go 1.24
- Docker & Docker Compose
- PostgreSQL
- GitHub Actions (para deploy automático)

### ⚙️ Deploy Automatizado com GitHub Actions

Este projeto inclui um workflow de CI/CD configurado em `.github/workflows/deploy.yml` que:

- Constrói a imagem Docker da API.
- Salva a imagem como arquivo `.tar`.
- Envia a imagem para um servidor remoto via SCP.
- Realiza o deploy no servidor via SSH, atualizando o container Docker em execução.

#### Configurações necessárias para o deploy automático:

No repositório GitHub, configure os *Secrets*:

| Nome do Secret   | Descrição                              |
|------------------|--------------------------------------|
| `ENV`            | Conteúdo do arquivo `.env` com variáveis sensíveis |
| `SSH_HOST`       | IP ou domínio do servidor remoto     |
| `SSH_USER`       | Usuário SSH do servidor               |
| `SSH_PRIVATE_KEY`| Chave privada SSH para autenticação  |

#### Como funciona o workflow:

1. O GitHub Actions é acionado em pushes na branch `master`.
2. Ele cria o arquivo `.env` com variáveis do *secret*.
3. Constrói a imagem Docker com as variáveis de ambiente.
4. Salva e transfere a imagem para o servidor.
5. No servidor, para o container antigo, remove e sobe um novo com a imagem atualizada.

---

### ▶️ Como Executar Localmente

Clone o repositório e use Docker Compose para subir o ambiente:

```bash
git clone https://github.com/gaabrielbrocco/go-api-banco.git
cd go-api-banco
docker compose up --build -d
```



## 🇺🇸 EN-US

This repository contains an API developed in Go (Golang) focused on applying **Clean Architecture** (Hexagonal Architecture).

---

### 📌 Purpose

This project serves as a study base for organizing and developing APIs with Go, using a decoupled structure that facilitates maintenance and code evolution. The application implements a simple CRUD for users, banks, and bank accounts.

### 📁 Project Structure

- `cmd/` — Application entry point.
- `internal/` — Business logic and internal domains.
- `pkg/di/` — Responsible for dependency injection.
- `Dockerfile` and `docker-compose.yml` — To containerize and ease application deployment.

### 🚀 Technologies Used

- Go 1.24
- Docker & Docker Compose
- PostgreSQL
- GitHub Actions (for automated deployment)

### ⚙️ Automated Deployment with GitHub Actions

This project includes a CI/CD workflow configured in `.github/workflows/deploy.yml` that:

- Builds the Docker image of the API.
- Saves the image as a `.tar` file.
- Sends the image to a remote server via SCP.
- Deploys the server via SSH, updating the running Docker container.

#### Required settings for automatic deployment:

Configure the following *Secrets* in the GitHub repository:

| Secret Name      | Description                          |
|------------------|------------------------------------|
| `ENV`            | Content of the `.env` file with sensitive variables |
| `SSH_HOST`       | Remote server IP or domain          |
| `SSH_USER`       | SSH user of the server              |
| `SSH_PRIVATE_KEY`| SSH private key for authentication |

#### How the workflow works:

1. GitHub Actions is triggered on pushes to the `master` branch.
2. It creates the `.env` file with the secret variables.
3. Builds the Docker image with environment variables.
4. Saves and transfers the image to the server.
5. On the server, stops the old container, removes it, and starts a new one with the updated image.

---

### ▶️ How to Run Locally

Clone the repository and use Docker Compose to spin up the environment:

```bash
git clone https://github.com/gaabrielbrocco/go-api-banco.git
cd go-api-banco
docker compose up --build -d


