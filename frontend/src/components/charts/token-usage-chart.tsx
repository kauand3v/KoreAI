// frontend/src/components/charts/token-usage-chart.tsx
"use client";

import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const data = [
  { name: "00:00", used: 1200, saved: 300 },
  { name: "01:00", used: 1800, saved: 450 },
  { name: "02:00", used: 1500, saved: 400 },
  { name: "03:00", used: 2100, saved: 600 },
  { name: "04:00", used: 2500, saved: 750 },
  { name: "05:00", used: 2200, saved: 700 },
  { name: "06:00", used: 2800, saved: 850 },
];

export function TokenUsageChart() {
  return (
    <div className="bg-zinc-900/50 rounded-lg border border-zinc-800 p-4">
      <h3 className="text-sm font-medium text-zinc-400 mb-4">Token Usage (last 6h)</h3>
      <ResponsiveContainer width="100%" height={250}>
        <AreaChart data={data}>
          <defs>
            <linearGradient id="usedGrad" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#8b5cf6" stopOpacity={0.3} />
              <stop offset="95%" stopColor="#8b5cf6" stopOpacity={0} />
            </linearGradient>
            <linearGradient id="savedGrad" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#10b981" stopOpacity={0.3} />
              <stop offset="95%" stopColor="#10b981" stopOpacity={0} />
            </linearGradient>
          </defs>
          <CartesianGrid stroke="#27272a" strokeDasharray="3 3" vertical={false} />
          <XAxis
            dataKey="name"
            tick={{ fill: "#a1a1aa", fontSize: 12 }}
            axisLine={{ stroke: "#27272a" }}
            tickLine={false}
          />
          <YAxis
            tick={{ fill: "#a1a1aa", fontSize: 12 }}
            axisLine={{ stroke: "#27272a" }}
            tickLine={false}
          />
          <Tooltip
            contentStyle={{
              backgroundColor: "#18181b",
              border: "1px solid #27272a",
              borderRadius: "6px",
              color: "#f4f4f5",
              fontSize: "12px",
            }}
            itemStyle={{ color: "#f4f4f5" }}
          />
          <Area
            type="monotone"
            dataKey="used"
            stroke="#8b5cf6"
            fill="url(#usedGrad)"
            strokeWidth={2}
            name="Tokens Used"
          />
          <Area
            type="monotone"
            dataKey="saved"
            stroke="#10b981"
            fill="url(#savedGrad)"
            strokeWidth={2}
            name="Tokens Saved"
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}