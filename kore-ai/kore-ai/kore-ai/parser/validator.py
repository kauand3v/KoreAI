from typing import List
from .parser import RouteConfig

class KoreValidator:
    ALLOWED_MODELS = ["hermes3:8b", "hermes3:70b", "llama3"]

    def validate(self, routes: List[RouteConfig]):
        for route in routes:
            # 1. Validação de Modelo
            if route.model not in self.ALLOWED_MODELS:
                raise ValueError(f"Modelo inválido na rota '{route.name}': {route.model}")
            
            # 2. Validação de Threshold (Circuit Breaker)
            if route.circuit_breaker:
                if not (0.0 <= route.circuit_breaker.threshold <= 1.0):
                    raise ValueError(f"Threshold inválido na rota '{route.name}'. Deve ser entre 0 e 1.")
            
            # 3. Validação de Prioridade
            if route.priority < 0:
                raise ValueError(f"Prioridade negativa na rota '{route.name}'")
        
        return True