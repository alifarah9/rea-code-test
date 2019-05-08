## Alinta code test

* App is initialised with four customers.
* To delete a customer select one from the list and hit the delete button on the left.
* To add a customer enter the details in the form fields and hit the "add" button.
* To edit customer select one from the list and hit the "update" button.

## Development

Install Go if not already installed `https://golang.org/doc/install`
Inside the root directory run
```bash
$ go run main.go
```
Which will use the test data inside the file `data.txt`.

To use data from input commands on the terminal just pass the flag `--input`

```bash
$ go run main.go --input
```

[Still in progress]