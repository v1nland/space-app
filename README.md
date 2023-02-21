# space-playground

This repository is a boilerplate for golang, using: gorm, clean architecture and graphql.

## Installation

Use go mod to install dependencies.

```bash
go mod tidy
```

This repository assumes you have a local container running postgres:alpine (15.1-alpine at the moment of updating this readme). If you don't have it, please install docker and then run the following command to create an empty database. Space-playground application will take care of running schema migration.

```bash
docker run --name psql-local -e POSTGRES_PASSWORD=test -d -p 5432:5432 postgres:alpine
```

## Usage

Run application.

```bash
go run app/cmd/main.go
```

Then, you can open your browser and go to <http://localhost:8080> to open the playground.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
