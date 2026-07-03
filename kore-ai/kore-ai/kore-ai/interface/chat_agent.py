# interface/chat_agent.py
import os
from pathlib import Path
from typing import List, Dict, Any

# Simulação de LLM (pode ser substituído por Ollama, Grok, LangChain, etc.)
class SimpleLLM:
    def generate(self, prompt: str) -> str:
        # Em produção: substitua por chamada real ao Ollama / OpenAI / Grok
        # Exemplo: client.chat.completions.create(...)
        return f"[Resposta simulada do modelo]\nBaseado na análise do sistema:\n{prompt[:300]}..."


class KoreChatAgent:
    def __init__(self):
        self.llm = SimpleLLM()
        self.engine = None  # Lazy load

    def _get_engine(self):
        if not self.engine:
            from engine.kore_engine import KoreEngine
            self.engine = KoreEngine()
        return self.engine

    def _read_policies(self) -> List[Dict[str, Any]]:
        """Lê todas as políticas .kore do diretório"""
        policies = []
        for file in Path(".").rglob("*.kore"):
            try:
                content = file.read_text()
                policies.append({
                    "file": str(file),
                    "content": content[:2000]  # limitar tamanho
                })
            except Exception:
                continue
        return policies

    def explain(self, query: str) -> str:
        """Responde perguntas sobre o sistema em linguagem natural"""
        
        context = {
            "policies": self._read_policies(),
            "system_status": self._get_engine().get_system_status(),
            "recent_errors": self._get_engine().get_recent_errors()
        }

        prompt = f"""
Você é o KoreAI Assistant, um SRE virtual especializado em infraestrutura declarativa.

Contexto atual do sistema:
{context}

Pergunta do usuário: {query}

Responda de forma clara, técnica quando necessário, mas acessível. 
Se identificar um problema, sugira soluções concretas.
Use emojis quando apropriado.
"""

        response = self.llm.generate(prompt)
        
        # Adiciona sugestão de ação
        if any(word in query.lower() for word in ["falha", "falhando", "error", "problema"]):
            response += "\n\n💡 Sugestão: Execute `kore status` ou `kore validate` para mais detalhes."
        
        return response


# Exemplo de uso direto (útil para debugging)
if __name__ == "__main__":
    agent = KoreChatAgent()
    print(agent.explain("Por que a rota de pagamentos está falhando?"))