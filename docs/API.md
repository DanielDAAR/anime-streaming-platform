# 📚 Documentación de la API

## Base URL
```
http://localhost:8080/api
```

## Autenticación

La API utiliza JWT Bearer tokens. Incluye el token en el header:
```
Authorization: Bearer <token>
```

## Respuestas

Todas las respuestas siguen el formato:
```json
{
  "success": true,
  "message": "Descripción",
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "totalPages": 5
  }
}
```

## Endpoints Detallados

### POST /auth/register
Registra un nuevo usuario.

**Body:**
```json
{
  "username": "nuevo_usuario",
  "email": "usuario@email.com",
  "password": "contraseña123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": "...",
      "username": "nuevo_usuario",
      "email": "usuario@email.com",
      "role": "user"
    }
  }
}
```

### POST /auth/login
Inicia sesión.

**Body:**
```json
{
  "email": "usuario@email.com",
  "password": "contraseña123"
}
```

### GET /animes
Lista animes con paginación y filtros.

**Query Parameters:**
- `page` (number): Página actual
- `limit` (number): Items por página
- `search` (string): Búsqueda por texto
- `genre` (string): Filtrar por género
- `status` (string): ongoing, completed, upcoming, cancelled
- `year` (number): Año de lanzamiento

### POST /animes
Crea un nuevo anime (requiere admin).

**Body:**
```json
{
  "title": "Nombre del Anime",
  "slug": "nombre-del-anime",
  "description": "Descripción detallada...",
  "genres": ["Acción", "Aventura"],
  "rating": 8.5,
  "images": {
    "poster": "https://...",
    "banner": "https://...",
    "thumbnail": "https://..."
  },
  "status": "ongoing",
  "year": 2024,
  "studio": "Studio Name"
}
```

### POST /episodes
Crea un nuevo episodio (requiere admin).

**Body:**
```json
{
  "animeId": "...",
  "number": 1,
  "title": "Título del Episodio",
  "description": "Descripción...",
  "servers": [
    {
      "name": "Server 1",
      "url": "https://embed.example.com/video",
      "quality": "1080p",
      "active": true
    }
  ],
  "duration": 24,
  "thumbnail": "https://..."
}
```

### POST /comments
Crea un comentario.

**Body:**
```json
{
  "animeId": "...",
  "content": "¡Gran episodio!"
}
```

### POST /comments/:id/reply
Responde a un comentario (máximo 1 nivel).

**Body:**
```json
{
  "content": "Totalmente de acuerdo"
}
```

### POST /comments/:id/like
Da like a un comentario.

## Códigos de Estado

| Código | Descripción |
|--------|-------------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

## Errores Comunes

```json
{
  "success": false,
  "error": "Mensaje de error descriptivo"
}
```
