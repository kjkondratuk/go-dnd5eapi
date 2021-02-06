# go-dnd5eapi

A go client for interacting with the [dnd5eapi](https://www.dnd5eapi.co/).  This is a first-pass at a go client with minimal dependencies, so there will likely be plenty of bugs at this stage--PRs are welcome.

## Install
```bash
go get github.com/kjkondratuk/go-dnd5eapi
```

## Create a client
Simply:
```go
client := godnd5eapi.NewApiClient("<Base URL Here>"))
```

Or, create a client with a client with a custom config:
```go
myCustomClient := &http.Client{
    // ... custom client options here ...
}

client := godnd5eapi.NewApiClient("<Base URL Here>", godnd5eapi.WithHttpClient(myCustomClient))
```

