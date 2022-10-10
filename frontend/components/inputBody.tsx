import { Textarea } from "@chakra-ui/react";
import React from "react";

type Props = {
    submit: any;
};

const submitTodo = (
    e: React.KeyboardEvent<HTMLTextAreaElement>,
    setBody: React.Dispatch<React.SetStateAction<string>>,
    body: string,
    submit: any
) => {
    if (e.nativeEvent.isComposing || e.key !== "Enter" || e.shiftKey) {
        return;
    }
    submit(0, body, "Untouched");
    setBody("");
    e.preventDefault();
};

const InputBody = ({ submit }: Props) => {
    const [body, setBody] = React.useState("");
    return (
        <Textarea
            placeholder="something to do"
            value={body}
            onChange={(e) => setBody(e.target.value)}
            onKeyDown={(e) => submitTodo(e, setBody, body, submit)}
        />
    );
};

export default InputBody;
