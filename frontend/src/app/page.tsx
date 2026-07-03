// src/app/dashboard/page.tsx
import { DashboardShell } from '@/components/ui/dashboard-shell';

export default function DashboardPage() {
  return (
    <DashboardShell title="Dashboard">
      {/* Todo o conteúdo da página vai aqui */}
      <div className="space-y-6">
        <div>
          <h2 className="text-xl font-semibold mb-4">Visão Geral</h2>
          {/* Seus componentes, cards, gráficos etc. */}
        </div>

        {/* Exemplo de conteúdo */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-white p-6 rounded-xl shadow">
            <p>Card 1</p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow">
            <p>Card 2</p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow">
            <p>Card 3</p>
          </div>
        </div>
      </div>
    </DashboardShell>
  );
}