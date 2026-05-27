// API Response Types
export interface ApiResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
  meta?: PaginationMeta;
}

export interface PaginationMeta {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

// User Types
export type UserRole = 'guest' | 'user' | 'admin';

export interface User {
  id: string;
  username: string;
  email: string;
  role: UserRole;
  avatar?: string;
  isActive: boolean;
  lastLogin?: string;
  createdAt: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

// Anime Types
export interface AnimeImages {
  poster: string;
  banner?: string;
  thumbnail?: string;
}

export interface SEOMetadata {
  metaTitle: string;
  metaDescription: string;
  keywords: string;
  canonicalUrl: string;
}

export interface Anime {
  id: string;
  slug: string;
  title: string;
  description: string;
  genres: string[];
  rating: number;
  images: AnimeImages;
  status: 'ongoing' | 'completed' | 'upcoming' | 'cancelled';
  episodesCount: number;
  year: number;
  studio?: string;
  seo: SEOMetadata;
  createdAt: string;
}

// Episode Types
export interface EmbedServer {
  name: string;
  url: string;
  quality: '360p' | '480p' | '720p' | '1080p' | 'unknown';
  active: boolean;
}

export interface Episode {
  id: string;
  animeId: string;
  number: number;
  title: string;
  description?: string;
  servers: EmbedServer[];
  duration?: number;
  thumbnail?: string;
  createdAt: string;
}

// Comment Types
export interface PublicUser {
  id: string;
  username: string;
  avatar?: string;
  role: string;
}

export interface Comment {
  id: string;
  animeId: string;
  userId: string;
  parentId?: string;
  content: string;
  likes: number;
  isDeleted: boolean;
  createdAt: string;
  user?: PublicUser;
  replies?: Comment[];
}

// History Types
export interface History {
  id: string;
  userId: string;
  animeId: string;
  episodeId: string;
  progress: number;
  completed: boolean;
  updatedAt: string;
  anime?: Anime;
  episode?: Episode;
}

// Dashboard Types
export interface DashboardStats {
  latestAnimes: Anime[];
  userStats: {
    totalAdmins: number;
    totalUsers: number;
    total: number;
  };
}

// Filter Types
export interface AnimeFilters {
  search?: string;
  genre?: string;
  status?: string;
  year?: number;
  page?: number;
  limit?: number;
}
