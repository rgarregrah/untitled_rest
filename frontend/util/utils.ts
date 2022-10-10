export const formatDate = (str: string) => {
    const date = new Date(str);
    date.setHours(date.getHours() - 9);

    const year = date.getFullYear();
    const month = date.getMonth().toString().padStart(2, "0");
    const day = date.getDate().toString().padStart(2, "0");

    const hour = date.getHours().toString().padStart(2, "0");
    const minutes = date.getMinutes().toString().padStart(2, "0");
    const seconds = date.getSeconds().toString().padStart(2, "0");

    return `${year}/${month}/${day} ${hour}:${minutes}:${seconds}`;
};
