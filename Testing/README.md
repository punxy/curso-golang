# Testing

para realizar el testing de nuestras funciones crearemos un archivo nuevo terminado en `_test.go`, en este caso se llamará `mate_test.go`.

las funciones comenzaran con la palabra `Test` y luego el nombre de la función a testear, en este caso llamaremos a nuestra primera prueba `TestSuma`, siendo `Suma` la función que vamos a testear.

cómo primer parametro de nuestra función `TestSuma` será `(t *testing.T)`, e importamos la librería `testing`

```
    import "testing"

    func TestSuma(t *testing.T){
```

podremos correr las pruebas en la terminal con el siguiente comando

```
    go test
```

## Coverage
para obtener en la terminal el porcentaje de covertura
```
    go test -cover
```

para generar un archivo con el coverage
```
    go test -coverprofile=coverage.out
```

con este archivo podemos ejecutar el comando 
```
    go tool cover -func=coverage.out
```
y observar de una manera más detallada y entendible el porcentaje de covertura
también podemos ver el detalle en html

```
    go tool cover -html=coverage.out
```

## testeo de tiempos
podemos generar un informa con la información de los tiempos de ejecucción,
con el siguiente comando podemos generar el archivo correspondiente
```
    go test -cpuprofile=cpu.out
```
para poder ver la información de este archivo es necesario acceder a el por medio de la terminal
```
    go tool pprof cpu.out
```
una vez dentro de la terminal que nos proporciona el archivo `cpu.out` podemos ejecutar los siguientes comandos para obtener los tiempos de ejecucción:
 - `top`: nos muestra una lista con los métodos que demoran más en ejecutarse
 - `list {search} `: nos muestra una lista indicando en qué linea se demora más y cuanto
    ```
        list Fibonacci
    ```
 - `web`: veremos de una manera más gráfica los tiempos de ejecución (svg), es necesario tener instalado `graphviz`
 - `pdf`: generamos la misma imagen que en `web` pero en un archivo pdf
 - `quit`: salimos del archivo `cpu.out`

