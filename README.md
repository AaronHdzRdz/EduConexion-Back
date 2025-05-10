# EduConexion - Backend en Go con GORM y Docker

## Descripción

Backend básico implementado con Go, usando el ORM GORM y una base de datos PostgreSQL, todo gestionado a través de Docker y Docker Compose.

## Requerimientos

* [Docker](https://www.docker.com/) instalado y configurado.
* [Docker Compose](https://docs.docker.com/compose/) instalado.
* Go (opcional, para desarrollo local sin Docker).

## Servicios Disponibles

### Base de datos

* PostgreSQL
* Puerto: `5432`

### Aplicación Go

* Framework: Gin
* Puerto: `3000`

## Endpoints disponibles

* `GET /ping`: Prueba de conexión.
* `POST /login`: Autenticación.
* CRUD usuarios:

  * `POST /users`: Crear usuario.
  * `GET /users`: Obtener usuarios.
  * `GET /users/:id`: Obtener usuario por ID.
  * `PUT /users/:id`: Actualizar usuario.
  * `DELETE /users/:id`: Eliminar usuario.

## Ejecutar la aplicación

Construye las imágenes y ejecuta los contenedores con el siguiente comando:

```bash
docker-compose up --build -d
```

### Detener la aplicación

```bash
docker-compose down
```

## Verificar estado de contenedores

```bash
docker-compose ps
```

## Variables de entorno

Configuradas en `docker-compose.yml`, incluyendo datos de conexión a la base de datos.
