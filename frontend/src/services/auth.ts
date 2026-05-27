import type { User, AuthResponse } from '../types';

const TOKEN_KEY = 'token';
const USER_KEY = 'user';
const COOKIE_MAX_AGE = 60 * 60 * 24;

function setCookie(name: string, value: string) {
  document.cookie = `${name}=${encodeURIComponent(value)}; path=/; max-age=${COOKIE_MAX_AGE}; SameSite=Lax`;
}

function clearCookie(name: string) {
  document.cookie = `${name}=; path=/; max-age=0; SameSite=Lax`;
}

export const authService = {
  setAuth: (auth: AuthResponse) => {
    if (typeof window !== 'undefined') {
      localStorage.setItem(TOKEN_KEY, auth.token);
      localStorage.setItem(USER_KEY, JSON.stringify(auth.user));
      setCookie(TOKEN_KEY, auth.token);
    }
  },

  clearAuth: () => {
    if (typeof window !== 'undefined') {
      localStorage.removeItem(TOKEN_KEY);
      localStorage.removeItem(USER_KEY);
      clearCookie(TOKEN_KEY);
    }
  },

  getToken: (): string | null => {
    if (typeof window !== 'undefined') {
      return localStorage.getItem(TOKEN_KEY);
    }
    return null;
  },

  getUser: (): User | null => {
    if (typeof window !== 'undefined') {
      const user = localStorage.getItem(USER_KEY);
      return user ? JSON.parse(user) : null;
    }
    return null;
  },

  isAuthenticated: (): boolean => {
    return !!authService.getToken();
  },

  isAdmin: (): boolean => {
    const user = authService.getUser();
    return user?.role === 'admin';
  },

  hasRole: (role: string): boolean => {
    const user = authService.getUser();
    return user?.role === role;
  },
};
