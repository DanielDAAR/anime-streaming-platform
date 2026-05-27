# 📊 Diagramas del Sistema

## Diagrama de Casos de Uso

```mermaid
graph TB
    subgraph Usuarios
        Guest[Usuario Invitado]
        User[Usuario Registrado]
        Admin[Administrador]
    end

    subgraph Funcionalidades
        F1[Ver Catálogo]
        F2[Buscar Anime]
        F3[Ver Episodios]
        F4[Registrarse]
        F5[Iniciar Sesión]
        F6[Comentar]
        F7[Dar Like]
        F8[Responder]
        F9[CRUD Animes]
        F10[CRUD Episodios]
        F11[Moderar Comentarios]
        F12[Gestionar Usuarios]
    end

    Guest --> F1
    Guest --> F2
    Guest --> F3
    Guest --> F4
    Guest --> F5

    User --> F1
    User --> F2
    User --> F3
    User --> F6
    User --> F7
    User --> F8

    Admin --> F9
    Admin --> F10
    Admin --> F11
    Admin --> F12
```

## Diagrama de Flujo - Autenticación

```mermaid
sequenceDiagram
    participant U as Usuario
    participant F as Frontend
    participant B as Backend
    participant DB as MongoDB

    U->>F: Ingresa credenciales
    F->>B: POST /api/auth/login
    B->>DB: Buscar usuario por email
    DB-->>B: Usuario + passwordHash
    B->>B: bcrypt.Compare()
    alt Credenciales válidas
        B->>B: Generar JWT
        B-->>F: {token, user}
        F->>F: Guardar en localStorage
        F-->>U: Redirigir a home
    else Credenciales inválidas
        B-->>F: 401 Unauthorized
        F-->>U: Mostrar error
    end
```

## Diagrama de Flujo - Crear Anime

```mermaid
sequenceDiagram
    participant A as Admin
    participant F as Frontend
    participant B as Backend
    participant V as Validators
    participant S as Service
    participant R as Repository
    participant DB as MongoDB

    A->>F: Completa formulario
    F->>B: POST /api/animes + JWT
    B->>B: AuthMiddleware
    B->>B: AdminMiddleware
    B->>V: ValidateAnime()
    V-->>B: OK / Error
    B->>S: CreateAnime()
    S->>R: Create()
    R->>DB: InsertOne()
    DB-->>R: ObjectID
    R-->>S: nil
    S-->>B: AnimeResponse
    B-->>F: 201 Created
    F-->>A: Mostrar éxito
```

## Diagrama NoSQL - Modelo de Datos

```mermaid
erDiagram
    ANIME {
        ObjectId _id
        string slug UK
        string title
        string description
        array genres
        float rating
        object images
        string status
        int episodesCount
        int year
        string studio
        object seo
        date createdAt
        date updatedAt
    }

    EPISODE {
        ObjectId _id
        ObjectId animeId FK
        int number
        string title
        string description
        array servers
        int duration
        string thumbnail
        date createdAt
        date updatedAt
    }

    USER {
        ObjectId _id
        string username UK
        string email UK
        string passwordHash
        string role
        string avatar
        bool isActive
        date lastLogin
        date createdAt
        date updatedAt
    }

    COMMENT {
        ObjectId _id
        ObjectId animeId FK
        ObjectId userId FK
        ObjectId parentId FK
        string content
        int likes
        array likedBy
        bool isDeleted
        date createdAt
        date updatedAt
    }

    HISTORY {
        ObjectId _id
        ObjectId userId FK
        ObjectId animeId FK
        ObjectId episodeId FK
        int progress
        bool completed
        date createdAt
        date updatedAt
    }

    ANIME ||--o{ EPISODE : "tiene"
    ANIME ||--o{ COMMENT : "recibe"
    USER ||--o{ COMMENT : "escribe"
    USER ||--o{ HISTORY : "registra"
    COMMENT ||--o{ COMMENT : "responde"
```

## Diagrama de Arquitectura General

```mermaid
graph TB
    subgraph Client
        Browser[Navegador]
    end

    subgraph Frontend_Server["Frontend Server (Node.js)"]
        Astro[Astro Framework]
        SSR[SSR Engine]
        Islands[Astro Islands]
        Bulma[Bulma CSS]
    end

    subgraph Backend_Server["Backend Server (Go)"]
        Gin[Gin Router]
        Middleware[Middleware Layer]
        Controllers[Controllers]
        Services[Services]
        Repositories[Repositories]
    end

    subgraph Database["MongoDB"]
        Collections[(Collections)]
        Indexes[(Indexes)]
    end

    Browser -->|HTTP| Astro
    Astro --> SSR
    Astro --> Islands
    Astro --> Bulma
    Islands -->|Fetch API| Gin
    SSR -->|Server Fetch| Gin
    Gin --> Middleware
    Middleware --> Controllers
    Controllers --> Services
    Services --> Repositories
    Repositories --> Collections
    Collections --> Indexes
```

## Flujo Backend Detallado

```mermaid
flowchart TD
    A[Request HTTP] --> B{Router Gin}
    B --> C[Middleware]
    C --> D{Auth?}
    D -->|Sí| E[Validate JWT]
    D -->|No| F[Continue]
    E --> G{Admin?}
    G -->|Sí| H[Check Role]
    G -->|No| F
    H -->|Admin| I[Continue]
    H -->|No Admin| J[403 Forbidden]
    F --> K[Controller]
    I --> K
    K --> L[Validate Input]
    L -->|Inválido| M[400 Bad Request]
    L -->|Válido| N[Service Layer]
    N --> O[Business Logic]
    O --> P[Repository]
    P --> Q[MongoDB Query]
    Q --> R[Result]
    R --> S[Response JSON]
    S --> T[Client]
    M --> T
    J --> T
```

## Wireframes - Páginas Principales

### Home
```
┌─────────────────────────────────────┐
│  🎬 AnimeStream    [Inicio][Catálogo][Buscar] [Login] │
├─────────────────────────────────────┤
│                                     │
│     Descubre tu próximo anime       │
│     [Explorar Catálogo] [Buscar]    │
│                                     │
├─────────────────────────────────────┤
│  📺 Últimos Agregados    [Ver Todo] │
│  ┌────┐ ┌────┐ ┌────┐ ┌────┐       │
│  │Img │ │Img │ │Img │ │Img │       │
│  │Tit │ │Tit │ │Tit │ │Tit │       │
│  └────┘ └────┘ └────┘ └────┘       │
├─────────────────────────────────────┤
│  ⭐ Mejor Valorados                 │
│  ┌────┐ ┌────┐ ┌────┐ ┌────┐       │
│  │Img │ │Img │ │Img │ │Img │       │
│  │Tit │ │Tit │ │Tit │ │Tit │       │
│  └────┘ └────┘ └────┘ └────┘       │
├─────────────────────────────────────┤
│  Footer                             │
└─────────────────────────────────────┘
```

### Anime Detail
```
┌─────────────────────────────────────┐
│  Navbar...                          │
├─────────────────────────────────────┤
│  ┌────────┐  Título del Anime       │
│  │        │  [En Emisión] [⭐9.1]   │
│  │ Poster │  Descripción...         │
│  │        │  [Acción] [Aventura]    │
│  └────────┘                         │
├─────────────────────────────────────┤
│  📺 Episodios                       │
│  ┌─────────────────────────────┐    │
│  │ #1  Título Episodio    [▶]  │    │
│  │ #2  Título Episodio    [▶]  │    │
│  └─────────────────────────────┘    │
├─────────────────────────────────────┤
│  💬 Comentarios                     │
│  [Escribe tu comentario...] [Enviar]│
│  ┌─────────────────────────────┐    │
│  │ 👤 Usuario - fecha          │    │
│  │ Contenido del comentario    │    │
│  │ [👍5] [💬 Responder]        │    │
│  └─────────────────────────────┘    │
└─────────────────────────────────────┘
```

### Admin Dashboard
```
┌─────────────────────────────────────┐
│ [Sidebar]     │  Dashboard          │
│               │  ┌──┐ ┌──┐ ┌──┐ ┌──┐│
│ 📊 Dashboard  │  │50│ │5 │ │45│ │10││
│ 🎬 Animes     │  │Us│ │Ad│ │Us│ │An││
│ 📺 Episodios  │  └──┘ └──┘ └──┘ └──┘│
│ 💬 Comentarios│                     │
│ 👥 Usuarios   │  📺 Recientes       │
│               │  ┌─────────────────┐│
│ ⚙️ Sistema    │  │ Tabla de animes ││
│ 🏠 Ver Sitio  │  └─────────────────┘│
│ 🚪 Logout     │                     │
└─────────────────────────────────────┘
```
