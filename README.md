# 🎓 Pardalis Backend

Bienvenido al backend de Pardalis, el corazón que impulsa nuestra plataforma educativa de inglés. Este proyecto está construido con Go, pensando en la simplicidad, el rendimiento y la escalabilidad.

## 🌟 Descripción

Pardalis Backend es el motor que hace posible que los estudiantes aprendan inglés de forma divertida y efectiva. Gestionamos todo, desde la autenticación hasta el seguimiento del progreso del aprendizaje, de manera segura y eficiente.

## 🚀 Características Principales

- Autenticación segura con JWT
- API RESTful intuitiva
- Sistema de control de progreso
- Gestión de contenido educativo
- Limitación de tasa de peticiones
- Manejo robusto de errores

## 🛠️ Tecnologías Principales

- **Go**: Nuestro lenguaje principal
- **MySQL**: Base de datos relacional
- **JWT**: Para autenticación segura
- **Gorilla Mux**: Router HTTP
- **GORM**: ORM para Go

## 🏁 Primeros Pasos

### Requisitos Previos

- Go 1.23 o superior
- MySQL 8.0+
- Make (opcional, pero recomendado)

### Configuración del Entorno

1. Clone el repositorio:
```bash
git clone https://gitlab.com/pardalis/pardalis-api.git
cd pardalis-api
```

2. Configure las variables de entorno:
```bash
cp .env.example .env
```

3. Configure su archivo .env:
```env
PORT=8080
PUBLIC_HOST=http://localhost

DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=pardalis_db

JWT_SECRET=your_secret_key
```

4. Inicie el servidor:
```bash
go run main.go
```

## 🔄 Endpoints API Principales

### Autenticación
- `POST /api/v1/login`: Inicio de sesión
- `POST /api/v1/register`: Registro de usuario

### Usuarios
- `GET /api/v1/users/{userApodo}`: Obtener perfil de usuario

## 🧪 Pruebas

Ejecute las pruebas con:
```bash
go test ./...
```

## 🤝 Contribución

Nos encanta recibir contribuciones. Por favor, lea nuestra guía de contribución (CONTRIBUTING.md) antes de enviar un pull request.

### Pasos para Contribuir

1. Fork del repositorio
2. Cree una rama para su funcionalidad
3. Realice sus cambios
4. Envíe un Pull Request

## 📋 Directrices de Código

- Use nombres descriptivos
- Documente sus funciones
- Mantenga las funciones pequeñas y enfocadas
- Escriba pruebas para su código
- Siga las convenciones de Go

## 🔐 Seguridad

Si descubre algún problema de seguridad, por favor repórtelo a través de un issue privado.

## 📜 Licencia

Este proyecto está bajo la Licencia GPL v3. Vea el archivo LICENSE para más detalles.

## 🐛 Reporte de Problemas

Si encuentra algún problema, por favor repórtelo usando nuestro sistema de issues. Incluya:

- Descripción del problema
- Pasos para reproducirlo
- Comportamiento esperado
- Capturas de pantalla (si aplica)
- Entorno (SO, versión de Go, etc.)

---

🍪 Si has leído hasta aquí, te has ganado una galleta virtual. ¡Gracias por tu interés en Pardalis!