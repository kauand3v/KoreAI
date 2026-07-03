// frontend/src/app/dashboard/page.tsx
"use client";

import { DashboardShell } from "@/components/ui/dashboard-shell";
import { GuardianStatus } from "@/components/security/guardian-status";
import { TokenUsageChart } from "@/components/charts/token-usage-chart";
import { LatencyChart } from "@/components/charts/latency-chart";
import { CreatePolicyForm } from "@/components/forms/create-policy-form";
import { useTelemetry } from "@/hooks/use-telemetry";

export default function DashboardPage() {
  useTelemetry(); // starts the telemetry simulation

  return (
    <DashboardShell title="Overview">
      <div className="space-y-6">
        <GuardianStatus />
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <TokenUsageChart />
          <LatencyChart />
        </div>
        <div className="max-w-lg">
          <h3 className="text-sm font-medium text-zinc-400 mb-3">Create New Policy</h3>
          <CreatePolicyForm />
        </div>
      </div>
    </DashboardShell>
  );
}