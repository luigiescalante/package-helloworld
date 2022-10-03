# GO Private Repository
Esta librería es un ejemplo para poder descargar un repositorio privado desde GitHub.

## Descripción
Es un ejemplo de una libreria que solo devuelve un método "Hello World" con el objetivo de poder probar el uso de repositorios privados dentro de GO

Fuente: https://codeandoseguro.tech/usar-repositorios-privados-de-github-en-go/

## Ejemplo en GO :
Archivo en el proyecto donde se descarga, o probar directamente en  main.go
<pre>
<code>
package main

import (
	"github.com/luigiescalante/package-helloworld/hellomessage"
	"log"
)

func main() {
	helloSrv := hellomessage.HelloMessage{}
	hello := helloSrv.SayHello("Codeando seguro")
	log.Println(hello)
}
</code>
</pre>

### Autor:
Luis Antonio Escalante Gutiérrez
