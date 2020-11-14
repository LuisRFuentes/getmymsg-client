# getmymsg-client

Cliente desarrollado por Luis Fuentes para la práctica de la Unidad 3 del curso de Sistemas Distribuidos

## Requerimientos
- Acceso al servidor getmymsg-server
- Tener Go instalado

## Instalación del cliente

1. Realice el clone de este repositorio en su maquina

```
$ git clone https://github.com/LuisRFuentes/lf-getmymsg-client.git
```

2. Usando Go, compile el archivo lf-getmymsg-client.go usando el comando:

```
$ go build lf-getmymsg-client.go
```

3. Ejecute el archivo ejecutable que le genere la compilación del archivo

```
$ ./lf-getmymsg-client
```

En caso que quiera ejecutar el cliente sin necesidad de generar un archivo ejecutable puede usar el siguiente comando:

```
$ go run lf-getmymsg-client.go
```

Esto compilará y ejecutará el código del archivo lf-getmymsg-client.go sin necesidad de usar los comandos de los pasos 2 y 3

## Uso del cliente

El cliente recibirá los parámetros necesarios para su funcionamiento durante la ejecución de este. Al iniciar le pedirá al usuario que ingrese los siguientes parámetros:

- Nombre de usuario
- Dirección IP del servidor
- Dirección IP desde donde esta ejecutando el cliente
- Puerto por donde desea reciibir el mensaje

Después de recibir esos parámetros, el cliente se comunicará con el servidor para autenticarse, recibir el mensaje y verificar que este último haya llegado completo. Después de esto, el cliente cerrará la conexión con el servidor y finalizará su ejecución.
