# Kio

Kio (Keep It Organized - inspired to Boilr) let manage simple template creation logic

  - Prepare your input file (with logic)
  - (Go Template) Magic
  - Profit!

### Why
Because i need i simple template creation tool that follow a simple logic.

### Installation

```sh
$ go get github.com/r3vit/kio/kio
```

### Template logic

Choice between `n`-options:
```sh
{{ choose "referenceName" "option1" "option2" "..." "optionN"}}
```

Show (`yes`/`no`):
```sh
    {{ if isset "referenceName"}}
        //show if referenceName is set (0 or 1)
    {{ end }}
```

Custom:
```sh
{{ custom "referenceName" }}
```


### Code

Simply:
```go
	newTemplate := kio.NewKio("Template-Name", "input.yml", "output.txt")
	newTemplate.Execute()
```
