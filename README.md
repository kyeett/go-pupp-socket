
# go-pupp-socket
Toy project to learn websockets and go


### Lessons learnt
#### `flag` package
Supports basic command line flag parsing. 

Creates a new flag `addr` with default value :8081 and help text
```
    addr = flag.String("addr", ":8081", "http service addr")
```

[gobyexample: command-line-flags](https://gobyexample.com/command-line-flags)

#### `html/template` package
