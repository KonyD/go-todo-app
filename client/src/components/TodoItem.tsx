import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import type { Todo } from "./TodoList";
import { BASE_URL } from "../App";
import { useTodoMutation } from "../hooks/useTodoMutation";
import {useColorModeValue} from "./ui/color-mode"

const TodoItem = ({ todo }: { todo: Todo }) => {
  const updateUrl = `${BASE_URL}/todos/${todo._id}`;
  const deleteUrl = `${BASE_URL}/todos/${todo._id}`;

  const { mutate: updateTodo, isPending: isUpdating } = useTodoMutation(
    "updateTodo",
    "PATCH",
    updateUrl
  );

  const { mutate: deleteTodo, isPending: isDeleting } = useTodoMutation(
    "deleteTodo",
    "DELETE",
    deleteUrl
  );

  return (
    <Flex gap={2} alignItems={"center"}  bg={useColorModeValue("blue.200", "gray.900")}>
      <Flex
        flex={1}
        alignItems={"center"}
        border={"1px"}
        borderColor={"gray.600"}
        p={2}
        borderRadius={"lg"}
        justifyContent={"space-between"}
      >
        <Text
          color={todo.completed ? useColorModeValue("green.500", "green.200") : useColorModeValue("yellow.400", "yellow.100")}
          textDecoration={todo.completed ? "line-through" : "none"}
        >
          {todo.body}
        </Text>
        {todo.completed && (
          <Badge ml="1" colorPalette="green">
            Done
          </Badge>
        )}
        {!todo.completed && (
          <Badge ml="1" colorPalette="yellow">
            In Progress
          </Badge>
        )}
      </Flex>
      <Flex gap={2} alignItems={"center"}>
        <Box
          color={"green.500"}
          cursor={"pointer"}
          onClick={() => updateTodo()}
        >
          {!isUpdating && <FaCheckCircle size={20} />}
          {isUpdating && <Spinner size={"sm"} />}
        </Box>
        <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
          {!isDeleting && <MdDelete size={25} />}
          {isDeleting && <Spinner size={"sm"} />}
        </Box>
      </Flex>
    </Flex>
  );
};
export default TodoItem;
