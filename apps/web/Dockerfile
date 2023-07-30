FROM node:18 AS deps

WORKDIR /app
COPY package.json yarn.lock .
RUN yarn install --frozen-lockfile


FROM node:18

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN yarn build

CMD ["yarn", "start"]
