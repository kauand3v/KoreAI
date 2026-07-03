import os
import logging
from typing import List, Optional

# Importações dos teus módulos internos
from kore_ai.parser.lexer import KoreLexer
from kore_ai.parser.parser import KoreParser
from kore_ai.parser.validator import KoreValidator
from kore_ai.openclaw.openclaw_bridge import OpenClawConnector

class KoreEngine:
    """
    O Orquestrador Central (Engine) do KoreAI.
    Este componente coordena o pipeline completo:
    Leitura -> Lexing -> Parsing -> Validação -> Integração.
    """
    
    def __init__(self, bridge_endpoint: str = "https://api.openclaw.com", api_key: str = "PROD_SECRET"):
        self.lexer = KoreLexer()
        self.validator = KoreValidator()
        self.bridge = OpenClawConnector(bridge_endpoint, api_key)
        self.logger = logging.getLogger("KoreEngine")
        logging.basicConfig(level=logging.INFO)

    def process_policy_file(self, file_path: str) -> Optional[List]:
        """
        Pipeline principal: processa um ficheiro .kore e retorna a AST validada.
        """
        try:
            self.logger.info(f"Iniciando processamento do ficheiro: {file_path}")
            
            # 1. Leitura
            if not os.path.exists(file_path):
                raise FileNotFoundError(f"Ficheiro não encontrado: {file_path}")
            
            with open(file_path, 'r', encoding='utf-8') as f:
                code = f.read()

            # 2. Lexing
            tokens = self.lexer.tokenize(code)
            
            # 3. Parsing
            parser = KoreParser(tokens)
            ast = parser.parse()
            
            # 4. Validação
            self.validator.validate(ast)
            
            # 5. Notificação de Sucesso via Bridge
            self.bridge.emit_log(
                "POLICY_DEPLOY_SUCCESS", 
                f"Configuração {file_path} validada com sucesso.", 
                severity="INFO"
            )
            
            return ast

        except Exception as e:
            error_msg = f"Falha no processamento: {str(e)}"
            self.logger.error(error_msg)
            
            # Reportar erro para a Bridge
            self.bridge.emit_log("POLICY_DEPLOY_ERROR", error_msg, severity="ERROR")
            raise e

# --- Exemplo de Execução (Entrypoint para testes) ---
if __name__ == "__main__":
    # Inicializa o motor
    engine = KoreEngine()
    
    # Caminho do teu ficheiro de regras
    caminho_politica = "kore-ai/policies/production.kore"
    
    try:
        resultado = engine.process_policy_file(caminho_politica)
        print("\n✅ Política carregada e validada com sucesso!")
        for rota in resultado:
            print(f"-> Rota: {rota.name} | Modelo: {rota.model} | Cache: {rota.cache}")
            
    except Exception as e:
        print(f"\n❌ Erro crítico: {e}")