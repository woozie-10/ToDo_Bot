FROM golang
WORKDIR /app
COPY . .
ENV MYSQL_DATABASE  todo_db
ENV MYSQL_USER user
ENV MYSQL_PASSWORD password
ENV MYSQL_ROOT_PASSWORD Password123#@!
ENV TOKEN your_token
RUN go get
COPY . .
RUN go build -o app
CMD ["./app"]