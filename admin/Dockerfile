FROM node:20

WORKDIR /app

COPY admin .

RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build

CMD ["pnpm", "start"]
