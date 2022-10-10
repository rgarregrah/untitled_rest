import { Flex, Box, Spacer, Button } from "@chakra-ui/react";
import { Text } from "@chakra-ui/react";
import Status from "./status";
import { formatDate } from "../util/utils";
import { CloseIcon } from "@chakra-ui/icons";

type Props = {
    todo: {
        id: number;
        body: string;
        updated_at: string;
        status: string;
        detail: {
            color: string;
        };
    };
    remove: any;
};
const Todo = ({ todo, remove }: Props) => {
    const removeTodo = (id: number, remove: any) => {
        remove(id);
    };

    return (
        <Box bg="#3700B3" height="60px" width="100%">
            <Flex m={2}>
                <Text color="#fff">{todo.body}</Text>
                <Spacer />
                <Status status={todo.status}></Status>
                <Text color="#fff">{formatDate(todo.updated_at)}</Text>
                <CloseIcon
                    ml={3}
                    color="#fff"
                    onClick={(e: any) => removeTodo(todo.id, remove)}
                />
            </Flex>
        </Box>
    );
};

export default Todo;
