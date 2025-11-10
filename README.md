# 📝 Taskify

![Go Version](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Build](https://github.com/shakarpg/taskify/actions/workflows/go.yml/badge.svg)

API REST escrita em **Golang** para gerenciamento de tarefas (To-Do List), com **autenticação JWT**, **banco PostgreSQL**, e **testes automatizados**.

---

## 🚀 Tecnologias

- **Go 1.22**
- **Chi Router** (rotas HTTP)
- **GORM** (ORM para PostgreSQL)
- **JWT** (autenticação)
- **Testify** (testes unitários)
- **Docker + Docker Compose**
- **GitHub Actions** (CI/CD)

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
