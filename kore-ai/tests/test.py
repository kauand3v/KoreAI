# tests/test.py
import pytest
from fastapi.testclient import TestClient
from proxy.main import app
from engine.parser.lexer import KoreLexer

client = TestClient(app)

def test_lexer_tokenization():
    lexer = KoreLexer()
    tokens = lexer.tokenize('route "test" { priority: 1 }')
    assert len(tokens) > 0
    assert tokens[0].type == 'ROUTE'

def test_proxy_unauthorized():
    # Testa se o Proxy bloqueia sem header de tenant
    response = client.get("/v1/chat/completions")
    assert response.status_code == 422 # Erro de validação de header

def test_proxy_authorized():
    # Testa com header correto
    response = client.get("/v1/chat/completions", headers={"x-tenant-id": "tenant_alpha"})
    assert response.status_code == 200