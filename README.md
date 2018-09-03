# Vistar Media State API

## Development

### Local
This project requires [Docker](https://docs.docker.com/v17.12/docker-for-mac/install/)

```
$ cp .env.example .env
$ make dev
```

## Testing
If your server isn't started, bring it up with `make dev`. Then to execute the tests:
```
$ make test
```

## Strategy

### Fixtures
I created a [CLI](https://github.com/emmac1016/state-api/blob/dev/cli/fixtures.go) which executed the command `./fixtures`. It was created by executing the following:
```
$ go build cli/fixtures.go
$ ./fixtures
```
The command makes use of the [Fixture Loader](https://github.com/emmac1016/state-api/blob/dev/internal/loader.go). This command is executed during the `make dev` command.

### API
I knew that MongoDB had geospatial searching capabilities, so the biggest task I had to do was to translate the data in `states.json` to be in a valid GeoJSON format. I then loaded it into the db using the strategy described above in "Fixtures". Then the actual API just had to call a query that checks the db for any state where a given pair of coordinates lies in their borders.
