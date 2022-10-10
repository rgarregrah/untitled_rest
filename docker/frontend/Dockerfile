FROM node:18.7-slim

WORKDIR /app

COPY ./frontend/package.json /app
RUN yarn install
COPY ./backend .

# CMD [ "yarn", "dev" ]
