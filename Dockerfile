FROM golang
WORKDIR /app
COPY . .

ENV JWT_TOKEN_TTL=5
ENV JWT_SECRET=secret
ENV PORT=8082

RUN go build -v -o dummy-boilerplate

EXPOSE $PORT

CMD ./dummy-boilerplate