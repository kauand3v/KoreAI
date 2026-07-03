# proxy/middleware.py
from fastapi import Request
from starlette.middleware.base import BaseHTTPMiddleware
import time

class LoggingMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        start_time = time.time()
        # Lógica de intercepção (antes da requisição)
        response = await call_next(request)
        # Lógica pós-requisição
        process_time = time.time() - start_time
        print(f"Request: {request.url.path} | Duration: {process_time:.4f}s")
        return response