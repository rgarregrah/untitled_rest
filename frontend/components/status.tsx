import { Tag } from "@chakra-ui/react";

type Props = {
    status: string;
};

const getColor = (status: string) => {
    if (status === "Doing") {
        return "yellow";
    }
    if (status === "Done") {
        return "green";
    }
    if (status === "Removed") {
        return "red";
    }
    return "purple";
};

const Status = ({ status }: Props) => {
    return (
        <Tag
            borderRadius="full"
            variant="solid"
            colorScheme={getColor(status)}
            mr={5}
        >
            {status}
        </Tag>
    );
};

export default Status;
