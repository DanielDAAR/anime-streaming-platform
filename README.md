# AnimeStream - Plataforma de Streaming de Anime

Proyecto integrador académico-profesional desarrollado con arquitectura MVC desacoplada, SSR híbrido y API REST.

## Tabla de Contenidos

- [Descripción](#descripción)
- [Stack Tecnológico](#stack-tecnológico)
- [Arquitectura](#arquitectura)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Uso](#uso)
- [API Endpoints](#api-endpoints)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Base de Datos](#base-de-datos)
- [Seguridad](#seguridad)
- [Licencia](#licencia)

## Descripción

AnimeStream es una plataforma web administrativa para la gestión y distribución de contenido multimedia tipo streaming mediante enlaces embebidos. El sistema permite:

- Administrar animes y episodios
- Gestionar usuarios con roles
- Manejar comentarios con likes y respuestas
- Autenticación JWT
- Búsqueda con índices de texto MongoDB
- Panel administrativo completo
- SEO optimizado con SSR híbrido

## Stack Tecnológico

### Frontend
- **Astro** - Framework web con SSR híbrido
- **TypeScript** - Tipado estático
- **Bulma CSS** - Framework CSS

### Backend
- **Go** - Lenguaje de programación
- **Gin Framework** - Web framework
- **MongoDB Driver** - Conexión a base de datos

### Base de Datos
- **MongoDB** - Base de datos NoSQL

### Autenticación
- **JWT** - JSON Web Tokens
- **bcrypt** - Hash de contraseñas

## 🏗 Arquitectura

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Frontend      │     │    Backend      │     │   Database      │
│   (Astro)       │◄────┤    (Go/Gin)     │◄────┤   (MongoDB)     │
│   SSR/CSR       │     │   API REST      │     │   NoSQL         │
└─────────────────┘     └─────────────────┘     └─────────────────┘
        │                       │                       │
        └────── MVC Desacoplado ┘                       │
                Separación por Capas                    │
                Servicios & Repositorios                │
```

### Patrones utilizados:
- **MVC Desacoplado**: Separación clara de responsabilidades
- **Repository Pattern**: Abstracción de acceso a datos
- **Service Layer**: Lógica de negocio independiente
- **Middleware**: Autenticación, CORS, seguridad

## Instalación

### Prerrequisitos
- Go 1.21+
- Node.js 18+
- MongoDB 6.0+
- Git

### 1. Clonar el repositorio
```bash
git clone <repository-url>
cd anime-streaming-platform
```

### 2. Configurar Backend
```bash
cd backend

# Copiar variables de entorno
cp .env.example .env

# Instalar dependencias
go mod download

# Compilar
go build -o server
```

### 3. Configurar Frontend
```bash
cd frontend

# Instalar dependencias
npm install

# Compilar para producción
npm run build
```

### 4. Configurar Base de Datos
```bash
# Asegurar que MongoDB esté corriendo
mongod --dbpath /path/to/data

# Ejecutar seeders
cd database/seeders
npm install bcryptjs mongodb
node seed.js
```

## ⚙ Configuración

### Variables de Entorno Backend (.env)
```env
# Server Configuration
PORT=8080
GIN_MODE=release

# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB_NAME=anime_streaming_db

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRATION_HOURS=24

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:4321,http://localhost:8080
```

### Variables de Entorno Frontend (.env)
```env
PUBLIC_API_URL=http://localhost:8080/api
```

## ▶ Uso

### Desarrollo
```bash
# Terminal 1 - Backend
cd backend
go run main.go

# Terminal 2 - Frontend
cd frontend
npm run dev
```

### Producción
```bash
# Backend
cd backend
./server

# Frontend (modo standalone)
cd frontend
npm run build
node ./dist/server/entry.mjs
```

### Acceso
- **Sitio Web**: http://localhost:4321
- **API**: http://localhost:8080
- **Admin Panel**: http://localhost:4321/admin/dashboard

### Credenciales por defecto
- **Admin**: admin@animestream.com / admin123
- **Usuario**: user@animestream.com / user123

## 🔌 API Endpoints

### Autenticación
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | /api/auth/register | Registrar usuario |
| POST | /api/auth/login | Iniciar sesión |
| GET | /api/auth/me | Obtener usuario actual |

### Animes
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | /api/animes | Listar animes (paginado) |
| GET | /api/animes/:slug | Obtener anime por slug |
| POST | /api/animes | Crear anime (admin) |
| PUT | /api/animes/:id | Actualizar anime (admin) |
| DELETE | /api/animes/:id | Eliminar anime (admin) |
| GET | /api/animes/latest | Últimos animes |
| GET | /api/animes/top-rated | Mejor valorados |

### Episodios
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | /api/animes/:id/episodes | Episodios por anime |
| GET | /api/episodes/:id | Obtener episodio |
| POST | /api/episodes | Crear episodio (admin) |
| PUT | /api/episodes/:id | Actualizar episodio (admin) |
| DELETE | /api/episodes/:id | Eliminar episodio (admin) |

### Comentarios
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | /api/comments/:animeId | Comentarios por anime |
| POST | /api/comments | Crear comentario |
| POST | /api/comments/:id/reply | Responder comentario |
| POST | /api/comments/:id/like | Dar like |
| GET | /api/comments | Moderación (admin) |
| DELETE | /api/comments/:id | Eliminar (admin) |

### Usuarios
| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | /api/users | Listar usuarios (admin) |
| GET | /api/users/:id | Obtener usuario (admin) |
| PUT | /api/users/:id/role | Cambiar rol (admin) |
| PUT | /api/users/:id/toggle-active | Activar/Desactivar (admin) |

## Estructura del Proyecto

```
anime-streaming-platform/
├── backend/
│   ├── controllers/      # Controladores HTTP
│   ├── services/         # Lógica de negocio
│   ├── repositories/     # Acceso a datos
│   ├── models/           # Modelos de datos
│   ├── routes/           # Definición de rutas
│   ├── middleware/       # Middlewares
│   ├── config/           # Configuración
│   ├── utils/            # Utilidades
│   ├── validators/       # Validaciones
│   └── main.go           # Punto de entrada
├── frontend/
│   ├── src/
│   │   ├── pages/        # Páginas Astro
│   │   ├── layouts/      # Layouts
│   │   ├── components/   # Componentes
│   │   ├── services/     # Servicios API
│   │   ├── styles/       # Estilos SCSS
│   │   ├── middleware/   # Middlewares
│   │   └── types/        # Tipos TypeScript
│   ├── public/           # Archivos estáticos
│   └── astro.config.mjs  # Config Astro
├── database/
│   ├── migrations/       # Migraciones
│   └── seeders/          # Datos iniciales
└── docs/                 # Documentación
```

## 🗄 Base de Datos

### Colecciones
1. **animes** - Información de series
2. **episodes** - Episodios con servidores embebidos
3. **users** - Usuarios y autenticación
4. **comments** - Comentarios y respuestas
5. **history** - Historial de visualización

### Índices
- Animes: slug (único), texto (title, description, genres), rating, createdAt
- Episodes: animeId + number, createdAt
- Users: email (único), username (único), role
- Comments: animeId + createdAt, userId, parentId
- History: userId + animeId (único), updatedAt

## Seguridad

- **bcrypt**: Hash de contraseñas con salt automático
- **JWT**: Tokens con expiración configurable
- **CORS**: Orígenes permitidos configurables
- **Headers de seguridad**: X-Content-Type-Options, X-Frame-Options, etc.
- **Validación**: Sanitización de inputs en todos los endpoints
- **Middleware**: Separación de roles (user/admin)

## Licencia

Este proyecto es de uso académico. Desarrollado como proyecto integrador universitario.

---

**Desarrollado con ❤️ para fines educativos**
