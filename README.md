## REA Group code test

## Development

Install Go if not already installed `https://golang.org/doc/install`

Go to the the root directory of project and run the command

```bash
$ go run .
```
Which will use the test data inside the file  named `data.txt`.

To use data from input commands on the terminal instead, just pass the flag `--false`

```bash
$ go run . --file=false
```

To run (very) basic tests run

```bash
$ go test
```
[Still in progress]