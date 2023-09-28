# servicio-rutas

## Instalacion de paqueteria

### `go mod tidy`

## Correr Proyecto
 Se tiene que estar en la raiz del proyecto y ejecutar
### `go run cmd/main.go`


## Configuracion
todas las configuraciones dentro del proyecto estan en `./data/config.yml`


## Construir Proyecto en docker

dentro de la raiz ejecutar
`docker build -t nombre_de_la_imagen .` 
con eso se contruye la imagen

# Correr Proyecto en docker
`docker run -p PORT:PORT nombre_de_la_imagen` 
con eso se corre el proyecto

tener en cuenta  que necesita db, para eso seria bueno ver el siguiente repositorio donde explicar el funcionamiento de todo el ecosistema.

## Orquestador
 https://github.com/time-to-visit/orquestador
