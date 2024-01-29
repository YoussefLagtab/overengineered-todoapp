# Overengineered TodoApp:

This is a fun project to discover `grpc` protocol, and Golang.

# How to run:

- copy `.env.example` to `.env` and change the values, or not :)
- run: ``env `cat .env | grep -v ^#` docker compose up -d``

# Endpoints:

-  `GET /todos`
-  `GET /todos/:id`
-  `POST /todos # (body: {content: string}) `
-  `PATCH /todos/:id # (body: {content: string}) `
-  `PATCH /todos/:id/mark-as-complete`
-  `PATCH /todos/:id/mark-as-incomplete`
-  `DELETE /todos/:id`

# Future ideas:

- add `users` and `auth` services, both will be written in a different language.

# Note:

- feedback is very welcomed