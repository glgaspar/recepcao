FROM golang:1.23.1 AS base

WORKDIR /app

FROM node:20 AS build
WORKDIR /src
COPY . .

RUN npm install
RUN npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css

FROM base AS final 
WORKDIR /app

COPY go.mod go.sum . ./
COPY --from=build ./src .

RUN CGO_ENABLED=0 GOOS=linux go build -o /recepcao

CMD ["/recepcao"]