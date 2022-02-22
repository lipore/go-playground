Go Dynamic Library Playground
=============================

## Definitions

The lib and the program must use the same definition. It will panic with error,  ` interface conversion: plugin.Symbol is func() main.Driver, not func() main.Driver (types from different scopes)` when define Driver in lib/main.go and main.go separately. 

## Lib

The `lib` folder contains the code of dynamic library.

### build
```shell
go build -buildmode=plugin -o driver.so lib/main.go
```

## Program

The `main.go` load the dynamic library from `driver.so` and call the function `NewDriver`

### build & run
```shell
go build -o go-dynamicLibrary && ./go-dynamicLibrary
```