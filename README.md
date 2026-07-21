#  KoreAI
### Plataforma Inteligente Multi-Agente com Orquestração de LLMs

<div align="center">

![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Node.js](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)
![LangChain](https://img.shields.io/badge/LangChain-121212?style=for-the-badge&logo=chainlink&logoColor=white)
![OpenAI](https://img.shields.io/badge/OpenAI-412991?style=for-the-badge&logo=openai&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Qdrant](https://img.shields.io/badge/Qdrant-00BFFF?style=for-the-badge&logo=vector&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)

![Architecture](https://img.shields.io/badge/Architecture-Event--Driven-8A2BE2?style=for-the-badge)
![AI](https://img.shields.io/badge/AI-Multi--Agent-FF6B6B?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)

</div>

---

## 📑 Índice / Table of Contents

1. [🇬🇧 English Version](#-english-version)
   - [Overview](#overview)
   - [The Name "KoreAI"](#the-name-koreai)
   - [Core Concepts & Philosophy](#core-concepts--philosophy)
   - [Architecture Deep Dive](#architecture-deep-dive)
   - [How It Works: A Complete Flow](#how-it-works-a-complete-flow)
   - [Tech Stack & Polyglot Strategy](#tech-stack--polyglot-strategy)
   - [Features](#features)
   - [Project Structure](#project-structure)
   - [Installation & Quick Start](#installation--quick-start)
   - [Configuration Guide](#configuration-guide)
   - [Usage Examples](#usage-examples)
   - [API Reference](#api-reference)
   - [Agent Design System](#agent-design-system)
   - [Memory & RAG Pipeline](#memory--rag-pipeline)
   - [Tool Integration Framework](#tool-integration-framework)
   - [Security Architecture](#security-architecture)
   - [Performance & Scalability](#performance--scalability)
   - [Testing & Quality Assurance](#testing--quality-assurance)
   - [Deploy & DevOps](#deploy--devops)
   - [Contributing](#contributing)
   - [Roadmap](#roadmap)
   - [FAQ](#faq)
   - [License](#license)
2. [🇧🇷 Versão em Português](#-versão-em-português)
   - (seções espelhadas em português)

---

# 🇬🇧 English Version

## Overview

**KoreAI** is a sophisticated multi-agent AI platform that orchestrates Large Language Models (LLMs), tools, and memory to solve complex, multi-step problems. It is not a simple chatbot – it is a **cognitive operating system** that assigns tasks to specialized AI agents, each with its own personality, capabilities, and access to external services.

The platform is designed for **enterprise automation**, **research acceleration**, **intelligent customer service**, and **creative content generation**. By combining structured reasoning with unstructured language understanding, KoreAI bridges the gap between human intent and machine execution.

### Who is KoreAI for?

- **Developers** building AI-powered applications
- **Data Scientists** experimenting with agent architectures
- **Enterprises** needing scalable, secure AI automation
- **Product Teams** wanting to embed intelligent assistants

## The Name "KoreAI"

The name **KoreAI** draws from the ancient Greek word *Kore* (κόρη), meaning "maiden" or "daughter", often a title of the goddess Persephone. Symbolically, it represents the emergence of new intelligence from the depths of data (the underworld) into the light of actionable insights. Phonetically, it also echoes **"Core AI"**, reinforcing the project's role as the central nervous system for AI agents.

## Core Concepts & Philosophy

KoreAI is built on five foundational principles:

1. **Agent Specialization** – Break complex tasks into subtasks handled by domain-specific agents (researcher, coder, analyst, etc.).
2. **Cognitive Planning** – Use advanced reasoning strategies like ReAct (Reason + Act), Tree of Thoughts, or custom state machines to decide *what* to do next.
3. **Memory & Context** – Maintain short-term conversation state (Redis) and long-term semantic memory (Qdrant vector DB) to personalize and ground interactions.
4. **Tool Empowerment** – Agents can call external APIs, execute code in sandboxes, query databases, and read documents – turning language into action.
5. **Observability & Governance** – Every decision, tool call, and response is logged, monitored, and auditable.

### How Is This Different from a Simple Chatbot?

| Feature | Simple Chatbot | KoreAI |
|--------|----------------|--------|
| Task Complexity | Single prompt-response | Multi-step planning |
| Memory | None or basic | Short + long term, persistent |
| Tools | None | APIs, code execution, DB queries |
| Reasoning | None | Chain-of-thought, ReAct, ToT |
| Multi-Agent | No | Yes, specialized agents collaborate |
| Security | Minimal | Sandboxed execution, prompt injection guard |
| Observability | Logs | Metrics, traces, cost tracking |

## Architecture Deep Dive

KoreAI follows an **event-driven, microservices-inspired** architecture, where the central orchestrator communicates with agents, memory, and tools via message passing (Redis Pub/Sub or direct REST calls). The diagram below illustrates the high-level components and their interactions.

```
┌─────────────────────────────────────────────────────────────────┐
│                          USER INTERFACES                         │
│  ┌─────────────┐  ┌──────────────┐  ┌─────────────┐            │
│  │ Web Chat UI │  │  REST API    │  │   CLI Tool  │            │
│  │ (Next.js)   │  │  (Node.js)   │  │  (Python)   │            │
│  └──────┬──────┘  └──────┬───────┘  └──────┬──────┘            │
└─────────┼─────────────────┼──────────────────┼──────────────────┘
          │                 │                  │
          └─────────────────┼──────────────────┘
                            │
┌───────────────────────────┼──────────────────────────────────────┐
│                    API GATEWAY (Node.js + TypeScript)            │
│  • Authentication (JWT)                                         │
│  • Rate Limiting                                                │
│  • Input Sanitization                                           │
│  • Request Routing to Orchestrator                              │
└───────────────────────────┬──────────────────────────────────────┘
                            │ (Redis Pub/Sub or gRPC)
┌───────────────────────────┼──────────────────────────────────────┐
│               ORCHESTRATION ENGINE (Python 3.11+)                │
│                                                                  │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                   TASK PLANNER                             │  │
│  │  • ReAct (Reason-Act) loop                                │  │
│  │  • Tree of Thoughts (exploration)                         │  │
│  │  • Custom finite-state machines                           │  │
│  └───────────────────────────────────────────────────────────┘  │
│                                                                  │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                    AGENT POOL                              │  │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐     │  │
│  │  │Researcher│ │  Coder   │ │  Analyst │ │ Creative │ ... │  │
│  │  │ Agent    │ │  Agent   │ │  Agent   │ │  Agent   │     │  │
│  │  └─────┬────┘ └─────┬────┘ └─────┬────┘ └─────┬────┘     │  │
│  └────────┼─────────────┼─────────────┼─────────────┼────────┘  │
│           │             │             │             │            │
└───────────┼─────────────┼─────────────┼─────────────┼────────────┘
            │             │             │             │
┌───────────┼─────────────┼─────────────┼─────────────┼────────────┐
│           │             │   SHARED SERVICES           │            │
│  ┌────────▼────────┐ ┌──▼──────────┐ ┌▼────────────┐▼──────────┐ │
│  │   Memory Store  │ │Vector Store │ │ Tool Registry│ LLM Pool │ │
│  │     (Redis)     │ │  (Qdrant)   │ │  (Python)   │ (OpenAI, │ │
│  │ - Conversation  │ │ - Documents │ │ - REST APIs │  Anthro…) │ │
│  │ - User Facts    │ │ - Embeddings│ │ - Code Exec │           │ │
│  └─────────────────┘ └─────────────┘ └─────────────┘───────────┘ │
└──────────────────────────────────────────────────────────────────┘
```

### Component Details

- **API Gateway (Node.js)**: Handles HTTP/WebSocket connections, validates JWT, enforces rate limits, and routes messages to the Orchestrator. Built with Express/Fastify for speed.
- **Orchestrator (Python)**: The brain. It receives a user intent, decomposes it into a task plan, spawns agents, monitors progress, and synthesizes final output.
- **Agent Pool**: A dynamic set of worker agents, each with a defined role (system prompt), available tools, and memory access. Agents can run concurrently.
- **Memory Store (Redis)**: Ultra-fast key-value store for conversation history and user-specific facts. Supports TTL and pub/sub for real-time updates.
- **Vector Store (Qdrant)**: Stores document embeddings for semantic search. Enables RAG (Retrieval-Augmented Generation) to ground answers in proprietary data.
- **Tool Registry**: A catalog of callable functions (Python/JS) with JSON schemas describing their inputs/outputs. Agents select tools based on task requirements.
- **LLM Pool**: Abstraction over multiple LLM providers, supporting failover, cost optimization, and model routing.

## How It Works: A Complete Flow

Let's walk through a complex user request: *"Analyze our Q2 sales data, create a chart, and write a summary email."*

1. **User Input**: Received by Next.js UI or REST API.
2. **Gateway**: Authenticates user, logs request, publishes to `orchestrator:new_task` channel.
3. **Orchestrator**:
   - Retrieves user context from Redis (past interactions, preferences).
   - Queries Qdrant for relevant documents (e.g., "sales report template").
   - Plans: "I need to (a) fetch sales data, (b) run statistical analysis, (c) generate chart, (d) compose email."
   - Assigns tasks:
     - **Analyst Agent**: fetch DB, compute stats.
     - **Coder Agent**: write Python script for chart.
     - **Creative Agent**: draft email text.
4. **Agents Execution**:
   - Analyst Agent queries PostgreSQL via tool, returns numbers.
   - Coder Agent generates matplotlib code, executes in sandbox, returns image URL.
   - Creative Agent uses LLM to craft email, incorporating stats and chart link.
5. **Synthesis**: Orchestrator collects outputs, formats final response (text + image + email).
6. **Response**: API returns JSON with structured content or streams via SSE.
7. **Post-Processing**: Conversation saved to Redis, new facts updated, cost logged.

This entire process happens in seconds, with parallelism where possible.

## Tech Stack & Polyglot Strategy

KoreAI deliberately uses multiple languages, each chosen for its strengths:

| Component | Language | Key Libraries | Justification |
|-----------|----------|---------------|---------------|
| **API Gateway** | TypeScript (Node.js) | Express, Fastify, ws | Non-blocking I/O handles thousands of concurrent connections. TypeScript adds type safety. |
| **Orchestrator & Agents** | Python 3.11+ | LangChain, asyncio, Pydantic | Python dominates AI/ML. LangChain provides agent/tool abstractions. asyncio enables concurrent agent execution. |
| **Web UI** | TypeScript (Next.js, React) | TailwindCSS, SWR | SSR for performance, React ecosystem, shared types with API. |
| **Memory & Cache** | Redis | ioredis (Node), redis-py | Sub-millisecond latency, perfect for conversation state and pub/sub messaging. |
| **Vector Database** | Qdrant | qdrant-client | Rust-based, high recall, payload filtering for multi-tenancy. |
| **Relational DB** | PostgreSQL | Prisma (Node), SQLAlchemy (Python) | User accounts, logs, billing. JSONB for flexible agent configs. |
| **Sandbox** | gVisor / Firecracker | REST API | Secure code execution for untrusted agent-generated code. |
| **Containerization** | Docker, Kubernetes | - | Consistent deploys, auto-scaling agent pools. |
| **Monitoring** | Prometheus, Grafana, OpenTelemetry | - | Metrics, traces, cost tracking. |

### Why Not Just Python?

While Python is great for AI, it's not ideal for high-concurrency APIs or real-time UIs. Node.js excels at I/O-heavy workloads, and Next.js provides a modern frontend experience. This **polyglot** approach ensures each layer uses the best tool for the job.

## Features

### 🧠 Intelligent Multi-Agent System
- **Role-based agents**: Researcher, Coder, Analyst, Creative, Custom.
- **Dynamic task decomposition**: Automatic or manual planning.
- **Parallel execution**: Multiple agents work simultaneously.

### 📚 Advanced RAG (Retrieval-Augmented Generation)
- **Document ingestion**: PDF, HTML, Markdown, CSV.
- **Chunking strategies**: Recursive, semantic, fixed-size.
- **Hybrid search**: Dense (embeddings) + sparse (BM25) retrieval.

### 🔌 Universal Tool Integration
- **REST API tools**: Auto-generate tool schemas from OpenAPI specs.
- **SQL tools**: Safe query generation with read-only permissions.
- **Code execution**: Python/JS sandboxes with resource limits.
- **Custom functions**: Decorators to expose Python functions as tools.

### 💾 Memory Management
- **Short-term**: Conversation buffer (Redis list).
- **Long-term**: User facts, preferences (Redis hashes + Qdrant).
- **Summarization**: Automatic conversation compression to avoid token overflow.

### 🛡️ Security & Governance
- **Prompt injection detection**: Input sanitization and LLM guardrails.
- **Sandboxed execution**: gVisor for untrusted code.
- **Audit trail**: Every action logged with user and agent ID.
- **Data isolation**: Multi-tenant support with strict separation.

### 📊 Observability & Analytics
- **Cost tracking**: Token usage and API costs per user/agent.
- **Agent performance metrics**: Success rate, latency, tool usage.
- **Live tracing**: Visualize agent decision paths.

## Project Structure

```
KoreAI/
├── api/                          # Node.js API Gateway
│   ├── src/
│   │   ├── controllers/          # Request handlers
│   │   ├── middleware/            # Auth, rate limit, sanitize
│   │   ├── services/             # Business logic, Redis, DB
│   │   ├── websocket/            # Real-time communication
│   │   └── utils/
│   ├── prisma/                   # Database schema
│   ├── tests/
│   ├── Dockerfile
│   └── package.json
│
├── core/                         # Python AI Engine
│   ├── orchestrator/             # Task planning and coordination
│   │   ├── planner.py
│   │   ├── strategies/           # ReAct, ToT, etc.
│   │   └── dispatcher.py
│   ├── agents/                   # Agent definitions
│   │   ├── base.py               # Abstract agent class
│   │   ├── researcher.py
│   │   ├── coder.py
│   │   ├── analyst.py
│   │   └── custom/
│   ├── memory/                   # Short & long term memory
│   │   ├── redis_store.py
│   │   └── vector_store.py
│   ├── tools/                    # Tool registry and implementations
│   │   ├── registry.py
│   │   ├── api_tool.py
│   │   ├── sql_tool.py
│   │   └── code_executor.py
│   ├── llm/                      # LLM abstraction layer
│   │   ├── providers/
│   │   └── router.py
│   ├── sandbox/                  # Secure code execution
│   ├── config/
│   ├── tests/
│   ├── Dockerfile
│   └── requirements.txt
│
├── web/                          # Next.js Frontend
│   ├── components/
│   │   ├── Chat/
│   │   ├── Dashboard/
│   │   └── Admin/
│   ├── pages/
│   ├── hooks/
│   ├── public/
│   ├── styles/
│   ├── Dockerfile
│   └── package.json
│
├── infrastructure/
│   ├── docker-compose.dev.yml
│   ├── docker-compose.prod.yml
│   ├── k8s/
│   │   ├── api-deployment.yaml
│   │   ├── core-deployment.yaml
│   │   └── ingress.yaml
│   └── terraform/
│       ├── main.tf
│       └── variables.tf
│
├── docs/
│   ├── architecture.md
│   ├── api.md
│   ├── agents.md
│   └── deployment.md
├── .env.example
├── README.md
└── LICENSE
```

## Installation & Quick Start

### Prerequisites

- **Python 3.11+** with `pip`
- **Node.js 20+** with `npm`
- **Docker & Docker Compose** (for Redis, Qdrant, PostgreSQL)
- **LLM Provider Keys** (OpenAI, etc.)

### 1. Clone & Setup Infrastructure

```bash
git clone https://github.com/kauandias747474-hue/KoreAI.git
cd KoreAI
cp .env.example .env   # Edit with your keys
docker-compose -f infrastructure/docker-compose.dev.yml up -d
```

### 2. Start Python Core

```bash
cd core
python -m venv venv
source venv/bin/activate  # Linux/Mac
# venv\Scripts\activate   # Windows
pip install -r requirements.txt
python main.py
```

### 3. Start Node.js API

```bash
cd api
npm install
npm run dev
```

### 4. Start Web UI

```bash
cd web
npm install
npm run dev
```

Visit `http://localhost:3000` and start chatting with your AI agents!

## Configuration Guide

Full `.env` reference:

```env
# --- LLM Providers ---
OPENAI_API_KEY=sk-...
ANTHROPIC_API_KEY=...
GOOGLE_API_KEY=...
# Model routing
DEFAULT_MODEL=gpt-4o
FALLBACK_MODEL=claude-3-opus

# --- Database ---
DATABASE_URL=postgresql://koreai:koreai@localhost:5432/koreai
REDIS_URL=redis://localhost:6379
QDRANT_URL=http://localhost:6333

# --- Agent Settings ---
MAX_AGENT_TOKENS=4096
AGENT_TIMEOUT_SEC=120
MAX_PARALLEL_AGENTS=5
MEMORY_TTL=3600

# --- Security ---
SANDBOX_ENABLED=true
MAX_CODE_EXEC_TIME_MS=5000

# --- Observability ---
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
LOG_LEVEL=info
```

## Usage Examples

### Python SDK (within your own scripts)

```python
from koreai import KoreAI

client = KoreAI()

# Simple question
answer = client.ask("Explain quantum computing in one paragraph")
print(answer)

# Multi-step task
task = """
1. Query our database for sales last month.
2. Compare to previous month.
3. Create a bar chart image.
4. Write an executive summary.
"""
response = client.solve(task, stream=True)
for chunk in response:
    print(chunk, end="")
```

### REST API

```bash
curl -X POST http://localhost:3001/api/v1/chat \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT" \
  -d '{
    "message": "Find the latest research on CRISPR and summarize key points.",
    "conversation_id": "optional",
    "stream": false
  }'
```

Response:

```json
{
  "id": "msg_abc123",
  "role": "assistant",
  "content": "Based on recent publications...",
  "sources": [
    {"title": "Nature CRISPR Review 2025", "url": "..."}
  ],
  "tokens_used": 1250,
  "cost": 0.02
}
```

### WebSocket (Real-time)

```javascript
const ws = new WebSocket('ws://localhost:3001/ws');
ws.send(JSON.stringify({ message: "Hello!", token: "..." }));
ws.onmessage = (event) => {
  console.log(JSON.parse(event.data));
};
```

## Agent Design System

Every agent inherits from `BaseAgent` and implements:

- `system_prompt`: Instructions defining role and constraints.
- `tools`: List of callable tools.
- `memory`: Access to short/long term memory.
- `execute(task) -> AgentResult`: Main method.

Example custom agent:

```python
from koreai.agents.base import BaseAgent, Tool

class LegalAdvisor(BaseAgent):
    system_prompt = """You are a legal advisor. Always cite relevant laws.
    Never give financial advice."""
    
    tools = [Tool(name="search_legislation", func=search_law_db)]
    
    async def execute(self, task):
        # Custom logic before LLM call
        research = await self.call_tool("search_legislation", task.query)
        response = await self.llm.generate(
            prompt=f"Based on {research}, answer: {task}"
        )
        return response
```

## Memory & RAG Pipeline

### Short-Term Memory (Redis)

- **Conversation buffer**: Stored as list with TTL.
- **User facts**: Hashes for preferences, e.g., `user:123:facts -> {name: "Alice", role: "Manager"}`.

### Long-Term Memory (Qdrant)

- Documents indexed after preprocessing (chunking, embedding).
- At query time, orchestrator fetches top-k relevant chunks.
- Payload filtering ensures multi-tenant isolation.

### Example: Adding Documents

```python
await vector_store.add_documents(
    documents=["Q2 report.pdf content..."],
    metadata={"user_id": "123", "category": "finance"}
)
```

## Tool Integration Framework

Tools are defined with a name, description, and JSON schema for parameters. They can be:

- **Python functions**: Decorated with `@tool`.
- **REST endpoints**: Described via OpenAPI import.
- **Database queries**: Auto-generated from schema.

Example:

```python
@tool
async def get_stock_price(symbol: str) -> dict:
    """Fetch current stock price for a symbol."""
    async with httpx.AsyncClient() as client:
        resp = await client.get(f"https://api.example.com/stock/{symbol}")
        return resp.json()
```

Agents decide which tool to call based on the task and tool descriptions.

## Security Architecture

| Threat | Mitigation |
|--------|------------|
| Prompt Injection | Input sanitization, LLM guardrails, regex filters. |
| Malicious Code Execution | Sandboxed execution (gVisor), resource limits, timeouts. |
| Data Leakage | Multi-tenant isolation in Redis/Qdrant, RLS in PostgreSQL. |
| API Abuse | JWT authentication, rate limiting by user/API key. |
| Model Theft | API gateway hides raw model endpoints, caching public responses. |

## Performance & Scalability

- **Async Python**: `asyncio` for parallel agent execution.
- **Redis Pub/Sub**: Decouples orchestrator and agents.
- **LLM Caching**: Identical requests served from Redis cache.
- **Horizontal Scaling**: Stateless agents can be replicated behind a load balancer.

Benchmarks (preliminary):

- Single agent task: ~2-4s
- 5-agent parallel task: ~3-6s
- 1000 concurrent users: 50ms API response time (non-streaming)

## Testing & Quality Assurance

```bash
# Python unit tests
cd core && pytest --cov

# Node.js API tests
cd api && npm test

# E2E tests
cd tests/e2e && npx playwright test
```

## Deploy & DevOps

### Docker Compose (Production)

```bash
docker-compose -f infrastructure/docker-compose.prod.yml up -d
```

### Kubernetes

```bash
kubectl apply -f infrastructure/k8s/
```

### Terraform (AWS/GCP)

```bash
cd infrastructure/terraform
terraform init && terraform apply
```

---
# 🤖 KoreAI — Explicação do Diagrama de Arquitetura



## 🇧🇷 Português

### 📊 Diagrama de Arquitetura do Sistema

O diagrama abaixo apresenta a arquitetura completa da plataforma **KoreAI**, uma plataforma multi-agente de orquestração de LLMs. Ele ilustra as interfaces de usuário, o API Gateway, o motor de orquestração, o pool de agentes especializados, os serviços compartilhados e a infraestrutura subjacente.

```mermaid
graph TB
    subgraph "KoreAI - Plataforma Multi-Agente"
        
        subgraph "Interfaces de Usuário"
            WebUI["🖥️ Web Chat UI<br/>(Next.js)"]
            REST_API["🔗 REST API<br/>(Clientes HTTP)"]
            CLI["⌨️ CLI Tool<br/>(Python)"]
            SDK["📦 Python SDK<br/>(koreai lib)"]
        end

        subgraph "API Gateway - Node.js + TypeScript"
            Auth["🔐 Autenticação<br/>(JWT)"]
            RateLimit["⏱️ Rate Limiting"]
            Sanitize["🧹 Input Sanitization"]
            Router["🚦 Request Router"]
        end

        subgraph "Orchestration Engine - Python 3.11+"
            Planner["🧠 Task Planner<br/>(ReAct / Tree of Thoughts / FSM)"]
            Dispatcher["📋 Dispatcher<br/>(Spawn & Monitor Agents)"]
            
            subgraph "Agent Pool"
                Researcher["🔍 Researcher Agent"]
                Coder["💻 Coder Agent"]
                Analyst["📊 Analyst Agent"]
                Creative["✨ Creative Agent"]
                Custom["⚙️ Custom Agent"]
            end
        end

        subgraph "Shared Services"
            Memory["💾 Memory Store<br/>(Redis)"]
            Vector["🧬 Vector Store<br/>(Qdrant)"]
            Tools["🔧 Tool Registry<br/>(Python/JS Functions)"]
            LLMPool["🤖 LLM Pool<br/>(OpenAI / Anthropic / Google)"]
            Sandbox["🛡️ Sandbox<br/>(gVisor / Firecracker)"]
        end

        subgraph "Data & Infrastructure"
            PostgreSQL[("🗄️ PostgreSQL<br/>(Users, Logs, Billing)")]
            Docker["🐳 Docker + Kubernetes"]
            Monitoring["📊 Prometheus + Grafana<br/>+ OpenTelemetry"]
        end

        WebUI --> Auth
        REST_API --> Auth
        CLI --> Auth
        SDK --> Auth
        Auth --> RateLimit
        RateLimit --> Sanitize
        Sanitize --> Router
        Router --> Planner
        Planner --> Dispatcher
        Dispatcher --> Researcher
        Dispatcher --> Coder
        Dispatcher --> Analyst
        Dispatcher --> Creative
        Dispatcher --> Custom

        Researcher --> Memory
        Researcher --> Vector
        Researcher --> Tools
        Coder --> Sandbox
        Coder --> Tools
        Analyst --> PostgreSQL
        Analyst --> Tools
        Creative --> LLMPool
        Custom --> Tools
        Custom --> LLMPool

        Memory --> PostgreSQL
        Vector --> PostgreSQL
        Tools --> PostgreSQL
        LLMPool --> Monitoring
        Router --> Monitoring
        Dispatcher --> Monitoring
    end
```

### 🧩 Explicação Detalhada

#### 1. Interfaces de Usuário (Topo)

O usuário pode interagir com o KoreAI de quatro maneiras principais:

- **Web Chat UI (Next.js)**: Uma interface de chat moderna e responsiva, com renderização no servidor (SSR) para performance.
- **REST API**: Permite que aplicações externas enviem comandos e recebam respostas estruturadas em JSON.
- **CLI Tool (Python)**: Para desenvolvedores e administradores, permite interagir com a plataforma via terminal.
- **Python SDK**: Uma biblioteca (`koreai`) que abstrai as chamadas à API e oferece uma experiência nativa em Python.

Todas as entradas são unificadas pelo **API Gateway**.

#### 2. API Gateway (Node.js + TypeScript)

Esta camada atua como a porta de entrada segura e controlada da plataforma:

- **Autenticação (JWT)**: Valida tokens de acesso, garantindo que apenas usuários autorizados utilizem o sistema.
- **Rate Limiting**: Protege contra abusos, limitando o número de requisições por usuário ou IP.
- **Input Sanitization**: Limpa e normaliza as entradas do usuário para prevenir injeção de prompts maliciosos.
- **Request Router**: Direciona as requisições válidas para o **Orchestration Engine**.

#### 3. Orchestration Engine (Python 3.11+)

O cérebro da plataforma, responsável por coordenar a resolução de tarefas complexas:

- **Task Planner**: Analisa a intenção do usuário e a decompõe em um plano de ação. Utiliza estratégias avançadas de raciocínio:
  - **ReAct** (Reason + Act): alterna entre pensar e agir.
  - **Tree of Thoughts**: explora múltiplos caminhos de solução em paralelo.
  - **Finite State Machines (FSM)**: para fluxos bem definidos e determinísticos.
- **Dispatcher**: Distribui as subtarefas para os agentes apropriados e monitora sua execução, podendo executar vários agentes simultaneamente.

#### 4. Agent Pool

Agentes especializados, cada um com um papel, ferramentas e acesso a memória:

- **Researcher Agent**: Busca informações em documentos (via RAG) e fontes externas.
- **Coder Agent**: Gera e executa código em sandbox seguro.
- **Analyst Agent**: Consulta bancos de dados, processa dados e gera insights.
- **Creative Agent**: Redige textos, e-mails, e conteúdo criativo.
- **Custom Agent**: Agentes definidos pelo usuário para tarefas específicas.

Cada agente pode acessar os serviços compartilhados conforme necessário.

#### 5. Shared Services (Serviços Compartilhados)

Infraestrutura de suporte acessível a todos os agentes:

- **Memory Store (Redis)**: Mantém o histórico de conversas (curto prazo) e fatos do usuário (preferências, contexto). Rápido, com TTL e suporte a pub/sub.
- **Vector Store (Qdrant)**: Armazena embeddings de documentos para busca semântica. Essencial para o funcionamento do RAG (Retrieval-Augmented Generation), permitindo que os agentes fundamentem suas respostas em dados proprietários.
- **Tool Registry**: Catálogo de funções invocáveis (APIs REST, consultas SQL, execução de código) com esquemas JSON que descrevem entradas e saídas.
- **LLM Pool**: Abstração sobre múltiplos provedores de LLM (OpenAI, Anthropic, Google). Oferece failover automático, otimização de custos e roteamento de modelos.
- **Sandbox (gVisor/Firecracker)**: Ambiente isolado para execução segura de código não confiável gerado pelos agentes.

#### 6. Data & Infrastructure

- **PostgreSQL**: Banco de dados relacional para informações estruturadas: contas de usuário, logs de auditoria, faturamento, configurações.
- **Docker + Kubernetes**: Containerização para deploys consistentes e orquestração para escalabilidade automática dos agentes.
- **Prometheus + Grafana + OpenTelemetry**: Stack de observabilidade para métricas, rastreamento distribuído e monitoramento de custos.

### 🔄 Fluxo de uma Requisição Típica

1. O usuário envia uma mensagem pelo chat, API, CLI ou SDK.
2. O API Gateway autentica, limita a taxa e sanitiza a entrada.
3. O Request Router encaminha para o Task Planner.
4. O Planner decompõe a tarefa e o Dispatcher aloca os agentes necessários.
5. Os agentes executam suas tarefas, consultando memória, vetores, ferramentas e LLMs.
6. Os resultados são coletados e sintetizados pelo Orchestrator.
7. A resposta é enviada de volta ao usuário pelo mesmo canal de entrada.
8. Todo o processo é registrado em logs, métricas e rastros para auditoria e otimização.

---

## 🇺🇸 English

### 📊 System Architecture Diagram

The diagram below presents the complete architecture of the **KoreAI** platform, a multi-agent LLM orchestration system. It illustrates user interfaces, the API Gateway, the orchestration engine, the specialized agent pool, shared services, and underlying infrastructure.

```mermaid
graph TB
    subgraph "KoreAI - Multi-Agent Platform"
        
        subgraph "User Interfaces"
            WebUI["🖥️ Web Chat UI<br/>(Next.js)"]
            REST_API["🔗 REST API<br/>(HTTP Clients)"]
            CLI["⌨️ CLI Tool<br/>(Python)"]
            SDK["📦 Python SDK<br/>(koreai lib)"]
        end

        subgraph "API Gateway - Node.js + TypeScript"
            Auth["🔐 Authentication<br/>(JWT)"]
            RateLimit["⏱️ Rate Limiting"]
            Sanitize["🧹 Input Sanitization"]
            Router["🚦 Request Router"]
        end

        subgraph "Orchestration Engine - Python 3.11+"
            Planner["🧠 Task Planner<br/>(ReAct / Tree of Thoughts / FSM)"]
            Dispatcher["📋 Dispatcher<br/>(Spawn & Monitor Agents)"]
            
            subgraph "Agent Pool"
                Researcher["🔍 Researcher Agent"]
                Coder["💻 Coder Agent"]
                Analyst["📊 Analyst Agent"]
                Creative["✨ Creative Agent"]
                Custom["⚙️ Custom Agent"]
            end
        end

        subgraph "Shared Services"
            Memory["💾 Memory Store<br/>(Redis)"]
            Vector["🧬 Vector Store<br/>(Qdrant)"]
            Tools["🔧 Tool Registry<br/>(Python/JS Functions)"]
            LLMPool["🤖 LLM Pool<br/>(OpenAI / Anthropic / Google)"]
            Sandbox["🛡️ Sandbox<br/>(gVisor / Firecracker)"]
        end

        subgraph "Data & Infrastructure"
            PostgreSQL[("🗄️ PostgreSQL<br/>(Users, Logs, Billing)")]
            Docker["🐳 Docker + Kubernetes"]
            Monitoring["📊 Prometheus + Grafana<br/>+ OpenTelemetry"]
        end

        WebUI --> Auth
        REST_API --> Auth
        CLI --> Auth
        SDK --> Auth
        Auth --> RateLimit
        RateLimit --> Sanitize
        Sanitize --> Router
        Router --> Planner
        Planner --> Dispatcher
        Dispatcher --> Researcher
        Dispatcher --> Coder
        Dispatcher --> Analyst
        Dispatcher --> Creative
        Dispatcher --> Custom

        Researcher --> Memory
        Researcher --> Vector
        Researcher --> Tools
        Coder --> Sandbox
        Coder --> Tools
        Analyst --> PostgreSQL
        Analyst --> Tools
        Creative --> LLMPool
        Custom --> Tools
        Custom --> LLMPool

        Memory --> PostgreSQL
        Vector --> PostgreSQL
        Tools --> PostgreSQL
        LLMPool --> Monitoring
        Router --> Monitoring
        Dispatcher --> Monitoring
    end
```

### 🧩 Detailed Explanation

#### 1. User Interfaces (Top)

Users can interact with KoreAI in four main ways:

- **Web Chat UI (Next.js)**: A modern, responsive chat interface with server-side rendering (SSR) for performance.
- **REST API**: Allows external applications to send commands and receive structured JSON responses.
- **CLI Tool (Python)**: For developers and administrators to interact with the platform via terminal.
- **Python SDK**: A library (`koreai`) that abstracts API calls and provides a native Python experience.

All inputs are unified by the **API Gateway**.

#### 2. API Gateway (Node.js + TypeScript)

This layer acts as the secure and controlled entry point:

- **Authentication (JWT)**: Validates access tokens, ensuring only authorized users access the system.
- **Rate Limiting**: Protects against abuse by limiting requests per user or IP.
- **Input Sanitization**: Cleans and normalizes user input to prevent malicious prompt injection.
- **Request Router**: Directs valid requests to the **Orchestration Engine**.

#### 3. Orchestration Engine (Python 3.11+)

The brain of the platform, responsible for coordinating complex tasks:

- **Task Planner**: Analyzes user intent and decomposes it into an action plan. Uses advanced reasoning strategies:
  - **ReAct** (Reason + Act): alternates between thinking and acting.
  - **Tree of Thoughts**: explores multiple solution paths in parallel.
  - **Finite State Machines (FSM)**: for well-defined, deterministic flows.
- **Dispatcher**: Distributes subtasks to the appropriate agents and monitors their execution, often running multiple agents concurrently.

#### 4. Agent Pool

Specialized agents, each with a role, tools, and memory access:

- **Researcher Agent**: Searches for information in documents (via RAG) and external sources.
- **Coder Agent**: Generates and executes code in a secure sandbox.
- **Analyst Agent**: Queries databases, processes data, and generates insights.
- **Creative Agent**: Writes texts, emails, and creative content.
- **Custom Agent**: User-defined agents for specific tasks.

Each agent can access shared services as needed.

#### 5. Shared Services

Support infrastructure accessible to all agents:

- **Memory Store (Redis)**: Maintains conversation history (short-term) and user facts (preferences, context). Fast, with TTL and pub/sub support.
- **Vector Store (Qdrant)**: Stores document embeddings for semantic search. Essential for RAG (Retrieval-Augmented Generation), allowing agents to ground answers in proprietary data.
- **Tool Registry**: Catalog of callable functions (REST APIs, SQL queries, code execution) with JSON schemas describing inputs and outputs.
- **LLM Pool**: Abstraction over multiple LLM providers (OpenAI, Anthropic, Google). Offers automatic failover, cost optimization, and model routing.
- **Sandbox (gVisor/Firecracker)**: Isolated environment for safe execution of untrusted code generated by agents.

#### 6. Data & Infrastructure

- **PostgreSQL**: Relational database for structured information: user accounts, audit logs, billing, configurations.
- **Docker + Kubernetes**: Containerization for consistent deploys and orchestration for auto-scaling agent pools.
- **Prometheus + Grafana + OpenTelemetry**: Observability stack for metrics, distributed tracing, and cost monitoring.

### 🔄 Typical Request Flow

1. The user sends a message via chat, API, CLI, or SDK.
2. The API Gateway authenticates, rate-limits, and sanitizes the input.
3. The Request Router forwards it to the Task Planner.
4. The Planner decomposes the task and the Dispatcher assigns the necessary agents.
5. Agents execute their tasks, consulting memory, vectors, tools, and LLMs.
6. Results are collected and synthesized by the Orchestrator.
7. The response is sent back to the user through the same channel.
8. The entire process is recorded in logs, metrics, and traces for auditing and optimization.

## Contributing

We welcome contributions! Please see `CONTRIBUTING.md` for details on coding standards, pull request process, and code of conduct.

## Roadmap

### Short-term (Q3 2026)
- [ ] Voice interface integration (Whisper + TTS)
- [ ] Agent marketplace for sharing custom agents
- [ ] Improved ReAct debugging UI

### Long-term (2027)
- [ ] Multi-modal agents (image, video understanding)
- [ ] Federated learning for privacy-preserving fine-tuning
- [ ] Autonomous agent swarms for complex simulations

## FAQ

**Q: Can I use local models?**
A: Yes! KoreAI supports Ollama and llama-cpp-python. Set `LLM_PROVIDER=ollama` in your `.env`.

**Q: Is multi-tenancy supported?**
A: Absolutely. User IDs, conversation isolation, and database RLS are built-in.

**Q: How do I add my own tools?**
A: Create a Python function, decorate with `@tool`, and register it. It becomes instantly available to agents.

**Q: What does it cost to run?**
A: You only pay for the LLM provider tokens (e.g., OpenAI API). Self-hosting the platform is free. Redis, Postgres, Qdrant can run on modest VMs.

## License

MIT License © 2026 KoreAI


---

#  KoreAI — Futuras Adições / Future Additions

[🇧🇷 Versão em Português](#-planeamento-de-expansão-da-plataforma-koreai)  
[🇺🇸 English Version](#-koreai-platform-expansion-roadmap)

---

## 🇧🇷 Planeamento de Expansão da Plataforma KoreAI

Este documento descreve as **novas funcionalidades e melhorias planeadas** para a plataforma KoreAI, organizadas em 8 eixos estratégicos. Cada eixo detalha adições que elevarão a inteligência, segurança, escalabilidade e valor do sistema multi‑agente.

### 1. 🧠 Expansão da Inteligência dos Agentes
- **Perfis cognitivos adaptativos** – agentes ajustam o estilo de comunicação, profundidade técnica e tom emocional com base num modelo de preferências do utilizador armazenado no Redis.
- **Loop de auto‑melhoria com feedback humano (RLHF)** – recolha automática de indicadores de qualidade (ex.: tempo de resolução, avaliações explícitas) para afinar modelos via *fine‑tuning* (LoRA/QLoRA) em ambiente fechado.
- **Agente revisor interno** – um agente de bastidor que verifica factos, coerência e conformidade ética antes da resposta final ser entregue, reduzindo alucinações.
- **Memória emocional** – histórico do estado emocional do utilizador para interações mais empáticas (ex.: detetar frustração e responder com mais paciência).

### 2. 📚 Melhorias no RAG e na Memória de Longo Prazo
- **Memória episódica vetorial** – guardar não apenas factos, mas também “episódios” completos de conversas anteriores, permitindo que o agente recupere contexto: “Na semana passada, falámos sobre X, e decidiste Y”.
- **Graph RAG** – utilização de grafos de conhecimento (Neo4j integrado ao Qdrant) para representar relações complexas entre documentos, entidades e conceitos.
- **Cache semântico inteligente** – embeddings frequentemente usados são armazenados em Redis para evitar recálculo e reduzir latência.
- **Indexação incremental de fontes externas** – suporte a feeds RSS, atualizações de bases de dados e repositórios de documentos, com re‑indexação automática.

### 3. 🔧 Ferramentas e Ações Mais Poderosas
- **Browser Agent** – agente com navegador headless (Playwright/Selenium) para extrair dados dinâmicos, interagir com formulários, autenticar‑se em portais e gerar relatórios baseados em informação da web.
- **Workflow Engine integrada** – capacidade de orquestrar pipelines complexos (ex.: “Extrai dados do S3, limpa, treina modelo e envia relatório por e‑mail”) utilizando motores como Temporal ou Prefect, expostos como ferramentas.
- **Registo automático de APIs externas** – dado um ficheiro OpenAPI/Swagger, o sistema gera automaticamente a *tool definition* e a disponibiliza no Tool Registry, sem código adicional.
- **Multi‑ferramenta condicional** – agentes podem combinar várias ferramentas numa única chamada de alto nível, com gestão automática de dependências.

### 4. 🎤🧠📷 Multi‑Modalidade e Interfaces Avançadas
- **Entrada/saída de voz** – integração total com Whisper (STT) e modelos TTS de última geração (ex.: ElevenLabs, Piper) com suporte a pistas emocionais.
- **Visão computacional** – processamento de imagens, gráficos e documentos digitalizados para extração de dados estruturados, reconhecimento de objetos e OCR avançado.
- **Interpretação de áudio** – análise de tom de voz, deteção de emoções e níveis de urgência para adaptar a resposta do agente.
- **Saída em realidade aumentada** – protótipo de agente que gera objetos 3D ou sobreposições para headsets, útil em simulações industriais e formação.

### 5. 🐝 Colaboração Multi‑Agente e Enxames
- **Debate entre agentes** – dois ou mais agentes argumentam sobre um tópico, um moderador avalia a qualidade dos argumentos e sintetiza uma conclusão ponderada (melhora precisão analítica).
- **Enxames para processamento massivo** – dezenas de agentes leves trabalham em paralelo para analisar grandes volumes de documentos, imagens ou registos, coordenados por um *swarm manager*.
- **Agentes proativos** – agentes de *background* monitorizam fontes (ex.: calendários, mercados financeiros, redes sociais) e iniciam conversas quando detetam eventos relevantes.
- **Hierarquias dinâmicas** – criação de hierarquias de agentes em tempo de execução, onde um agente “líder” pode delegar sub‑tarefas e escalar problemas.

### 6. 🛡️ Segurança e Governança Avançadas
- **RBAC para ferramentas e agentes** – cada utilizador/tenant define quais ferramentas e tipos de agente estão disponíveis, com políticas granulares.
- **Execução em enclaves confidenciais** – suporte a ambientes de execução seguros (Intel SGX, AWS Nitro Enclaves) para processamento de dados altamente sensíveis sem exposição ao operador da plataforma.
- **Guardião ético e de conformidade** – um agente separado que audita todas as saídas contra políticas de uso aceitável, regulações (GDPR, HIPAA) e vieses.
- **Criptografia ponta‑a‑ponta para ferramentas** – possibilidade de encriptar parâmetros e resultados de ferramentas críticas, usando chaves geridas pelo cliente.

### 7. 📊 Observabilidade, Escalabilidade e DevOps
- **Agent Performance Score (APS)** – métrica composta baseada em precisão, latência, custo e feedback do utilizador, usada para seleção automática de agentes e modelos.
- **Multi‑cloud e híbrido** – templates Terraform para Azure, GCP e ambientes on‑premise; suporte a execução de agentes em edge com modelos locais (Ollama, vLLM) para baixa latência.
- **Marketplace API** – API pública para que terceiros submetam agentes e ferramentas, com testes automáticos de segurança e desempenho antes da publicação.
- **Dashboard avançado para developers** – visualização em tempo real da árvore de decisão dos agentes, execução de ferramentas, tokens gastos e tracing distribuído completo.

### 8. 💰 Monetização e Produto
- **Planos de subscrição modulares** – pacotes baseados no número de agentes especializados, volume de documentos indexados e chamadas a ferramentas premium.
- **Loja de agentes premium** – agentes certificados por parceiros (ex.: análise financeira, suporte jurídico) comercializados como add‑ons.
- **White‑label** – empresas podem incorporar a plataforma KoreAI nos seus produtos, personalizando agentes, memória e interfaces com a própria marca.
- **Insights e exportação de relatórios** – painel de produtividade que mostra tempo poupado, tarefas concluídas e ROI estimado, com possibilidade de exportar relatórios executivos gerados pelos agentes.

---

## 🇺🇸 KoreAI Platform Expansion Roadmap

This document outlines the **planned new features and improvements** for the KoreAI platform, grouped into 8 strategic pillars. Each pillar details additions that will raise the intelligence, security, scalability, and value of the multi‑agent system.

### 1. 🧠 Agent Intelligence Expansion
- **Adaptive cognitive profiles** – agents adjust communication style, technical depth, and emotional tone based on a user preference model stored in Redis.
- **Self‑improvement loop with human feedback (RLHF)** – automatic collection of quality signals (e.g., resolution time, explicit ratings) to fine‑tune models via LoRA/QLoRA in a closed environment.
- **Internal reviewer agent** – a behind‑the‑scenes agent that checks facts, coherence, and policy compliance before the final answer is delivered, reducing hallucinations.
- **Emotional memory** – history of the user’s emotional state for more empathetic interactions (e.g., detecting frustration and responding with greater patience).

### 2. 📚 RAG & Long‑term Memory Upgrades
- **Episodic vector memory** – store not only facts but also complete conversational “episodes”, enabling recall like: “Last week we talked about X, and you decided Y.”
- **Graph RAG** – integration with knowledge graphs (Neo4j + Qdrant) to represent complex relationships among documents, entities, and concepts.
- **Intelligent semantic cache** – frequently used embeddings are cached in Redis to avoid recomputation and reduce latency.
- **Incremental external source indexing** – support for RSS feeds, database change streams, and live document repositories with automatic re‑indexing.

### 3. 🔧 More Powerful Tools & Actions
- **Browser Agent** – headless browser agent (Playwright/Selenium) for dynamic data extraction, form interaction, portal authentication, and web‑based report generation.
- **Integrated workflow engine** – ability to orchestrate complex pipelines (e.g., “Pull CSV from S3, clean, train model, email report”) using engines like Temporal or Prefect, exposed as tools.
- **Auto‑registration of external APIs** – given an OpenAPI/Swagger spec, the system automatically generates the tool definition and registers it in the Tool Registry, no extra code required.
- **Conditional multi‑tool** – agents can combine several tools into a single high‑level call with automatic dependency management.

### 4. 🎤🧠📷 Multi‑Modality & Advanced Interfaces
- **Voice input/output** – full integration with Whisper (STT) and state‑of‑the‑art TTS (e.g., ElevenLabs, Piper) with support for emotional cues.
- **Computer vision** – image, chart, and scanned‑document processing for structured data extraction, object recognition, and advanced OCR.
- **Audio interpretation** – tone‑of‑voice analysis, emotion detection, and urgency levels to adapt the agent’s response.
- **Augmented reality output** – prototype agent that generates 3D objects or overlays for headsets, useful in industrial simulations and training.

### 5. 🐝 Multi‑Agent Collaboration & Swarms
- **Agent debate** – two or more agents argue a topic; a moderator evaluates the quality of arguments and synthesizes a well‑reasoned conclusion (improves analytical accuracy).
- **Massive processing swarms** – dozens of lightweight agents work in parallel to analyze large volumes of documents, images, or logs, coordinated by a swarm manager.
- **Proactive agents** – background agents monitor sources (calendars, financial markets, social media) and initiate conversations when relevant events are detected.
- **Dynamic hierarchies** – runtime creation of agent hierarchies, where a “leader” agent can delegate sub‑tasks and escalate issues.

### 6. 🛡️ Advanced Security & Governance
- **RBAC for tools and agents** – each user/tenant defines which tools and agent types are available, with granular access policies.
- **Confidential enclave execution** – support for secure execution environments (Intel SGX, AWS Nitro Enclaves) to process highly sensitive data without exposing it to the platform operator.
- **Ethics & compliance guardian** – a separate agent that audits every output against acceptable use policies, regulations (GDPR, HIPAA), and bias detection.
- **End‑to‑end tool encryption** – ability to encrypt parameters and results of critical tools using customer‑managed keys.

### 7. 📊 Observability, Scalability & DevOps
- **Agent Performance Score (APS)** – composite metric based on accuracy, latency, cost, and user feedback, used for automatic agent and model selection.
- **Multi‑cloud & hybrid** – Terraform templates for Azure, GCP, and on‑premise deployments; support for edge agents running local models (Ollama, vLLM) for low latency.
- **Marketplace API** – public API for third parties to submit agents and tools, with automated security and performance testing before publication.
- **Advanced developer dashboard** – real‑time visualization of agent decision trees, tool execution, token consumption, and full distributed tracing.

### 8. 💰 Monetization & Product
- **Modular subscription plans** – packages based on the number of specialized agents, volume of indexed documents, and premium tool calls.
- **Premium agent store** – partner‑certified agents (e.g., financial analysis, legal support) sold as add‑ons.
- **White‑label** – enterprises can embed the KoreAI platform into their own products, customizing agents, memory, and interfaces with their brand.
- **Insights & report export** – productivity dashboard showing time saved, tasks completed, and estimated ROI, with the ability to export executive reports generated by the agents.

---

Este documento descreve as novas funcionalidades e melhorias planeadas para a plataforma **KoreAI**, organizadas em **8 eixos estratégicos**. Cada eixo detalha adições que elevarão a inteligência, segurança, escalabilidade e valor do sistema multi‑agente.

### 1. 🧠 Expansão da Inteligência dos Agentes

- **Perfis cognitivos adaptativos** – agentes ajustam o estilo de comunicação, profundidade técnica e tom emocional com base num modelo de preferências do utilizador armazenado no Redis.
- **Loop de auto‑melhoria com feedback humano (RLHF)** – recolha automática de indicadores de qualidade (ex.: tempo de resolução, avaliações explícitas) para afinar modelos via fine‑tuning (LoRA/QLoRA) em ambiente fechado.
- **Agente revisor interno** – um agente de bastidor que verifica factos, coerência e conformidade ética antes da resposta final ser entregue, reduzindo alucinações.
- **Memória emocional** – histórico do estado emocional do utilizador para interações mais empáticas (ex.: detetar frustração e responder com mais paciência).

### 2. 📚 Melhorias no RAG e na Memória de Longo Prazo

- **Memória episódica vetorial** – guardar não apenas factos, mas também “episódios” completos de conversas anteriores, permitindo que o agente recupere contexto: “Na semana passada, falámos sobre X, e decidiste Y”.
- **Graph RAG** – utilização de grafos de conhecimento (Neo4j integrado ao Qdrant) para representar relações complexas entre documentos, entidades e conceitos.
- **Cache semântico inteligente** – embeddings frequentemente usados são armazenados em Redis para evitar recálculo e reduzir latência.
- **Indexação incremental de fontes externas** – suporte a feeds RSS, atualizações de bases de dados e repositórios de documentos, com re‑indexação automática.

### 3. 🔧 Ferramentas e Ações Mais Poderosas

- **Browser Agent** – agente com navegador headless (Playwright/Selenium) para extrair dados dinâmicos, interagir com formulários, autenticar‑se em portais e gerar relatórios baseados em informação da web.
- **Workflow Engine integrada** – capacidade de orquestrar pipelines complexos (ex.: “Extrai dados do S3, limpa, treina modelo e envia relatório por e‑mail”) utilizando motores como Temporal ou Prefect, expostos como ferramentas.
- **Registo automático de APIs externas** – dado um ficheiro OpenAPI/Swagger, o sistema gera automaticamente a tool definition e a disponibiliza no Tool Registry, sem código adicional.
- **Multi‑ferramenta condicional** – agentes podem combinar várias ferramentas numa única chamada de alto nível, com gestão automática de dependências.

### 4. 🎤🧠📷 Multi‑Modalidade e Interfaces Avançadas

- **Entrada/saída de voz** – integração total com Whisper (STT) e modelos TTS de última geração (ex.: ElevenLabs, Piper) com suporte a pistas emocionais.
- **Visão computacional** – processamento de imagens, gráficos e documentos digitalizados para extração de dados estruturados, reconhecimento de objetos e OCR avançado.
- **Interpretação de áudio** – análise de tom de voz, deteção de emoções e níveis de urgência para adaptar a resposta do agente.
- **Saída em realidade aumentada** – protótipo de agente que gera objetos 3D ou sobreposições para headsets, útil em simulações industriais e formação.

### 5. 🐝 Colaboração Multi‑Agente e Enxames

- **Debate entre agentes** – dois ou mais agentes argumentam sobre um tópico, um moderador avalia a qualidade dos argumentos e sintetiza uma conclusão ponderada (melhora precisão analítica).
- **Enxames para processamento massivo** – dezenas de agentes leves trabalham em paralelo para analisar grandes volumes de documentos, imagens ou registos, coordenados por um swarm manager.
- **Agentes proativos** – agentes de background monitorizam fontes (ex.: calendários, mercados financeiros, redes sociais) e iniciam conversas quando detetam eventos relevantes.
- **Hierarquias dinâmicas** – criação de hierarquias de agentes em tempo de execução, onde um agente “líder” pode delegar sub‑tarefas e escalar problemas.

### 6. 🛡️ Segurança e Governança Avançadas

- **RBAC para ferramentas e agentes** – cada utilizador/tenant define quais ferramentas e tipos de agente estão disponíveis, com políticas granulares.
- **Execução em enclaves confidenciais** – suporte a ambientes de execução seguros (Intel SGX, AWS Nitro Enclaves) para processamento de dados altamente sensíveis sem exposição ao operador da plataforma.
- **Guardião ético e de conformidade** – um agente separado que audita todas as saídas contra políticas de uso aceitável, regulações (GDPR, HIPAA) e vieses.
- **Criptografia ponta‑a‑ponta para ferramentas** – possibilidade de encriptar parâmetros e resultados de ferramentas críticas, usando chaves geridas pelo cliente.

### 7. 📊 Observabilidade, Escalabilidade e DevOps

- **Agent Performance Score (APS)** – métrica composta baseada em precisão, latência, custo e feedback do utilizador, usada para seleção automática de agentes e modelos.
- **Multi‑cloud e híbrido** – templates Terraform para Azure, GCP e ambientes on‑premise; suporte a execução de agentes em edge com modelos locais (Ollama, vLLM) para baixa latência.
- **Marketplace API** – API pública para que terceiros submetam agentes e ferramentas, com testes automáticos de segurança e desempenho antes da publicação.
- **Dashboard avançado para developers** – visualização em tempo real da árvore de decisão dos agentes, execução de ferramentas, tokens gastos e tracing distribuído completo.

### 8. 💰 Monetização e Produto

- **Planos de subscrição modulares** – pacotes baseados no número de agentes especializados, volume de documentos indexados e chamadas a ferramentas premium.
- **Loja de agentes premium** – agentes certificados por parceiros (ex.: análise financeira, suporte jurídico) comercializados como add‑ons.
- **White‑label** – empresas podem incorporar a plataforma KoreAI nos seus produtos, personalizando agentes, memória e interfaces com a própria marca.
- **Insights e exportação de relatórios** – painel de produtividade que mostra tempo poupado, tarefas concluídas e ROI estimado, com possibilidade de exportar relatórios executivos gerados pelos agentes.

---

## 🇬🇧 English Version

### KoreAI Platform Expansion Roadmap

This document outlines the planned new features and improvements for the **KoreAI** platform, grouped into **8 strategic pillars**. Each pillar details additions that will raise the intelligence, security, scalability, and value of the multi‑agent system.

### 1. 🧠 Agent Intelligence Expansion

- **Adaptive cognitive profiles** – agents adjust communication style, technical depth, and emotional tone based on a user preference model stored in Redis.
- **Self‑improvement loop with human feedback (RLHF)** – automatic collection of quality signals (e.g., resolution time, explicit ratings) to fine‑tune models via LoRA/QLoRA in a closed environment.
- **Internal reviewer agent** – a behind‑the‑scenes agent that checks facts, coherence, and policy compliance before the final answer is delivered, reducing hallucinations.
- **Emotional memory** – history of the user’s emotional state for more empathetic interactions (e.g., detecting frustration and responding with greater patience).

### 2. 📚 RAG & Long‑term Memory Upgrades

- **Episodic vector memory** – store not only facts but also complete conversational “episodes”, enabling recall like: “Last week we talked about X, and you decided Y.”
- **Graph RAG** – integration with knowledge graphs (Neo4j + Qdrant) to represent complex relationships among documents, entities, and concepts.
- **Intelligent semantic cache** – frequently used embeddings are cached in Redis to avoid recomputation and reduce latency.
- **Incremental external source indexing** – support for RSS feeds, database change streams, and live document repositories with automatic re‑indexing.

### 3. 🔧 More Powerful Tools & Actions

- **Browser Agent** – headless browser agent (Playwright/Selenium) for dynamic data extraction, form interaction, portal authentication, and web‑based report generation.
- **Integrated workflow engine** – ability to orchestrate complex pipelines (e.g., “Pull CSV from S3, clean, train model, email report”) using engines like Temporal or Prefect, exposed as tools.
- **Auto‑registration of external APIs** – given an OpenAPI/Swagger spec, the system automatically generates the tool definition and registers it in the Tool Registry, no extra code required.
- **Conditional multi‑tool** – agents can combine several tools into a single high‑level call with automatic dependency management.

### 4. 🎤🧠📷 Multi‑Modality & Advanced Interfaces

- **Voice input/output** – full integration with Whisper (STT) and state‑of‑the‑art TTS (e.g., ElevenLabs, Piper) with support for emotional cues.
- **Computer vision** – image, chart, and scanned‑document processing for structured data extraction, object recognition, and advanced OCR.
- **Audio interpretation** – tone‑of‑voice analysis, emotion detection, and urgency levels to adapt the agent’s response.
- **Augmented reality output** – prototype agent that generates 3D objects or overlays for headsets, useful in industrial simulations and training.

### 5. 🐝 Multi‑Agent Collaboration & Swarms

- **Agent debate** – two or more agents argue a topic; a moderator evaluates the quality of arguments and synthesizes a well‑reasoned conclusion (improves analytical accuracy).
- **Massive processing swarms** – dozens of lightweight agents work in parallel to analyze large volumes of documents, images, or logs, coordinated by a swarm manager.
- **Proactive agents** – background agents monitor sources (calendars, financial markets, social media) and initiate conversations when relevant events are detected.
- **Dynamic hierarchies** – runtime creation of agent hierarchies, where a “leader” agent can delegate sub‑tasks and escalate issues.

### 6. 🛡️ Advanced Security & Governance

- **RBAC for tools and agents** – each user/tenant defines which tools and agent types are available, with granular access policies.
- **Confidential enclave execution** – support for secure execution environments (Intel SGX, AWS Nitro Enclaves) to process highly sensitive data without exposing it to the platform operator.
- **Ethics & compliance guardian** – a separate agent that audits every output against acceptable use policies, regulations (GDPR, HIPAA), and bias detection.
- **End‑to‑end tool encryption** – ability to encrypt parameters and results of critical tools using customer‑managed keys.

### 7. 📊 Observability, Scalability & DevOps

- **Agent Performance Score (APS)** – composite metric based on accuracy, latency, cost, and user feedback, used for automatic agent and model selection.
- **Multi‑cloud & hybrid** – Terraform templates for Azure, GCP, and on‑premise deployments; support for edge agents running local models (Ollama, vLLM) for low latency.
- **Marketplace API** – public API for third parties to submit agents and tools, with automated security and performance testing before publication.
- **Advanced developer dashboard** – real‑time visualization of agent decision trees, tool execution, token consumption, and full distributed tracing.

### 8. 💰 Monetization & Product

- **Modular subscription plans** – packages based on the number of specialized agents, volume of indexed documents, and premium tool calls.
- **Premium agent store** – partner‑certified agents (e.g., financial analysis, legal support) sold as add‑ons.
- **White‑label** – enterprises can embed the KoreAI platform into their own products, customizing agents, memory, and interfaces with their brand.
- **Insights & report export** – productivity dashboard showing time saved, tasks completed, and estimated ROI, with the ability to export executive reports generated by the agents.


---
#  KoreAI – Roadmap & Arquitetura Futura / Future Architecture

**🇧🇷 Português** | **[🇬🇧 English](#-english-version)**

---

## 🗺️ Diagrama da Arquitetura Expandida (Visão Futura)

```mermaid
flowchart TB
    subgraph Frontend["Frontend (Next.js)"]
        Dashboard["Dashboard"]
        Billing["Billing / Store"]
    end

    subgraph Proxy["Proxy (Go)"]
        Gateway["API Gateway"]
        Cache["Cache (Redis)"]
        CB["Circuit Breaker"]
        Marketplace["Marketplace API"]
    end

    subgraph Core["Core (Python) – novos módulos"]
        Engine["Engine (engine.py)"]
        TaskPlanner["Task Planner"]
        
        subgraph AgentPool["Agent Pool"]
            Researcher["Researcher"]
            Coder["Coder"]
            Analyst["Analyst"]
            Creative["Creative"]
            Browser["Browser Agent"]
            Voice["Voice Agent"]
            Proactive["Proactive Agent"]
            Reviewer["Reviewer"]
            Ethics["Ethics Guardian"]
            Moderator["Moderator"]
        end
        
        subgraph Memory["Memory"]
            Episodic["Episodic Store"]
            Graph["Graph Store (Neo4j)"]
            Emotional["Emotional Store"]
            SemanticCache["Semantic Cache"]
        end
        
        subgraph Tools["Tools"]
            BrowserTool["Browser Tool"]
            Vision["Vision Tool"]
            AR["AR Tool"]
            APIReg["API Registrar"]
        end
        
        subgraph Workflows["Workflows"]
            WEngine["Temporal/Prefect Engine"]
        end
        
        subgraph Swarm["Swarm"]
            SwManager["Swarm Manager"]
            Strategies["Debate/Hierarchy"]
        end
    end

    subgraph Security["Security (Go + Python)"]
        Enclave["Enclave (SGX)"]
        E2EEnc["E2E Encryption"]
        RBAC["RBAC"]
        GuardMiddleware["Guardian Middleware"]
    end

    subgraph Infrastructure["Infrastructure"]
        Terraform["Terraform (multi-cloud)"]
        Edge["Edge (Ollama)"]
    end

    Frontend --> Proxy
    Proxy --> Engine
    Engine --> TaskPlanner
    TaskPlanner --> AgentPool
    AgentPool --> Memory
    AgentPool --> Tools
    AgentPool --> Swarm
    AgentPool --> Workflows
    AgentPool --> Security
    Infrastructure --> Proxy
    Infrastructure --> Core
