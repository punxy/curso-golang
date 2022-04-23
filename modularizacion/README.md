# Modularización

para poder importar bien los paquetes y trabajar de manera correcta en visual studio code, es necesario abrir el proyecto desde la carpeta donde se ejecuta el main.go

## struct
para poder incicializar los struct, estos deben tener sus propiedades en modo público
```
    type Circulo struct {
        Radio float32
    }
```

## Módulos de terceros

primero debemos obtener el paquete
```
    go get github.com/donvito/hellomod
```
se creará un nuevo archivo `go.sum` con los hash correspondientes,
luego se debe hacer un import 
```
    import github.com/donvito/hellomod
```

## Creación de módulos propios
primero iniciamos un repositorio en github, 
el package a crear será la de github
```
    go mod init github.com/punxy/figuras
```

## Paquetes
los paquetes de terceros o propios son descargargados donde se creó la carpeta `go/pkg`