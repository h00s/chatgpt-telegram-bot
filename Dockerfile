FROM golang:1.21-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/chatgpt-telegram-bot

FROM alpine

COPY --from=build /out/chatgpt-telegram-bot /bin

CMD [ "/bin/chatgpt-telegram-bot" ]
