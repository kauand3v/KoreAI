// frontend/src/components/forms/create-policy-form.tsx
"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { korePolicySchema, type KorePolicy } from "@/schemas/kore-policy-schema";
import { cn } from "@/lib/utils";

export function CreatePolicyForm() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<KorePolicy>({
    resolver: zodResolver(korePolicySchema),
    defaultValues: {
      cache: false,
      circuitBreaker: { threshold: 50, duration: 30 },
    },
  });

  const onSubmit = (data: KorePolicy) => {
    console.log("Policy created:", data);
    // TODO: integrate with backend
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-4 p-4 bg-zinc-900/50 rounded-lg border border-zinc-800"
    >
      <div>
        <label htmlFor="routeName" className="block text-sm font-medium text-zinc-400">
          Route Name
        </label>
        <input
          id="routeName"
          {...register("routeName")}
          className={cn(
            "mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600",
            errors.routeName && "border-red-500"
          )}
          placeholder="e.g., /api/chat"
        />
        {errors.routeName && (
          <p className="mt-1 text-xs text-red-500">{errors.routeName.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="priority" className="block text-sm font-medium text-zinc-400">
          Priority (0-100)
        </label>
        <input
          id="priority"
          type="number"
          {...register("priority", { valueAsNumber: true })}
          className={cn(
            "mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600",
            errors.priority && "border-red-500"
          )}
          placeholder="50"
        />
        {errors.priority && (
          <p className="mt-1 text-xs text-red-500">{errors.priority.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="model" className="block text-sm font-medium text-zinc-400">
          Model
        </label>
        <input
          id="model"
          {...register("model")}
          className={cn(
            "mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600",
            errors.model && "border-red-500"
          )}
          placeholder="gpt-4"
        />
        {errors.model && (
          <p className="mt-1 text-xs text-red-500">{errors.model.message}</p>
        )}
      </div>

      <div>
        <label htmlFor="fallback" className="block text-sm font-medium text-zinc-400">
          Fallback (optional)
        </label>
        <input
          id="fallback"
          {...register("fallback")}
          className="mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600"
          placeholder="gpt-3.5-turbo"
        />
      </div>

      <div className="flex items-center gap-2">
        <input
          id="cache"
          type="checkbox"
          {...register("cache")}
          className="rounded border-zinc-700 bg-zinc-900 text-zinc-100 focus:ring-zinc-600"
        />
        <label htmlFor="cache" className="text-sm text-zinc-400">
          Enable Cache
        </label>
      </div>

      <fieldset className="border border-zinc-800 rounded-md p-3">
        <legend className="text-sm text-zinc-400 px-1">Circuit Breaker</legend>
        <div className="space-y-3">
          <div>
            <label htmlFor="circuitBreaker.threshold" className="block text-sm font-medium text-zinc-400">
              Threshold (%)
            </label>
            <input
              id="circuitBreaker.threshold"
              type="number"
              {...register("circuitBreaker.threshold", { valueAsNumber: true })}
              className={cn(
                "mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600",
                errors.circuitBreaker?.threshold && "border-red-500"
              )}
              placeholder="50"
            />
            {errors.circuitBreaker?.threshold && (
              <p className="mt-1 text-xs text-red-500">{errors.circuitBreaker.threshold.message}</p>
            )}
          </div>
          <div>
            <label htmlFor="circuitBreaker.duration" className="block text-sm font-medium text-zinc-400">
              Duration (seconds)
            </label>
            <input
              id="circuitBreaker.duration"
              type="number"
              {...register("circuitBreaker.duration", { valueAsNumber: true })}
              className={cn(
                "mt-1 w-full bg-zinc-900 border border-zinc-800 rounded-md px-3 py-2 text-sm text-zinc-100 placeholder-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-600",
                errors.circuitBreaker?.duration && "border-red-500"
              )}
              placeholder="30"
            />
            {errors.circuitBreaker?.duration && (
              <p className="mt-1 text-xs text-red-500">{errors.circuitBreaker.duration.message}</p>
            )}
          </div>
        </div>
      </fieldset>

      <button
        type="submit"
        className="w-full bg-zinc-800 hover:bg-zinc-700 text-zinc-100 font-medium py-2 px-4 rounded-md transition-colors"
      >
        Create Policy
      </button>
    </form>
  );
}