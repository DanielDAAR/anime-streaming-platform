export const API_BASE_URL = import.meta.env.PUBLIC_API_URL || 'http://localhost:8080/api';

export function apiUrl(endpoint: string): string {
  return `${API_BASE_URL}${endpoint}`;
}

export async function fetchApi<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const url = apiUrl(endpoint);

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...((options.headers as Record<string, string>) || {}),
  };

  // Add auth token if available
  if (typeof window !== 'undefined') {
    const token = localStorage.getItem('token');
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
  }

  const response = await fetch(url, {
    ...options,
    headers,
  });

  const data = response.status === 204 ? null : await response.json();

  if (!response.ok) {
    throw new ApiError(response.status, data?.error || 'An error occurred');
  }

  return data;
}

class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = 'ApiError';
  }
}

// Auth API
export const authApi = {
  login: (email: string, password: string) =>
    fetchApi<{ data: import('../types').AuthResponse }>('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),

  register: (username: string, email: string, password: string) =>
    fetchApi<{ data: import('../types').AuthResponse }>('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ username, email, password }),
    }),

  me: () =>
    fetchApi<{ data: import('../types').User }>('/auth/me'),
};

// Anime API
export const animeApi = {
  getAll: (params?: Record<string, string>) => {
    const query = params ? '?' + new URLSearchParams(params).toString() : '';
    return fetchApi<{ data: import('../types').Anime[]; meta: import('../types').PaginationMeta }>(`/animes${query}`);
  },

  getBySlug: (slug: string) =>
    fetchApi<{ data: import('../types').Anime }>(`/animes/${slug}`),

  getLatest: (limit = '10') =>
    fetchApi<{ data: import('../types').Anime[] }>(`/animes/latest?limit=${limit}`),

  getTopRated: (limit = '10') =>
    fetchApi<{ data: import('../types').Anime[] }>(`/animes/top-rated?limit=${limit}`),

  create: (anime: Partial<import('../types').Anime>) =>
    fetchApi<{ data: import('../types').Anime }>('/animes', {
      method: 'POST',
      body: JSON.stringify(anime),
    }),

  update: (id: string, anime: Partial<import('../types').Anime>) =>
    fetchApi<{ data: import('../types').Anime }>(`/animes/${id}`, {
      method: 'PUT',
      body: JSON.stringify(anime),
    }),

  delete: (id: string) =>
    fetchApi<{}>(`/animes/${id}`, { method: 'DELETE' }),
};

// Episode API
export const episodeApi = {
  getByAnime: (animeId: string, params?: Record<string, string>) => {
    const query = params ? '?' + new URLSearchParams(params).toString() : '';
    return fetchApi<{ data: import('../types').Episode[]; meta: import('../types').PaginationMeta }>(`/animes/${animeId}/episodes${query}`);
  },

  getById: (id: string) =>
    fetchApi<{ data: import('../types').Episode }>(`/episodes/${id}`),

  getByAnimeAndNumber: (animeRef: string, number: string | number) =>
    fetchApi<{ data: import('../types').Episode }>(`/animes/${animeRef}/episodes/${number}`),

  getLatest: (limit = '10') =>
    fetchApi<{ data: import('../types').Episode[] }>(`/episodes/latest?limit=${limit}`),

  create: (episode: Partial<import('../types').Episode>) =>
    fetchApi<{ data: import('../types').Episode }>('/episodes', {
      method: 'POST',
      body: JSON.stringify(episode),
    }),

  update: (id: string, episode: Partial<import('../types').Episode>) =>
    fetchApi<{ data: import('../types').Episode }>(`/episodes/${id}`, {
      method: 'PUT',
      body: JSON.stringify(episode),
    }),

  delete: (id: string) =>
    fetchApi<{}>(`/episodes/${id}`, { method: 'DELETE' }),
};

// Comment API
export const commentApi = {
  getByAnime: (animeId: string, params?: Record<string, string>) => {
    const query = params ? '?' + new URLSearchParams(params).toString() : '';
    return fetchApi<{ data: import('../types').Comment[]; meta: import('../types').PaginationMeta }>(`/comments/${animeId}${query}`);
  },

  create: (animeId: string, content: string) =>
    fetchApi<{ data: import('../types').Comment }>('/comments', {
      method: 'POST',
      body: JSON.stringify({ animeId, content }),
    }),

  reply: (commentId: string, content: string) =>
    fetchApi<{ data: import('../types').Comment }>(`/comments/${commentId}/reply`, {
      method: 'POST',
      body: JSON.stringify({ content }),
    }),

  like: (commentId: string) =>
    fetchApi<{}>(`/comments/${commentId}/like`, { method: 'POST' }),

  getRecent: (params?: Record<string, string>) => {
    const query = params ? '?' + new URLSearchParams(params).toString() : '';
    return fetchApi<{ data: import('../types').Comment[]; meta: import('../types').PaginationMeta }>(`/comments${query}`);
  },

  delete: (id: string) =>
    fetchApi<{}>(`/comments/${id}`, { method: 'DELETE' }),
};

// User API
export const userApi = {
  getAll: (params?: Record<string, string>) => {
    const query = params ? '?' + new URLSearchParams(params).toString() : '';
    return fetchApi<{ data: import('../types').User[]; meta: import('../types').PaginationMeta }>(`/users${query}`);
  },

  getById: (id: string) =>
    fetchApi<{ data: import('../types').User }>(`/users/${id}`),

  updateRole: (id: string, role: string) =>
    fetchApi<{}>(`/users/${id}/role`, {
      method: 'PUT',
      body: JSON.stringify({ role }),
    }),

  toggleActive: (id: string) =>
    fetchApi<{}>(`/users/${id}/toggle-active`, { method: 'PUT' }),

  getStats: () =>
    fetchApi<{ data: { totalAdmins: number; totalUsers: number; total: number } }>('/users/stats'),
};

// Dashboard API
export const dashboardApi = {
  getStats: () =>
    fetchApi<{ data: import('../types').DashboardStats }>('/dashboard/stats'),
};

export { ApiError };
