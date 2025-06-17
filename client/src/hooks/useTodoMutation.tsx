import { useMutation, useQueryClient } from "@tanstack/react-query";

export function useTodoMutation(
  key: string,
  method: "PATCH" | "DELETE",
  url: string
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [key],
    mutationFn: async () => {
      const res = await fetch(url, { method });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || `${key} failed`);
      return data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      console.error(`${key} error:`, error);
    },
  });
}
