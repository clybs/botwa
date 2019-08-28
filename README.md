# botwa
Breath of the Wild Articles

### Installation

- Requires latest [Golang](https://golang.org/doc/install) and
- Requires latest [Docker](https://docs.docker.com) to run
- Requires latest [Docker Compose](https://docs.docker.com/compose/install/) to run

### Run application
Go to project folder and type:

```sh
$ cd botwa
$ docker-compose up -d
```

### Shutdown application
Go to project folder and type:

```sh
$ cd botwa
$ docker-compose down
```

### Interact with the app
List articles

```sh
$ curl http://localhost:8080/articles
```

Read an article
```sh
$ curl http://localhost:8080/articles/{article_id}
```

Create and article.
```sh
$ curl -X POST localhost:8080/articles -d '{"title":"hello", "content":"world", "author":"cliburn"}'
```

### Tests
Run the tests:
```sh
$ go test -cover ./... -v
```
### Documentation
Run the docs:
```sh
$ godoc -http=":6060"
```
Then visit: [http://localhost:6060/pkg/github.com/clybs/](http://localhost:6060/pkg/github.com/clybs/)
 