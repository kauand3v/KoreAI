// frontend/src/schemas/kore-policy-schema.ts
import { z } from "zod";

export const korePolicySchema = z.object({
  routeName: z.string().min(3, "Route name must be at least 3 characters"),
  priority: z.number().min(0).max(100),
  model: z.string().min(1, "Model is required"),
  fallback: z.string().optional(),
  cache: z.boolean(),
  circuitBreaker: z.object({
    threshold: z.number().min(0).max(100),
    duration: z.number().min(1),
  }),
});

export type KorePolicy = z.infer<typeof korePolicySchema>;