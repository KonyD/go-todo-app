import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { Todo } from "../components/TodoList";
import { BASE_URL } from "../App";

export function useTodoMutation(
  key: string,
  method: "PATCH" | "DELETE",
  todo: Todo
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [key],
    mutationFn: async () => {
      if (method === "PATCH" && todo.completed) return alert("Todo is already completed");
      
      const res = await fetch(`${BASE_URL}/todos/${todo._id}`, { method });
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
