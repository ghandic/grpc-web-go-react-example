FROM node

WORKDIR /app

RUN npm i -g pnpm
COPY package.json /app
COPY pnpm-lock.yaml /app

RUN	cd /app && pnpm install

COPY . /app

EXPOSE 5173
CMD ["pnpm", "run", "dev", "--", "--host"]
