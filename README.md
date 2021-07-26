# Labpro BE

## Prerequisite

- [Oculi](http://github.com/ravielze/oculi)
- Postgresql
- apib2swagger (optional)

## Manual Instalation/Run

1. Copy `.env.example` to `.env`.
2. Setup `.env`. Don't forget to start the database server before moving to the next step.
3. For development, run `go run main.go`. Or, `go build main.go` and run the executable.

## Usage

__All API are documented__ in [localhost/docs](http://localhost/docs).

*(Please run the program in port 80 to use this hyperlink, or directly just access localhost:**YOUR_PORT**/docs)*

## Regenerate Swagger JSON from API Blueprint

1. Install api2bswagger using `npm install -g apib2swagger`
2. Run `apib2swagger -i ./api/blueprint.apib -o ./resources/external/docs.json`

## Author

Coded with <3 by [ravielze](https://github.com/ravielze)