# 🏗 Arquitectura del Sistema

## Visión General

AnimeStream utiliza una arquitectura **MVC desacoplada** con separación clara de responsabilidades:

```
┌─────────────────────────────────────────────────────────────┐
│                      CLIENTE (Navegador)                     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Páginas   │  │   Islands   │  │  Client-side Auth   │  │
│  │    SSR      │  │    CSR      │  │      (JWT)          │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└──────────────────────────┬──────────────────────────────────┘
                           │ HTTP/REST
┌──────────────────────────▼──────────────────────────────────┐
│                    FRONTEND (Astro)                          │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Layouts   │  │ Components  │  │   API Services      │  │
│  │   Pages     │  │   Astro     │  │   TypeScript        │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└──────────────────────────┬──────────────────────────────────┘
                           │ Fetch API
┌──────────────────────────▼──────────────────────────────────┐
│                    BACKEND (Go/Gin)                          │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Controllers │  │  Services   │  │   Repositories      │  │
│  │   (HTTP)    │──│  (Logic)    │──│    (MongoDB)        │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Middleware  │  │ Validators  │  │      Utils          │  │
│  │  Auth/CORS  │  │   Input     │  │   JWT/Password      │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└──────────────────────────┬──────────────────────────────────┘
                           │ MongoDB Protocol
┌──────────────────────────▼──────────────────────────────────┐
│                   DATABASE (MongoDB)                         │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐        │
│  │ animes  │  │episodes │  │  users  │  │comments │        │
│  └─────────┘  └─────────┘  └─────────┘  └─────────┘        │
│  ┌─────────┐                                                │
│  │ history │                                                │
│  └─────────┘                                                │
└─────────────────────────────────────────────────────────────┘
```

## Flujo de Datos

### 1. Solicitud de Página Pública (SSR)
```
Usuario → Astro (SSR) → API Go → MongoDB → HTML Completo → Usuario
```

### 2. Interacción Dinámica (Islands)
```
Usuario → Island (CSR) → API Go → MongoDB → JSON → DOM Update
```

### 3. Operación Admin (CSR)
```
Admin → Dashboard → API Go + JWT → MongoDB → Response
```

## Capas del Backend

### Controllers
- Reciben requests HTTP
- Validan parámetros de ruta/query
- Llaman a Services
- Retornan responses JSON

### Services
- Contienen la lógica de negocio
- Coordinan operaciones entre repositorios
- Aplican reglas de negocio
- Transforman datos para responses

### Repositories
- Abstraen el acceso a MongoDB
- Manejan queries y aggregations
- Implementan paginación
- Gestionan índices

### Models
- Definen estructuras de datos
- Incluyen tags BSON/JSON
- Implementan métodos de transformación

## Patrones de Diseño

1. **Repository Pattern**: Aislamiento de la capa de datos
2. **Dependency Injection**: Servicios reciben repositorios
3. **Middleware Chain**: Autenticación, CORS, seguridad
4. **DTO Pattern**: Separación entre modelos internos y responses
5. **Singleton**: Conexión a base de datos

## Seguridad

### Autenticación
```
Login → bcrypt compare → JWT generate → Client storage
Request → JWT validate → Context set → Handler execution
```

### Autorización
```
Admin middleware → Role check → Continue/403
```

## Escalabilidad

- **Stateless**: Backend sin estado, escalable horizontalmente
- **Índices MongoDB**: Optimización de queries frecuentes
- **Paginación**: Todos los listados paginados
- **Caché**: Preparado para implementar Redis
