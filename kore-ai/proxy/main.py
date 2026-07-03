# proxy/main.py
from fastapi import FastAPI, Depends
from .auth import AuthManager
from .middleware import LoggingMiddleware

app = FastAPI(title="KoreAI Proxy Control Plane")
app.add_middleware(LoggingMiddleware)

@app.get("/health")
async def health_check():
    return {"status": "ok"}

@app.get("/v1/chat/completions", dependencies=[Depends(AuthManager.verify_tenant)])
async def chat_proxy():
    # Aqui a lógica do seu Engine seria chamada
    return {"message": "Requisição validada e encaminhada pelo Proxy"}