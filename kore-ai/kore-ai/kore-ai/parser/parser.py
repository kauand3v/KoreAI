from dataclasses import dataclass, field
from typing import List, Optional, Any
from .lexer import Token

@dataclass
class CircuitBreakerConfig:
    threshold: float = 0.0
    duration: int = 0

@dataclass
class RouteConfig:
    name: str
    priority: int = 0
    model: str = ""
    cache: bool = False
    circuit_breaker: Optional[CircuitBreakerConfig] = None

class KoreParser:
    def __init__(self, tokens: List[Token]):
        self.tokens = tokens
        self.pos = 0

    def consume(self, expected_type: str):
        if self.pos < len(self.tokens) and self.tokens[self.pos].type == expected_type:
            val = self.tokens[self.pos].value
            self.pos += 1
            return val
        raise SyntaxError(f"Esperava {expected_type}, encontrou {self.tokens[self.pos].type if self.pos < len(self.tokens) else 'EOF'}")

    def parse(self) -> List[RouteConfig]:
        routes = []
        while self.pos < len(self.tokens):
            if self.tokens[self.pos].value == 'route':
                routes.append(self.parse_route())
        return routes

    def parse_route(self) -> RouteConfig:
        self.consume('ROUTE')
        name = self.consume('STRING').strip('"')
        self.consume('LBRACE')
        
        config = RouteConfig(name=name)
        
        while self.tokens[self.pos].type != 'RBRACE':
            key = self.consume('ID')
            self.consume('COLON')
            
            if key == 'priority': config.priority = int(self.consume('NUMBER'))
            elif key == 'model': config.model = self.consume('STRING').strip('"')
            elif key == 'cache': config.cache = (self.consume('BOOL') == 'true')
            elif key == 'circuit_breaker':
                self.consume('LBRACE')
                cb = CircuitBreakerConfig()
                while self.tokens[self.pos].type != 'RBRACE':
                    cb_key = self.consume('ID')
                    self.consume('COLON')
                    if cb_key == 'threshold': cb.threshold = float(self.consume('NUMBER'))
                    elif cb_key == 'duration': cb.duration = int(self.consume('NUMBER'))
                self.consume('RBRACE')
                config.circuit_breaker = cb
        
        self.consume('RBRACE')
        return config