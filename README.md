**Cloud-Instance-Manager**

Este proyecto es un gestor de instancias en la nube minimalista escrito en Go. Permite crear, listar, iniciar, detener y terminar instancias en la nube de manera segura y eficiente. El objetivo de este proyecto es proporcionar una herramienta simple y fácil de usar para administrar instancias en la nube de Amazon Web Services (AWS).

**Características**

* Crear instancias en la nube con una sola llamada a la API
* Listar instancias en la nube en una región específica
* Iniciar, detener y terminar instancias en la nube de manera segura
* Modo DryRun para probar el código sin afectar recursos reales
* Soporte para instancias en la nube de bajo costo (t2.micro)
* Soporte para instancias en la nube con diferentes tipos de instancias (t2.micro, t2.small, etc.)

**Requisitos**

* Go 1.17 o superior
* AWS SDK para Go
* Credenciales de AWS con permisos limitados
* Una cuenta de AWS válida

**Uso**

1. Clona el repositorio: `git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git`
2. Instala las dependencias: `go get -u github.com/aws/aws-sdk-go-v2/aws`
3. Configura tus credenciales de AWS: `export AWS_ACCESS_KEY_ID=tu-llave-de-acceso` y `export AWS_SECRET_ACCESS_KEY=tu-llave-secreta`
4. Ejecuta el programa: `go run main.go`

**Pruebas**

* Modo DryRun: `export DRY_RUN=true` y ejecuta el programa
* Pruebas unitarias: `go test -v`
* Pruebas de integración: `go test -v -tags=integration`

**Licencia**

Este proyecto está licenciado bajo la licencia MIT. Esto significa que puedes usar, modificar y distribuir el código de manera libre, siempre y cuando se incluya la licencia original y se respeten los términos y condiciones de la misma.

**Contribuciones**

Si deseas contribuir a este proyecto, por favor, sigue los siguientes pasos:

1. Clona el repositorio: `git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git`
2. Crea una rama nueva: `git branch mi-rama`
3. Realiza los cambios que desees
4. Envía un pull request con tus cambios

**Problemas y errores**

Si encuentras algún problema o error en el código, por favor, abre un issue en el repositorio de GitHub. Esto nos ayudará a identificar y solucionar el problema de manera más rápida.

---

**Cloud-Instance-Manager**

This project is a minimalistic cloud instance manager written in Go. It allows you to securely and efficiently create, list, start, stop, and terminate cloud instances. The goal of this project is to provide a simple and easy-to-use tool for managing AWS cloud instances.

**Features**

* Create cloud instances with a single API call
* List cloud instances in a specific region
* Securely start, stop, and terminate cloud instances
* DryRun mode to test code without affecting real resources
* Support for low-cost cloud instances (t2.micro)
* Support for different instance types (t2.micro, t2.small, etc.)

**Requirements**

* Go 1.17 or later
* AWS SDK for Go
* Limited AWS credentials with restricted permissions
* A valid AWS account

**Usage**

1. Clone the repository: `git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git`
2. Install dependencies: `go get -u github.com/aws/aws-sdk-go-v2/aws`
3. Configure your AWS credentials: `export AWS_ACCESS_KEY_ID=your-access-key` and `export AWS_SECRET_ACCESS_KEY=your-secret-key`
4. Run the program: `go run main.go`

**Testing**

* DryRun mode: `export DRY_RUN=true` and run the program
* Unit tests: `go test -v`
* Integration tests: `go test -v -tags=integration`

**License**

This project is licensed under the MIT license. This means you can use, modify, and distribute the code freely, as long as you include the original license and respect its terms and conditions.

**Contributions**

If you want to contribute to this project, please follow these steps:

1. Clone the repository: `git clone https://github.com/elliotsecops/Cloud-Instance-Manager.git`
2. Create a new branch: `git branch my-branch`
3. Make your desired changes
4. Submit a pull request with your changes

**Issues and Bugs**

If you find any issues or bugs in the code, please open an issue on the GitHub repository. This will help us identify and resolve the problem more quickly.
