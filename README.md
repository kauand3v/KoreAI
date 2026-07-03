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

