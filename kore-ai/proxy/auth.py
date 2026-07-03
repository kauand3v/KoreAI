# proxy/auth.py
from fastapi import HTTPException, Header

class AuthManager:
    # Simulação de um banco de dados de tenants
    VALID_TENANTS = {"tenant_alpha": "secret_key_1", "tenant_beta": "secret_key_2"}

    @staticmethod
    async def verify_tenant(x_tenant_id: str = Header(...)):
        if x_tenant_id not in AuthManager.VALID_TENANTS:
            raise HTTPException(status_code=403, detail="Tenant inválido ou não autorizado")
        return x_tenant_id