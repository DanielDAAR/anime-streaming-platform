```markdown
# AnimeStream - Plataforma de Streaming de Anime

Proyecto integrador academico-profesional desarrollado con arquitectura MVC desacoplada, SSR hibrido y API REST.

## Tabla de Contenidos

- [Descripcion](#descripcion)
- [Stack Tecnologico](#stack-tecnologico)
- [Arquitectura](#arquitectura)
- [Instalacion](#instalacion)
- [Configuracion](#configuracion)
- [Uso](#uso)
- [API Endpoints](#api-endpoints)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Base de Datos](#base-de-datos)
- [Seguridad](#seguridad)
- [Licencia](#licencia)

## Descripcion

AnimeStream es una plataforma web administrativa para la gestion y distribucion de contenido multimedia tipo streaming mediante enlaces embebidos. El sistema permite:

- Administrar animes y episodios
- Gestionar usuarios con roles
- Manejar comentarios con likes y respuestas
- Autenticacion JWT
- Busqueda con indices de texto MongoDB
- Panel administrativo completo
- SEO optimizado con SSR hibrido

## Stack Tecnologico

### Frontend
- **Astro** - Framework web con SSR hibrido
- **TypeScript** - Tipado estatico
- **Bulma CSS** - Framework CSS

### Backend
- **Go** - Lenguaje de programacion
- **Gin Framework** - Web framework
- **MongoDB Driver** - Conexion a base de datos

### Base de Datos
- **MongoDB** - Base de datos NoSQL

### Autenticacion
- **JWT** - JSON Web Tokens
- **bcrypt** - Hash de contrasenas

## Arquitectura

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Frontend      │     │    Backend      │     │   Database      │
│   (Astro)       │◄────┤    (Go/Gin)     │◄────┤   (MongoDB)     │
│   SSR/CSR       │     │   API REST      │     │   NoSQL         │
└─────────────────┘     └─────────────────┘     └─────────────────┘
        │                       │                       │
        └────── MVC Desacoplado ┘                       │
                Separacion por Capas                    │
                Servicios & Repositorios                │
```

### Patrones utilizados:
- **MVC Desacoplado**: Separacion clara de responsabilidades
- **Repository Pattern**: Abstraccion de acceso a datos
- **Service Layer**: Logica de negocio independiente
- **Middleware**: Autenticacion, CORS, seguridad

## Instalacion

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

# Compilar para produccion
npm run build
```

### 4. Configurar Base de Datos
```bash
# Asegurar que MongoDB este corriendo
mongod --dbpath /path/to/data

# Ejecutar seeders
cd database/seeders
npm install bcryptjs mongodb
node seed.js
```

## Configuracion

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

## Uso

### Desarrollo
```bash
# Terminal 1 - Backend
cd backend
go run main.go

# Terminal 2 - Frontend
cd frontend
npm run dev
```

### Produccion
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

## API Endpoints

### Autenticacion
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| POST | /api/auth/register | Registrar usuario |
| POST | /api/auth/login | Iniciar sesion |
| GET | /api/auth/me | Obtener usuario actual |

### Animes
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/animes | Listar animes (paginado) |
| GET | /api/animes/:slug | Obtener anime por slug |
| POST | /api/animes | Crear anime (admin) |
| PUT | /api/animes/:id | Actualizar anime (admin) |
| DELETE | /api/animes/:id | Eliminar anime (admin) |
| GET | /api/animes/latest | Ultimos animes |
| GET | /api/animes/top-rated | Mejor valorados |

### Episodios
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/animes/:id/episodes | Episodios por anime |
| GET | /api/episodes/:id | Obtener episodio |
| POST | /api/episodes | Crear episodio (admin) |
| PUT | /api/episodes/:id | Actualizar episodio (admin) |
| DELETE | /api/episodes/:id | Eliminar episodio (admin) |

### Comentarios
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/comments/:animeId | Comentarios por anime |
| POST | /api/comments | Crear comentario |
| POST | /api/comments/:id/reply | Responder comentario |
| POST | /api/comments/:id/like | Dar like |
| GET | /api/comments | Moderacion (admin) |
| DELETE | /api/comments/:id | Eliminar (admin) |

### Usuarios
| Metodo | Endpoint | Descripcion |
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
│   ├── services/         # Logica de negocio
│   ├── repositories/     # Acceso a datos
│   ├── models/           # Modelos de datos
│   ├── routes/           # Definicion de rutas
│   ├── middleware/       # Middlewares
│   ├── config/           # Configuracion
│   ├── utils/            # Utilidades
│   ├── validators/       # Validaciones
│   └── main.go           # Punto de entrada
├── frontend/
│   ├── src/
│   │   ├── pages/        # Paginas Astro
│   │   ├── layouts/      # Layouts
│   │   ├── components/   # Componentes
│   │   ├── services/     # Servicios API
│   │   ├── styles/       # Estilos SCSS
│   │   ├── middleware/   # Middlewares
│   │   └── types/        # Tipos TypeScript
│   ├── public/           # Archivos estaticos
│   └── astro.config.mjs  # Config Astro
├── database/
│   ├── migrations/       # Migraciones
│   └── seeders/          # Datos iniciales
└── docs/                 # Documentacion
```

## Base de Datos

### Colecciones
1. **animes** - Informacion de series
2. **episodes** - Episodios con servidores embebidos
3. **users** - Usuarios y autenticacion
4. **comments** - Comentarios y respuestas
5. **history** - Historial de visualizacion

### Indices
- Animes: slug (unico), texto (title, description, genres), rating, createdAt
- Episodes: animeId + number, createdAt
- Users: email (unico), username (unico), role
- Comments: animeId + createdAt, userId, parentId
- History: userId + animeId (unico), updatedAt

## Seguridad

- **bcrypt**: Hash de contrasenas con salt automatico
- **JWT**: Tokens con expiracion configurable
- **CORS**: Origenes permitidos configurables
- **Headers de seguridad**: X-Content-Type-Options, X-Frame-Options, etc.
- **Validacion**: Sanitizacion de inputs en todos los endpoints
- **Middleware**: Separacion de roles (user/admin)

## Licencia

Este proyecto es de uso academico. Desarrollado como proyecto integrador universitario.

---

Desarrollado para fines educativos por DCSW - Daniel and Cesar Softworks
```
