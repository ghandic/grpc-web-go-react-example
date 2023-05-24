FROM node

RUN apt-get update && \
    apt-get install ca-certificates 
WORKDIR /app

COPY package.json /app
COPY pnpm-lock.yaml /app
RUN	cd /app && npm install

COPY . /app

EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host"]
