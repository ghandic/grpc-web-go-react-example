FROM golang

WORKDIR /app

RUN curl -fsSL https://deb.nodesource.com/setup_20.x | \
    bash - && apt-get install -y nodejs && npm install -g npm@9.6.7
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest 
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1 
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
RUN npm install @bufbuild/buf npx

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . /app
RUN cd /app && npx buf generate
RUN	cd /app && npm install
RUN cd /app && sqlc generate
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
