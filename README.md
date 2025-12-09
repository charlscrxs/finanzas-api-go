# API REST de Finanzas Personales

## Stack Tecnológico
* Lenguaje: Go (Golang)
* Framework Web: Gin Gonic
* Base de Datos: PostgreSQL
* ORM: GORM
* Autenticación: JSON Web Tokens (JWT)
* Arquitectura: MVC Modular

*## Funcionalidades Principales
* **Seguridad de Usuarios:** Sistema de registro y acceso seguro (Login) para proteger la información privada de cada cuenta.
* **Control Financiero:** Herramientas completas para registrar, consultar, modificar y eliminar tanto ingresos como gastos.
* **Organización de Gastos:** Clasificación automática de las transacciones por categorías para entender mejor en qué se gasta el dinero.

## Instrucciones de Ejecución

1. Clonar el repositorio:
   git clone https://github.com/charlscrxs/finanzas-api-go.git

2. Instalar dependencias:
   go mod tidy

3. Configuración:
   Crear un archivo .env en la raíz del proyecto configurando las credenciales de la base de datos PostgreSQL (Host, User, Password, DB_Name, Port).

4. Iniciar el servidor:
   go run cmd/main.go
