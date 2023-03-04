FROM golang:1.20-alpine AS build

WORKDIR /src

COPY . ./

RUN go mod download && \
    go build -o /out/husakgpt

FROM alpine

COPY --from=build /out/husakgpt /bin

CMD [ "/bin/husakgpt" ]
