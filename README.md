### Prerequisites

```
go 1.15.8
```

### Installing

A step by step series of examples that tell you how to get a development env running, after clonning repository into your machine:

After clonning repository into your machine you need to install required libraries:
```bash
#: Install packages
go mod download

#: Environment should have those variables.
PORT=
JWT_SECRET=
JWT_TOKEN_TTL=

#: Run application
go run main.go
```
