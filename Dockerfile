FROM node:alpine
COPY . /app
WORKDIR /app
RUN npm install
RUN npm run build

FROM golang
COPY . /app
WORKDIR /app
RUN go build api/main.go

EXPOSE 8080
ENTRYPOINT [ "/app/main" ]