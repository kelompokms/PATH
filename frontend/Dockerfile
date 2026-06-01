FROM ghcr.io/pnpm/pnpm:11 AS builder
RUN pnpm runtime set node 24 -g

WORKDIR /app
COPY package*.json ./

RUN pnpm install

COPY . .
RUN pnpm build
RUN pnpm prune --production

FROM node:24-alpine
WORKDIR /app
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .
EXPOSE 3001
ENV NODE_ENV=production
