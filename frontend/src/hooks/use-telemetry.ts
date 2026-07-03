// frontend/src/hooks/use-telemetry.ts
"use client";

import { useEffect } from "react";
import { useTitanStore } from "@/store/use-titan-store";

export const useTelemetry = () => {
  const {
    setActiveRequests,
    setTotalTokensSaved,
    setBlockedThreats,
    incrementBlockedThreats,
  } = useTitanStore();

  useEffect(() => {
    const interval = setInterval(() => {
      // Simulate random telemetry data
      const newRequests = Math.floor(Math.random() * 50) + 10;
      const newTokensSaved = Math.floor(Math.random() * 2000) + 500;
      const newBlocked = Math.random() > 0.9 ? 1 : 0;

      setActiveRequests(newRequests);
      setTotalTokensSaved(newTokensSaved);
      if (newBlocked) incrementBlockedThreats();
      else setBlockedThreats(0); // reset if no threat
    }, 2000);

    return () => clearInterval(interval);
  }, [
    setActiveRequests,
    setTotalTokensSaved,
    setBlockedThreats,
    incrementBlockedThreats,
  ]);
};