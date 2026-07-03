// frontend/src/components/charts/latency-chart.tsx
"use client";

import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const data = [
  { time: "00:00", latency: 45 },
  { time: "01:00", latency: 52 },
  { time: "02:00", latency: 38 },
  { time: "03:00", latency: 65 },
  { time: "04:00", latency: 48 },
  { time: "05:00", latency: 55 },
  { time: "06:00", latency: 42 },
];

export function LatencyChart() {
  return (
    <div className="bg-zinc-900/50 rounded-lg border border-zinc-800 p-4">
      <h3 className="text-sm font-medium text-zinc-400 mb-4">Proxy Latency (ms)</h3>
      <ResponsiveContainer width="100%" height={250}>
        <LineChart data={data}>
          <CartesianGrid stroke="#27272a" strokeDasharray="3 3" vertical={false} />
          <XAxis
            dataKey="time"
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
          <Line
            type="monotone"
            dataKey="latency"
            stroke="#f59e0b"
            strokeWidth={2}
            dot={false}
            name="Latency"
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}