# Cloud-Instance-Manager

Este proyecto es un gestor de instancias en la nube minimalista escrito en Go. Permite crear, listar, iniciar, detener y terminar instancias en la nube de manera segura y eficiente. El objetivo de este proyecto es proporcionar una herramienta simple y fácil de usar para administrar instancias en la nube de Amazon Web Services (AWS).

## Características

* Crear instancias en la nube con una sola llamada a la API
* Listar instancias en la nube en una región específica
* Iniciar, detener y terminar instancias en la nube de manera segura
* Modo DryRun para probar el código sin afectar recursos reales
* Soporte para instancias en la nube de bajo costo (t2.micro)
* Soporte para instancias en la nube con diferentes tipos de instancias (t2.micro, t2.small, etc.)

## Requisitos

* Go 1.17 o superior
* AWS SDK para Go
* Credenciales de AWS con permisos limitados
* Una cuenta de AWS válida

## Configuración

### 1. Configuración de Credenciales de AWS

Asegúrate de tener un archivo de credenciales de AWS configurado en tu máquina. Puedes hacerlo ejecutando el siguiente comando:

```sh
aws configure
```

Esto te pedirá que ingreses tu `AWS Access Key ID`, `AWS Secret Access Key`, región predeterminada y formato de salida. Si ya tienes configurado este archivo, puedes omitir este paso.

### 2. Configuración de la Región

Puedes configurar la región de AWS a través de una variable de entorno. Si no se especifica, se usará la región `us-west-2` por defecto.

```sh
export AWS_REGION=tu-region
```

### 3. Modo DryRun

Puedes habilitar el modo DryRun para probar el código sin realizar cambios reales en los recursos de AWS.

```sh
export DRY_RUN=true
```

## Instalación

1. Clona el repositorio:

   ```sh
   git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git
   ```

2. Navega al directorio del proyecto:

   ```sh
   cd Cloud-Instance-Manager
   ```

3. Instala las dependencias:

   ```sh
   go get -u github.com/aws/aws-sdk-go-v2/aws
   ```

## Uso

### 1. Ejecutar el Programa

Para ejecutar el programa, simplemente ejecuta el siguiente comando:

```sh
go run main.go
```

### 2. Ejemplos de Comandos

#### Crear una Instancia

El programa creará una instancia con la AMI `ami-12345678` por defecto. Puedes cambiar la AMI en el código si lo deseas.

```go
instanceID, err := crearInstancia("ami-12345678")
if err != nil {
    log.Fatalf("Error al crear instancia: %v", err)
}
fmt.Printf("Instancia creada: %s\n", instanceID)
```

#### Listar Instancias

El programa listará todas las instancias en la región especificada.

```go
instances, err := listarInstancias()
if err != nil {
    log.Fatalf("Error al listar instancias: %v", err)
}
for _, instance := range instances {
    fmt.Printf("ID: %s, Estado: %s\n", *instance.InstanceId, instance.State.Name)
}
```

#### Iniciar una Instancia

El programa iniciará la instancia creada anteriormente.

```go
if err := iniciarInstancia(instanceID); err != nil {
    log.Fatalf("Error al iniciar instancia: %v", err)
}
fmt.Println("Instancia iniciada")
```

#### Detener una Instancia

El programa detendrá la instancia creada anteriormente.

```go
if err := detenerInstancia(instanceID); err != nil {
    log.Fatalf("Error al detener instancia: %v", err)
}
fmt.Println("Instancia detenida")
```

#### Terminar una Instancia

El programa terminará la instancia creada anteriormente.

```go
if err := terminarInstancia(instanceID); err != nil {
    log.Fatalf("Error al terminar instancia: %v", err)
}
fmt.Println("Instancia terminada")
```

## Pruebas

### Modo DryRun

Para probar el código sin realizar cambios reales en los recursos de AWS, puedes habilitar el modo DryRun:

```sh
export DRY_RUN=true
go run main.go
```

### Pruebas Unitarias

Para ejecutar las pruebas unitarias, utiliza el siguiente comando:

```sh
go test -v
```

### Pruebas de Integración

Para ejecutar las pruebas de integración, utiliza el siguiente comando:

```sh
go test -v -tags=integration
```

## Licencia

Este proyecto está licenciado bajo la licencia MIT. Esto significa que puedes usar, modificar y distribuir el código de manera libre, siempre y cuando se incluya la licencia original y se respeten los términos y condiciones de la misma.

## Contribuciones

Si deseas contribuir a este proyecto, por favor, sigue los siguientes pasos:

1. Clona el repositorio:

   ```sh
   git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git
   ```

2. Crea una rama nueva:

   ```sh
   git branch my-branch
   ```

3. Realiza los cambios que desees

4. Envía un pull request con tus cambios

## Problemas y Errores

Si encuentras algún problema o error en el código, por favor, abre un issue en el repositorio de GitHub. Esto nos ayudará a identificar y solucionar el problema de manera más rápida.

---

