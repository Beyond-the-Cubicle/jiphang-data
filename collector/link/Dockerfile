FROM node:20-slim AS runner

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

RUN apt-get update -y && apt-get install -y openssl

COPY . /app
WORKDIR /app

RUN pnpm install
CMD ["./entrypoint.sh"]
