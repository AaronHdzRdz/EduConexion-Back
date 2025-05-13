# EduConexión

Repositorio monorepo del proyecto EduConexión, consistente en:

- **Back**: API RESTful en Go con Gin y Gorm
- **Front**: SPA en React + TypeScript + Tailwind CSS
- **Base de datos**: PostgreSQL

---

## 📦 Estructura

```
/Back         # Código del backend en Go
/Font         # Código del frontend en React/TS
/docker-compose.yml  # Despliegue local con Docker Compose
README.md     # Esta documentación
```

---

## 🚀 Levantar en local sin Docker

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/educonexion.git
   cd educonexion
   ```
2. Frontend:
   ```bash
   cd Font/educonexion-front
   npm install
   npm run dev
   ```
3. Backend:
   ```bash
   cd Back/EDUCONEXION/cmd
   go mod tidy
   go run main.go
   ```
4. Abre tu navegador en `http://localhost:5173` (frontend) y la API escucha en `http://localhost:3000`.

---

## 🐳 Despliegue con Docker

### 1. Descargar la imagen del backend (Go + API)
```bash
# Reemplaza "latest" por la etiqueta que corresponda si fuera necesario
docker pull sresendiz/educonexion-stack:latest
```

### 2. Levantar la base de datos y la API con Docker Compose
```bash
# Desde la raíz del proyecto, donde está docker-compose.yml
docker-compose up -d
```

Esto creará y levantará dos servicios:

- **database**: PostgreSQL en el puerto 5432
- **app**: tu API Go en el puerto 3000, conectada a la base de datos

#### Variables de entorno
El contenedor `app` leerá automáticamente del archivo `.env` (colocado junto a `docker-compose.yml`) o de las variables definidas en el `docker-compose.yml`:

```env
# .env (en la raíz junto a docker-compose.yml)
DB_HOST=database
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=294332
DB_NAME=EduConexion
JWT_SECRET=tu_secreto_jwt
```

### 3. Levantar el frontend en Docker
Si deseas también dockerizar el frontend, crea un `docker-compose.front.yml` (en `/Font/educonexion-front`) con:

```yaml
version: '3.8'
services:
  front:
    build:
      context: .
      dockerfile: Dockerfile
    image: educonexion-frontend:0.1
    container_name: educonexion-front
    restart: always
    ports:
      - "5173:80"
    networks:
      - educonexion_default
networks:
  educonexion_default:
    external: true
```

Y luego:
```bash
cd Font/educonexion-front
docker-compose -f docker-compose.front.yml up -d
```

Ahora el frontend estará accesible en `http://localhost:5173`.

---

## 📝 API Endpoints

- `POST /login`  → Autenticación, devuelve JWT + información de usuario
- Rutas protegidas bajo `/api` (requieren header `Authorization: Bearer <token>`):
  - `GET /api/students`
  - `POST /api/students`
  - `GET /api/students/:id`
  - `PUT /api/students/:id`
  - `DELETE /api/students/:id`

… y similares para `subjects`, `grades`, etc.

---

## 📖 Licencia
MIT © Tu Nombre