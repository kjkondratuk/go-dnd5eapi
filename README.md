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
| classes-proficiency-rel | - | X | - |
| damage-types | - | X | - |
| equipment | - | X | - |
| languages | - | X | - |
| proficiencies | - | X | - |
| races | - | X | - |
| skills | - | X | - |
| skill-ability-rel | - | X | - |
| subclasses | - | X | - |
| traits | - | X | - |
| ------------------------------- | --- | --- | --- |
| conditions | - | - | - |
| equipment-categories | - | - | - |
| features | - | - | - |
| magic-schools | - | - | - |
| monsters | - | - | - |
| spellcasting | - | - | - |
| spells | - | - | - |
| starting-equipment | - | - | - |
| subraces | - | - | - |
| weapon-properties | - | - | - |
