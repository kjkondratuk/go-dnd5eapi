# go-dnd5eapi

![Go](https://github.com/kjkondratuk/go-dnd5eapi/workflows/Go/badge.svg)

A go client for interacting with the [dnd5eapi](https://www.dnd5eapi.co/).  This is a first-pass at a go client with minimal dependencies, so there will likely be plenty of bugs at this stage--PRs are welcome.

## Install
```bash
go get github.com/kjkondratuk/go-dnd5eapi
```

## Create a client
Simply:
```go
client := go_dnd5eapi.NewApiClient("<Base URL Here>"))
```

Or, create a client with a client with a custom config:
```go
myCustomClient := &http.Client{
    // ... custom client options here ...
}

client := go_dnd5eapi.NewApiClient("<Base URL Here>", go_dnd5eapi.WithHttpClient(myCustomClient))
```

## Sample Calls

Check out the examples in the `examples` folder!
