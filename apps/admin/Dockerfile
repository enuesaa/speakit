FROM node:20

WORKDIR /workspace
COPY . .

WORKDIR /workspace/apps/admin
RUN npm install -g pnpm
RUN pnpm config set store-dir /.pnpm-store
RUN pnpm install
RUN pnpm build

CMD ["pnpm", "dev"]
