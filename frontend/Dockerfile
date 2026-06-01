FROM ghcr.io/pnpm/pnpm:11 AS builder
RUN pnpm runtime set node 24 -g

WORKDIR /app
COPY package*.json ./

RUN pnpm install

COPY . .
RUN pnpm build
RUN pnpm prune --production

FROM nginx:1.30-alpine
COPY --from=builder /app/build /usr/share/nginx/html
