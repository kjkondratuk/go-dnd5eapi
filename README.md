# go-dnd5eapi

A go client for interacting with the [dnd5eapi](https://www.dnd5eapi.co/).

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

## Supported Endpoints
*(Work-in-progress)*

| Endpoint | Fully Implemented | Partially Implemented | Not Implemented |
| --- | :---: | :---: | :---: |
| endpoints | X | - | - |
| ------------------------------- | --- | --- | --- |
| ability-scores | - | X | - |
| classes | - | X | - |
| proficiencies | - | X | - |
| skills | - | X | - |
| ------------------------------- | --- | --- | --- |
| conditions | - | - | - |
| damage-types | - | - | - |
| equipment-categories | - | - | - |
| equipment | - | - | - |
| features | - | - | - |
| languages | - | - | - |
| magic-schools | - | - | - |
| monsters | - | - | - |
| races | - | - | - |
| spellcasting | - | - | - |
| spells | - | - | - |
| starting-equipment | - | - | - |
| subclasses | - | - | - |
| subraces | - | - | - |
| traits | - | - | - |
| weapon-properties | - | - | - |
