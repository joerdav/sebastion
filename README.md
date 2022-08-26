# sebastion

A simple framework for running Go scripts.

This is a work in progress still.

Run the example:

Web:

```sh
go run examples/web/main.go
```

TUI:

```sh
go run examples/tui/main.go
```


## Tasks

These tasks follow [xc](https://github.com/joerdav/xc) syntax, therefore can be ran with `xc [taskname]`.

### tools

Install development tools.

- <https://github.com/a-h/templ>: A Go HTML templating language.
- <https://github.com/cespare/reflex>: For watch tasks

```shell
go install github.com/a-h/templ/cmd/templ@v0.2.184
go install github.com/cespare/reflex@v0.3.1
```

### web-example

Run the web example located at `examples/web/main.go`.

```shell
reflex -r '.*\.(go|templ)' -R '.*_templ\.go' -s -- sh -c 'templ generate && go run examples/web/main.go'
```

### tui-example

Run the cli example located at `examples/tui/main.go`.

```shell
go run examples/tui/main.go
```

### build-templates

Build all [templ](https://github.com/a-h/templ) templates.

```shell
templ generate
```
