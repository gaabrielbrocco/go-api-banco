# go-api-banco

Este repositório contém uma API desenvolvida em Go (Golang) com foco na aplicação da **Clean Architecture**

## 📌 Objetivo

O objetivo deste projeto é servir como base de estudo e prática de arquitetura limpa no desenvolvimento de APIs com Go. A aplicação implementa um CRUD simples e demonstra como estruturar uma aplicação de forma organizada e desacoplada.

## 📁 Estrutura do Projeto

- `cmd/` — Ponto de entrada da aplicação.
- `internal/` — Contém a lógica de negócio e domínios internos.
- `pkg/di/` — Responsável pela injeção de dependências.
- `Dockerfile` e `docker-compose.yaml` — Para facilitar o deploy da aplicação em containers.

## 🚀 Tecnologias Utilizadas

- Go
- Docker Compose

## ▶️ Como Executar

Clone o repositório e utilize o Docker para subir a aplicação:

```bash
git clone https://github.com/gaabrielbrocco/go-api-banco.git
cd go-api-banco
docker compose up --build -d
