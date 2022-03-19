FROM golang:1.17

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .

RUN go mod download

COPY api/*.go ./

RUN go build -o /api_challenge

EXPOSE 17854

CMD [ "/api_challenge" ]