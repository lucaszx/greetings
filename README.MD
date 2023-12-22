#Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalisados

## Instalación 

Ejectutar el siguiente comando para instalar el paquete:

```bash
go get -u github.com/lucaszx/greetings
```

## Uso 
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:


```go
package main

import (

	"fmt"
	"github.com/lucaszx/greetings"
)

func main() {

	message, err := greetings.Hello("Lucas")

	if err != nil {
		fmt.Println("Ocurrió un error:",err)
        return
	}

	fmt.Println(message)
}
```

Este ejemplo importa el paquete github.com/lucaszx/greetings y llama a la función Hello la cual
retorna un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.