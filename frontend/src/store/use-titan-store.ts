// frontend/src/store/use-titan-store.ts
import { create } from "zustand";

interface TitanStore {
  activeRequests: number;
  totalTokensSaved: number;
  blockedThreats: number;
  setActiveRequests: (value: number) => void;
  setTotalTokensSaved: (value: number) => void;
  setBlockedThreats: (value: number) => void;
  incrementBlockedThreats: () => void;
}

export const useTitanStore = create<TitanStore>((set) => ({
  activeRequests: 0,
  totalTokensSaved: 0,
  blockedThreats: 0,
  setActiveRequests: (value) => set({ activeRequests: value }),
  setTotalTokensSaved: (value) => set({ totalTokensSaved: value }),
  setBlockedThreats: (value) => set({ blockedThreats: value }),
  incrementBlockedThreats: () =>
    set((state) => ({ blockedThreats: state.blockedThreats + 1 })),
}));