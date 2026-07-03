# interface/cli.py
import argparse
import sys
from pathlib import Path
from typing import Optional

from engine.kore_engine import KoreEngine
from interface.chat_agent import KoreChatAgent


class KoreCLI:
    def __init__(self):
        self.engine = KoreEngine()
        self.chat_agent = KoreChatAgent()

    def validate(self, file_path: str) -> None:
        """Valida um arquivo .kore"""
        path = Path(file_path)
        if not path.exists():
            print(f"❌ Arquivo não encontrado: {file_path}")
            sys.exit(1)

        try:
            result = self.engine.validate_policy(path.read_text())
            if result.is_valid:
                print(f"✅ Política válida: {path.name}")
                print(f"   Rotas detectadas: {len(result.routes)}")
            else:
                print(f"❌ Política inválida: {path.name}")
                for error in result.errors:
                    print(f"   • {error}")
        except Exception as e:
            print(f"❌ Erro durante validação: {e}")

    def deploy(self, file_path: str, tenant_id: str = "tenant_alpha") -> None:
        """Faz deploy de uma política"""
        print(f"🚀 Iniciando deploy para tenant: {tenant_id}")
        result = self.engine.deploy_policy(file_path, tenant_id)
        
        if result.success:
            print("✅ Deploy realizado com sucesso!")
            print(f"   Versão: {result.version}")
            print(f"   Rotas ativadas: {len(result.deployed_routes)}")
        else:
            print("❌ Deploy falhou:")
            print(result.message)

    def status(self) -> None:
        """Mostra status do sistema"""
        status = self.engine.get_system_status()
        print("📊 Status do KoreAI")
        print(f"   Engine: {'🟢 Online' if status['healthy'] else '🔴 Degradado'}")
        print(f"   Rotas ativas: {status['active_routes']}")
        print(f"   Último deploy: {status['last_deploy']}")

    def explain(self, query: str) -> None:
        """Explica algo usando o agente cognitivo"""
        print("🤖 Kore Assistant pensando...\n")
        response = self.chat_agent.explain(query)
        print(response)


def main():
    parser = argparse.ArgumentParser(
        prog="kore",
        description="KoreAI CLI - Gerenciamento de Políticas de Roteamento Inteligente"
    )
    subparsers = parser.add_subparsers(dest="command", help="Comandos disponíveis")

    # kore validate
    validate_parser = subparsers.add_parser("validate", help="Valida um arquivo .kore")
    validate_parser.add_argument("file", help="Caminho do arquivo .kore")

    # kore deploy
    deploy_parser = subparsers.add_parser("deploy", help="Faz deploy de uma política")
    deploy_parser.add_argument("file", help="Caminho do arquivo .kore")
    deploy_parser.add_argument("--tenant", default="tenant_alpha", help="ID do tenant")

    # kore status
    subparsers.add_parser("status", help="Mostra status do sistema")

    # kore explain
    explain_parser = subparsers.add_parser("explain", help="Pergunta ao assistente cognitivo")
    explain_parser.add_argument("query", nargs="+", help="Pergunta sobre o sistema")

    args = parser.parse_args()

    cli = KoreCLI()

    if args.command == "validate":
        cli.validate(args.file)
    elif args.command == "deploy":
        cli.deploy(args.file, args.tenant)
    elif args.command == "status":
        cli.status()
    elif args.command == "explain":
        query = " ".join(args.query)
        cli.explain(query)
    else:
        parser.print_help()


if __name__ == "__main__":
    main()