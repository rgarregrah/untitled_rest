import type { NextPage } from "next";
import { Flex, VStack } from "@chakra-ui/react";
import Head from "next/head";
import Todo from "../components/todo";
import InputBody from "../components/inputBody";
import useSWR, { mutate } from "swr";
import axios from "axios";

type Todo = {
    id: number;
    body: string;
    updated_at: string;
    status: string;
    detail: {
        color: string;
    };
};

const Todos: NextPage = () => {
    const submitTodo = async (id: number, body: string, status: string) => {
        await axios.post(`http://localhost:8080/api/todo/${id}`, {
            id: id,
            body: body,
            status: status,
        });
        mutate("http://localhost:8080/api/todos");
    };

    const deleteTodo = async (id: number) => {
        await axios.delete(`http://localhost:8080/api/todo/${id}`);
        mutate("http://localhost:8080/api/todos");
    };

    const fetcher = async (key: string, init?: RequestInit) => {
        return fetch(key, init).then(
            (res) => res.json() as Promise<Todo[] | null>
        );
    };

    const { data: todos, error } = useSWR(
        "http://localhost:8080/api/todos",
        fetcher
    );

    if (!todos) {
        return <div>loading...</div>;
    }

    return (
        <div>
            <Head>
                <title>Todo</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main>
                <VStack>
                    <Flex mt={50} w="100%" justify="center" maxWidth={"960px"}>
                        <VStack w="80%">
                            <InputBody submit={submitTodo} />
                        </VStack>
                    </Flex>

                    <Flex
                        mt={50}
                        w={"100%"}
                        maxWidth={"960px"}
                        justify="center"
                    >
                        <VStack w="80%" mt={50}>
                            {todos.map((todo) => {
                                return (
                                    <Todo
                                        todo={todo}
                                        remove={deleteTodo}
                                        key={todo.id}
                                    ></Todo>
                                );
                            })}
                        </VStack>
                    </Flex>
                </VStack>
            </main>
        </div>
    );
};

export default Todos;
