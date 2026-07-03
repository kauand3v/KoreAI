"use client";

export function GuardianStatus() {
  return (
    <div className="rounded-lg border border-zinc-800 bg-zinc-900/50 p-4">
      <h3 className="text-sm font-medium text-zinc-400 mb-2">Guardian Status</h3>
      <div className="flex items-center gap-2">
        <span className="h-2 w-2 rounded-full bg-emerald-500" />
        <span className="text-sm text-zinc-300">All tenants isolated</span>
      </div>
    </div>
  );
}