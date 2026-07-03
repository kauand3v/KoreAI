"""
policy_validator.py
Motor de políticas do Guardian: valida se um tenant pode rotear uma
requisição para um modelo específico (definido em um arquivo .kore),
prevenindo acesso a recursos fora do escopo de isolamento do tenant.
"""

from typing import Tuple

from tenant_registry import TenantNotFoundError, TenantRegistry


class GuardianPolicyValidator:
    """Consome o TenantRegistry para decidir se uma rota é segura."""

    def __init__(self, registry: TenantRegistry) -> None:
        self._registry = registry

    def validate_route_safety(
        self, tenant_id: str, requested_model: str
    ) -> Tuple[bool, str]:
        """
        Verifica se requested_model está dentro do escopo permitido
        de tenant_id.

        Returns:
            (True, "")                  se o acesso é permitido.
            (False, "<mensagem de erro>") se não for.
        """
        if not tenant_id:
            return False, "tenant_id ausente: requisição rejeitada pelo Guardian."

        if not requested_model:
            return False, "requested_model ausente: nenhum modelo especificado no .kore."

        try:
            config = self._registry.get_tenant_config(tenant_id)
        except TenantNotFoundError as exc:
            return False, f"Violação de isolamento: {exc}"

        if requested_model not in config.allowed_models:
            return False, (
                f"Violação de cross-tenant: o tenant '{tenant_id}' não tem "
                f"permissão para acessar o modelo '{requested_model}'. "
                f"Modelos permitidos: {config.allowed_models}."
            )

        return True, ""