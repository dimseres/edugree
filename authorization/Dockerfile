FROM golang:bullseye
#RUN apt install nodejs npm
RUN mkdir "/app"
WORKDIR /app
ADD go.mod ./
ADD go.sum ./
RUN go mod download
COPY . .
#RUN npm install ./cmd/migration
#RUN npx prisma dev --name init --prefix ./cmd/migration
RUN go build -o bin/auth/service ./cmd/auth-service
EXPOSE 7001
CMD ["./bin/auth/service"]
