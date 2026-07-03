"""
tenant_registry.py
Plano de controle: registro em memória dos recursos permitidos por tenant.
Simula um banco de dados / Redis até a persistência real ser plugada.
"""

from dataclasses import dataclass, field
from threading import RLock
from typing import Dict, List


class TenantNotFoundError(Exception):
    """Levantada quando tenant_id não possui configuração registrada."""

    def __init__(self, tenant_id: str) -> None:
        self.tenant_id = tenant_id
        super().__init__(f"Tenant '{tenant_id}' não encontrado no registry.")


@dataclass(frozen=True)
class TenantConfig:
    tenant_id: str
    allowed_models: List[str] = field(default_factory=list)
    rate_limit: int = 1000  # requisições por minuto


class TenantRegistry:
    """
    Mapeia tenant_id -> TenantConfig.
    Thread-safe via RLock (workers concorrentes do control plane).
    """

    def __init__(self) -> None:
        self._lock = RLock()
        self._tenants: Dict[str, TenantConfig] = {}
        self._seed_defaults()

    def _seed_defaults(self) -> None:
        defaults = [
            TenantConfig(
                tenant_id="tenant-acme",
                allowed_models=["hermes3:8b", "llama3"],
                rate_limit=1000,
            ),
            TenantConfig(
                tenant_id="tenant-globex",
                allowed_models=["llama3"],
                rate_limit=500,
            ),
        ]
        for cfg in defaults:
            self._tenants[cfg.tenant_id] = cfg

    def register_tenant(self, config: TenantConfig) -> None:
        """Cria ou substitui a configuração de um tenant."""
        with self._lock:
            self._tenants[config.tenant_id] = config

    def get_tenant_config(self, tenant_id: str) -> TenantConfig:
        """
        Retorna o TenantConfig de tenant_id.

        Raises:
            TenantNotFoundError: se tenant_id não estiver registrado.
        """
        with self._lock:
            config = self._tenants.get(tenant_id)
        if config is None:
            raise TenantNotFoundError(tenant_id)
        return config

    def list_tenants(self) -> List[str]:
        with self._lock:
            return list(self._tenants.keys())