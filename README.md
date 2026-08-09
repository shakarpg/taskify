# 🚀 Taskify — Secure REST API em Go

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql&logoColor=white)
![JWT Auth](https://img.shields.io/badge/Auth-JWT-black?style=flat&logo=jsonwebtokens)
![CI Build](https://github.com/shakarpg/taskify-projectgo/actions/workflows/ci.yml/badge.svg)

API RESTful segura para gerenciamento de tarefas e projetos desenvolvida em Golang, aplicando boas práticas de arquitetura limpa, autenticação via JWT, persistência com PostgreSQL e testes automatizados.

---

## 📌 Funcionalidades

- 🔑 **Autenticação & Autorização**: Registro e login de usuários com hash seguro de senha (Bcrypt) e tokens JWT.
- 📝 **Gestão de Tarefas (CRUD)**: Criação, listagem, atualização e exclusão de tarefas.
- 🛡️ **Segurança**: Middlewares de autenticação e validação de requisições.
- 🗄️ **Persistência Relacional**: Integração com PostgreSQL e migrações de banco de dados.
- 🧪 **Testes Automatizados**: Suporte a testes unitários e de integração com cobertura de código.

---

## 🏗️ Arquitetura do Sistema

graph TD
    Client[Cliente / Frontend / Postman] -->|HTTP Request + JWT| Middleware[Auth Middleware]
    Middleware -->|Requisição Autenticada| Handler[HTTP Handlers / Controllers]
    Handler -->|Regras de Negócio| Service[Service / Business Layer]
    Service -->|Acesso a Dados| Repository[Repository Layer]
    Repository -->|SQL Queries| DB[(PostgreSQL Database)]

---
## 🧰 Como rodar o projeto

### 1️⃣ Clone o repositório
```bash
git clone https://github.com/shakarpg/taskify.git
cd taskify
```

### 2️⃣ Instale as dependências
```bash
go mod tidy
```

### 3️⃣ Suba o banco de dados com Docker
```bash
make docker-up
```

### 4️⃣ Rode a aplicação
```bash
make run
```

Acesse: [http://localhost:8080/health](http://localhost:8080/health)

---

## 🧪 Rodar os testes

```bash
make test
```

---

## 📡 Endpoints da API

### 🔓 Públicos
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/api/login` | Fazer login e obter token JWT |
| GET | `/health` | Health check |

### 🔒 Protegidos (requer JWT no header `Authorization: Bearer <token>`)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/api/tasks` | Listar todas as tarefas |
| POST | `/api/tasks` | Criar nova tarefa |
| GET | `/api/tasks/{id}` | Obter tarefa específica |
| PUT | `/api/tasks/{id}` | Atualizar tarefa |
| DELETE | `/api/tasks/{id}` | Deletar tarefa |

---

## 🧾 Exemplo de uso

### 1. Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"123456"}'
```

**Resposta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 2. Criar tarefa (com token)
```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -d '{"title":"Estudar Go","completed":false}'
```

### 3. Listar tarefas
```bash
curl -X GET http://localhost:8080/api/tasks \
  -H "Authorization: Bearer SEU_TOKEN_AQUI"
```

---

## 🐳 Docker

### Subir tudo (app + banco)
```bash
docker-compose up --build
```

### Parar containers
```bash
make docker-down
```

---

## 📂 Estrutura do Projeto

```
taskify/
├── .github/
│   └── workflows/
│       └── go.yml           # GitHub Actions CI/CD
├── cmd/
│   └── main.go              # Entry point da aplicação
├── internal/
│   ├── database/
│   │   └── db.go            # Conexão com PostgreSQL
│   ├── handlers/
│   │   ├── auth.go          # Handler de autenticação
│   │   └── tasks.go         # Handlers de tarefas
│   ├── models/
│   │   ├── task.go          # Modelo de Task
│   │   └── user.go          # Modelo de User
│   ├── middleware/
│   │   └── auth.go          # Middleware JWT
│   └── router/
│       └── router.go        # Configuração de rotas
├── tests/
│   └── tasks_test.go        # Testes automatizados
├── .env                     # Variáveis de ambiente
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
└── README.md
```

---

## 🧠 Próximos passos (melhorias)

- [ ] Adicionar hash de senha (bcrypt)
- [ ] Implementar refresh token
- [ ] Adicionar paginação nas listagens
- [ ] Documentação Swagger
- [ ] Deploy em produção (Railway, Render, Fly.io)
- [ ] Adicionar rate limiting
- [ ] Implementar logs estruturados (zerolog/zap)

---

## 📄 Licença

MIT License - sinta-se livre para usar e modificar!

---

## 👤 Autor

**Shakarpg**  
GitHub: [@shakarpg](https://github.com/shakarpg)
