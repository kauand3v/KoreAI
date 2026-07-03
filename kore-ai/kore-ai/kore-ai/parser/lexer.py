import re
from typing import NamedTuple, List

 class Token(NamedTuple):
     type: str
     value: str
     
    class KoreLexer:
        TOKEN_SPEC = [
           ('ROUTE',     r'route'),
        ('LBRACE',    r'\{'),
        ('RBRACE',    r'\}'),
        ('COLON',     r':'),
        ('BOOL',      r'true|false'),
        ('NUMBER',    r'\d+(\.\d+)?'),
        ('STRING',    r'"[^"]*"'),
        ('ID',        r'[a-zA-Z_][a-zA-Z0-9_]*'),
        ('SKIP',      r'[ \t\n\r]+'),
        ('MISMATCH',  r'.'),
    ]

    def __init__(self):
        self.regex = re.compile('|'.join(f'(?P<{name}>{pattern})' for name, pattern in self.TOKEN_SPEC))

    def tokenize(self, code: str) -> List[Token]:
        tokens = []
        for mo in self.regex.finditer(code):
            kind = mo.lastgroup
            value = mo.group(kind)
            if kind == 'SKIP':
                continue
            elif kind == 'MISMATCH':
                raise SyntaxError(f"Caractere inesperado: {value}")
            tokens.append(Token(kind, value))
        return tokens 