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

# Cloud-Instance-Manager

This project is a minimalist cloud instance manager written in Go. It allows you to create, list, start, stop, and terminate cloud instances in a secure and efficient manner. The goal of this project is to provide a simple and easy-to-use tool for managing cloud instances on Amazon Web Services (AWS).

## Features

* Create cloud instances with a single API call
* List cloud instances in a specific region
* Start, stop, and terminate cloud instances securely
* DryRun mode to test the code without affecting real resources
* Support for low-cost cloud instances (t2.micro)
* Support for cloud instances with different instance types (t2.micro, t2.small, etc.)

## Requirements

* Go 1.17 or higher
* AWS SDK for Go
* AWS credentials with limited permissions
* A valid AWS account

## Setup

### 1. AWS Credentials Configuration

Ensure you have an AWS credentials file configured on your machine. You can do this by running the following command:

```sh
aws configure
```

This will prompt you to enter your `AWS Access Key ID`, `AWS Secret Access Key`, default region, and output format. If you already have this file configured, you can skip this step.

### 2. Region Configuration

You can configure the AWS region via an environment variable. If not specified, the default region `us-west-2` will be used.

```sh
export AWS_REGION=your-region
```

### 3. DryRun Mode

You can enable DryRun mode to test the code without making actual changes to AWS resources.

```sh
export DRY_RUN=true
```

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git
   ```

2. Navigate to the project directory:

   ```sh
   cd Cloud-Instance-Manager
   ```

3. Install the dependencies:

   ```sh
   go get -u github.com/aws/aws-sdk-go-v2/aws
   ```

## Usage

### 1. Run the Program

To run the program, simply execute the following command:

```sh
go run main.go
```

### 2. Command Examples

#### Create an Instance

The program will create an instance with the default AMI `ami-12345678`. You can change the AMI in the code if you wish.

```go
instanceID, err := crearInstancia("ami-12345678")
if err != nil {
    log.Fatalf("Error creating instance: %v", err)
}
fmt.Printf("Instance created: %s\n", instanceID)
```

#### List Instances

The program will list all instances in the specified region.

```go
instances, err := listarInstancias()
if err != nil {
    log.Fatalf("Error listing instances: %v", err)
}
for _, instance := range instances {
    fmt.Printf("ID: %s, State: %s\n", *instance.InstanceId, instance.State.Name)
}
```

#### Start an Instance

The program will start the previously created instance.

```go
if err := iniciarInstancia(instanceID); err != nil {
    log.Fatalf("Error starting instance: %v", err)
}
fmt.Println("Instance started")
```

#### Stop an Instance

The program will stop the previously created instance.

```go
if err := detenerInstancia(instanceID); err != nil {
    log.Fatalf("Error stopping instance: %v", err)
}
fmt.Println("Instance stopped")
```

#### Terminate an Instance

The program will terminate the previously created instance.

```go
if err := terminarInstancia(instanceID); err != nil {
    log.Fatalf("Error terminating instance: %v", err)
}
fmt.Println("Instance terminated")
```

## Testing

### DryRun Mode

To test the code without making actual changes to AWS resources, you can enable DryRun mode:

```sh
export DRY_RUN=true
go run main.go
```

### Unit Tests

To run the unit tests, use the following command:

```sh
go test -v
```

### Integration Tests

To run the integration tests, use the following command:

```sh
go test -v -tags=integration
```

## License

This project is licensed under the MIT License. This means you can use, modify, and distribute the code freely, as long as the original license is included and its terms and conditions are respected.

## Contributions

If you wish to contribute to this project, please follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git
   ```

2. Create a new branch:

   ```sh
   git branch my-branch
   ```

3. Make your changes

4. Submit a pull request with your changes

## Issues and Bugs

If you find any issues or bugs in the code, please open an issue on the GitHub repository. This will help us identify and fix the problem more quickly.