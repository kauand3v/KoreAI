# 🇧🇷 VERSÃO EM PORTUGUÊS

# 🧠 KoreAI: O Core de Orquestração para LLMs

O **KoreAI** é um *AI Gateway* e orquestrador de alta performance projetado para resolver as maiores dores na adoção de Inteligência Artificial corporativa: **altos custos de tokens, instabilidade de APIs e vendor lock-in**.

Ao introduzir uma DSL própria (**KoreScript**), o KoreAI atua como um middleware inteligente que roteia requisições dinamicamente, garantindo resiliência, economia e baixa latência sem a necessidade de reescrever o código da aplicação cliente.

---

## 🎯 O Problema vs. A Solução KoreAI

| A Dor do Mercado (O Problema) | Como o KoreAI Resolve |
| --- | --- |
| **Vendor Lock-in** | Compatibilidade nativa com a API da OpenAI. Troque de provedor por trás dos panos (Anthropic, Gemini, Ollama) sem mudar o código cliente. |
| **Custos Exorbitantes** | **Roteamento Inteligente:** Tarefas complexas vão para o GPT-4o; tarefas simples são roteadas para modelos locais ou mais baratos (Llama-3). |
| **Instabilidade / Downtime** | **Circuit Breaker:** Se uma API falhar, o KoreAI intercepta o erro e faz um fallback automático para o próximo modelo disponível na política. |
| **Requisições Redundantes** | **Semantic Caching:** Respostas são cacheadas com base no significado da pergunta, evitando pagar duas vezes pela mesma resposta. |

---

## 🇧🇷 Arquitetura Poliglota: O Propósito de Cada Linguagem

O **KoreAI** utiliza uma arquitetura baseada em microsserviços onde cada linguagem foi escolhida de forma pragmática para resolver gargalos específicos do ecossistema de Inteligência Artificial. Não acreditamos em uma "linguagem de prata", mas sim em usar a ferramenta certa para o trabalho certo.

### 🔵 Go (Golang): A Espinha Dorsal de Rede e Concorrência

* **Por que usamos:** O gargalo de um *AI Gateway* não é a IA em si, mas a rede (I/O). Go possui um modelo de concorrência nativo (*Goroutines*) extremamente leve e um coletor de lixo otimizado para baixa latência.
* **Como é implementado:** O Go atua como o servidor de borda (Reverse Proxy). Ele gerencia as conexões HTTP de entrada e se comunica com as outras camadas via chamadas gRPC de alta velocidade.
* **Para que serve:** * Receber milhares de requisições simultâneas da aplicação cliente.
* Gerenciar o estado do **Circuit Breaker** e o *Rate Limiting*.
* Roteamento de rede de forma assíncrona, garantindo que o gateway nunca trave a aplicação cliente enquanto aguarda a resposta da OpenAI.


### 🐍 Python: O Plano de Controle, DSL e Ecossistema

* **Por que usamos:** Python é o padrão ouro (*lingua franca*) da Inteligência Artificial. Ele possui o ecossistema mais rico para integração, tooling e construção de analisadores sintáticos (parsers).
* **Como é implementado:** Atua na camada de "Control Plane" (Plano de Controle) e interface de usuário. Não lida com o tráfego pesado de rede, mas orquestra as regras do sistema.
* **Para que serve:**
* **KoreScript Parser:** Ler, tokenizar e compilar a nossa DSL (arquivos `.kore`), transformando texto em regras lógicas para o Go executar.
* **CLI e Ferramentas:** O utilitário `kore-cli` para desenvolvedores validarem políticas antes do deploy.
* **Agentes de Interface:** Executar a integração com frameworks como LangChain/LangGraph para criar o chatbot interno que auxilia no debug do roteamento.



---



## 🇺🇸 Polyglot Architecture: The Purpose of Each Language

**KoreAI** utilizes a microservices-based architecture where each language was pragmatically chosen to solve specific bottlenecks in the Artificial Intelligence ecosystem. We do not believe in a "silver bullet" language, but rather in using the right tool for the right job.

### 🔵 Go (Golang): The Network & Concurrency Backbone

* **Why we use it:** The bottleneck of an AI Gateway is not the AI itself, but the network (I/O). Go features a remarkably lightweight native concurrency model (*Goroutines*) and a garbage collector optimized for ultra-low latency.
* **How it is implemented:** Go acts as the edge server (Reverse Proxy). It handles incoming HTTP connections and communicates with the other architectural layers via high-speed gRPC calls.
* **What it is used for:**
* Handling thousands of simultaneous requests from client applications.
* Managing **Circuit Breaker** state and Rate Limiting operations.
* Asynchronous network routing, ensuring the gateway never blocks the client application while waiting for OpenAI's response.



### 🐍 Python: The Control Plane, DSL & Ecosystem Glue

* **Why we use it:** Python is the gold standard (*lingua franca*) of Artificial Intelligence. It possesses the richest ecosystem for integrations, tooling, and building syntactic analyzers (parsers).
* **How it is implemented:** It operates in the "Control Plane" and user interface layer. It doesn't handle the heavy network traffic but rather orchestrates the system's underlying rules.
* **What it is used for:**
* **KoreScript Parser:** Reading, tokenizing, and compiling our custom DSL (`.kore` files), transforming plain text into logical execution rules for Go.
* **CLI and Tooling:** Powering the `kore-cli` utility for developers to validate routing policies prior to deployment.
* **Agent Interfaces:** Running integrations with frameworks like LangChain/LangGraph to build the internal chatbot that assists in routing debugging.




## 🏗️ KoreScript: Infraestrutura como Código para IA

O grande diferencial do KoreAI é a separação entre lógica de negócios e política de IA. Em vez de espalhar `if/else` pelo código, você define regras de roteamento usando o **KoreScript**, uma linguagem declarativa projetada para este ecossistema.

### `policies/production.kore`

```hcl
route "finance-analysis" {
  priority: 1
  model: "gpt-4o"
  fallback: "claude-3-5-sonnet"
  cache: true
}

route "customer-support" {
  priority: 2
  model: "llama-3-local"
  circuit_breaker {
    threshold: 0.8
    duration: "60s"
  }
}

```

---

## 🚀 Quickstart

### 1. Inicie a KoreVM e o Dashboard

O projeto já inclui um `docker-compose.yml` pré-configurado com Redis (para o cache semântico) e Grafana (para métricas de economia de custos).

```bash
docker-compose up -d

```

### 2. Valide suas Políticas (Kore CLI)

O motor de compilação valida as regras antes do deploy, garantindo segurança em produção.

```bash
kore-cli validate policies/production.kore

```

### 3. Integração Transparente (Drop-in Replacement)

Apenas mude a `BASE_URL` da sua aplicação. O KoreAI faz o resto.

```python
from openai import OpenAI

client = OpenAI(
    api_key="sua-chave",
    base_url="http://localhost:8080/v1" # Apontando para o gateway KoreAI
)

response = client.chat.completions.create(
    model="customer-support", # Nome da rota definida no KoreScript
    messages=[{"role": "user", "content": "Qual o status do meu pedido?"}]
)

```

---

## 👤 Sobre o Autor & Contato

Este projeto foi construído para demonstrar habilidades em **Engenharia de Software de Backend, Arquitetura de Sistemas Distribuídos e LLMOps**.

Estou aberto a novas oportunidades e desafios em Engenharia de Software (Sênior/Staff) e posições focadas em arquitetura de IA.

* **LinkedIn:** (https://www.linkedin.com/in/kauan-oliveira-324264378/)
* **Email:** kauandias747474@gmail.com

---

---

# 🇺🇸 ENGLISH VERSION

# 🧠 KoreAI: The Core LLM Orchestrator

**KoreAI** is a high-performance AI Gateway and orchestrator designed to solve the biggest pain points in enterprise AI adoption: **high token costs, API instability, and vendor lock-in**.

By introducing a proprietary Domain-Specific Language (**KoreScript**), KoreAI acts as an intelligent middleware that dynamically routes requests, ensuring resilience, cost-efficiency, and low latency without requiring application-level code changes.

---

## 🎯 The Problem vs. The KoreAI Solution

| Industry Pain Point (The Problem) | How KoreAI Solves It |
| --- | --- |
| **Vendor Lock-in** | Native compatibility with the OpenAI API structure. Switch providers (Anthropic, Gemini, Ollama) under the hood with zero client-side code changes. |
| **Skyrocketing Costs** | **Intelligent Routing:** Complex tasks are sent to GPT-4o; simple tasks are dynamically routed to local or cheaper models (Llama-3). |
| **API Instability / Downtime** | **Circuit Breaker:** If an API fails, KoreAI intercepts the error and triggers an automatic fallback to the next available model. |
| **Redundant Requests** | **Semantic Caching:** Answers are cached based on the semantic meaning of the prompt, saving tokens on repetitive queries. |

---

## 🏗️ KoreScript: Infrastructure as Code for AI

KoreAI's killer feature is the separation of business logic and AI routing policies. Instead of hardcoding complex fallback `if/else` statements, developers define rules using **KoreScript**, a declarative language built for AI routing.

### `policies/production.kore`

```hcl
route "finance-analysis" {
  priority: 1
  model: "gpt-4o"
  fallback: "claude-3-5-sonnet"
  cache: true
}

route "customer-support" {
  priority: 2
  model: "llama-3-local"
  circuit_breaker {
    threshold: 0.8
    duration: "60s"
  }
}

```

---

## 🚀 Quickstart

### 1. Boot up the KoreVM and Metrics Dashboard

The project includes a ready-to-use `docker-compose.yml` with Redis (for semantic caching) and Grafana (for cost-saving metrics).

```bash
docker-compose up -d

```

### 2. Validate your Policies (Kore CLI)

The built-in parser validates your rules before deployment to ensure production safety.

```bash
kore-cli validate policies/production.kore

```

### 3. Transparent Integration (Drop-in Replacement)

Simply change your application's `BASE_URL`. KoreAI handles the rest.

```python
from openai import OpenAI

client = OpenAI(
    api_key="your-key",
    base_url="http://localhost:8080/v1" # Pointing to the KoreAI Gateway
)

response = client.chat.completions.create(
    model="customer-support", # The route name defined in KoreScript
    messages=[{"role": "user", "content": "What is the status of my order?"}]
)

```

---
# 🎛️ KoreAI Control Plane (Frontend)

[![Stack](https://img.shields.io/badge/Stack-Next.js_%7C_TypeScript_%7C_Tailwind-black)](#)
[![Design](https://img.shields.io/badge/UI-Shadcn%20%7C%20Zustand%20%7C%20Recharts-zinc)](#)

> O **KoreAI Control Plane** é a interface visual de alta densidade (Dashboard SaaS) desenvolvida para gerenciar e monitorar o gateway poliglota de inteligência artificial. Inspirado no design de ferramentas líderes como Vercel e Linear.

---

## 🇧🇷 Conteúdo em Português

### 📋 Visão Geral
A interface atua como o painel de controle do sistema, abstraindo a complexidade do roteamento de IA. Ela foca em **Observabilidade** e **Mutação Segura de Estado**, permitindo que os desenvolvedores monitorem a latência da rede, ajustem regras de *Circuit Breaker* e adicionem novas políticas de roteamento em tempo real.

### 🛠️ Stack Tecnológica
* **Next.js (App Router) & TypeScript:** Renderização de ponta e tipagem rigorosa para escalabilidade.
* **TailwindCSS & Shadcn/ui:** Estética minimalista focada no *Dark Mode* (`zinc-950` e bordas finas), reduzindo a fadiga cognitiva.
* **Zustand:** Gerenciamento de estado global leve para lidar com streams de telemetria sem engasgar a renderização.
* **React Hook Form + Zod:** Validação implacável no cliente, garantindo que regras de infraestrutura malformadas nunca cheguem ao backend.
* **Recharts:** Visualização tática e interativa de consumo de tokens e latência.

### 📂 Estrutura de Pastas (Frontend)
```text
frontend/
├── src/
│   ├── app/                  # Next.js App Router (Layouts e Páginas)
│   │   └── dashboard/        # Painel central de telemetria e gestão
│   ├── components/           # Componentes atômicos (Shadcn) e formulários
│   ├── store/                # Estado global de telemetria com Zustand
│   └── schemas/              # Validações Zod para mutação de infraestrutura

```

### 🚀 Como Executar Localmente

1. Certifique-se de ter o Node.js instalado (v18+ recomendado).
2. Entre na pasta do frontend e instale as dependências:

```bash
   cd frontend
   npm install

```

3. Inicie o servidor de desenvolvimento:

```bash
   npm run dev

```

4. Acesse o painel pelo navegador em: `http://localhost:3000`

---

## 🇺🇸 Content in English

### 📋 Overview

The interface serves as the command center for the system, abstracting the complexity of AI routing. It focuses deeply on **Observability** and **Secure State Mutation**, enabling developers to monitor network latency, tune *Circuit Breaker* thresholds, and provision new routing policies in real-time.

### 🛠️ Tech Stack

* **Next.js (App Router) & TypeScript:** Cutting-edge rendering paradigms and strict end-to-end type safety.
* **TailwindCSS & Shadcn/ui:** Minimalist, *Dark Mode* focused aesthetic (`zinc-950` and 1px borders) crafted to prevent cognitive overload.
* **Zustand:** Atomic, low-latency global state management that processes incoming telemetry without cascading DOM re-renders.
* **React Hook Form + Zod:** Rigorous client-side payload validation, ensuring malformed infrastructure mutations never breach the backend.
* **Recharts:** Tactical and interactive data visualization for token burn rates and proxy latency.

### 📂 Directory Structure (Frontend)

```text
frontend/
├── src/
│   ├── app/                  # Next.js App Router (Layouts and Pages)
│   │   └── dashboard/        # Centralized telemetry and management panel
│   ├── components/           # Atomic UI components (Shadcn) and forms
│   ├── store/                # Zustand global state (Live telemetry)
│   └── schemas/              # Zod schemas for infrastructure mutations

```

### 🚀 How to Run Locally

1. Ensure Node.js is installed (v18+ recommended).
2. Navigate to the frontend directory and install dependencies:

```bash
   cd frontend
   npm install

```

3. Spin up the development server:

```bash
   npm run dev

```

4. Access the dashboard via your browser at: `http://localhost:3000`



## 👤 About the Author & Contact

This project was built to demonstrate expertise in **Backend Software Engineering, Distributed Systems Architecture, and LLMOps**.

I am currently open to new opportunities and challenges in Software Engineering and AI Architecture roles. Let's connect!

* **LinkedIn:** https://www.linkedin.com/in/kauan-oliveira-324264378/
* **Email:** kauandias747474@gmail.com
