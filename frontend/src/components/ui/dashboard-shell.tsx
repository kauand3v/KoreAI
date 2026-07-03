// src/components/ui/dashboard-shell.tsx
import { ReactNode } from 'react';

interface DashboardShellProps {
  title: string;
  children: ReactNode;        // ← Esta prop é obrigatória
}

export function DashboardShell({ title, children }: DashboardShellProps) {
  return (
    <div className="flex min-h-screen flex-col">
      {/* Header */}
      <header className="border-b bg-white">
        <div className="flex h-16 items-center px-6">
          <h1 className="text-2xl font-semibold">{title}</h1>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 p-6">
        {children}
      </main>
    </div>
  );
}