# How to install GO en MAC

## Install GO

1. Download GO from [https://golang.org/dl/](https://golang.org/dl/)

2. Install GO

## Hello World program

Create a file called hello.go

```go
package main // Package name, this is about the project package name

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Run the program

```bash
go run hello.go
```

### Build the program

```bash
go build hello.go
```

En go es necesario compilarlo para poder ejecutarlo, en otros lenguajes como python no es necesario.
Pero también existe un comando que nos permite ejecutarlo sin necesidad de compilarlo y es RUN

### Build 

Cuando ejecutamos el comando build se genera un archivo binario que es el que se ejecuta, en este caso se genera un archivo llamado hello

```bash 
./hello
```
