import type { MiddlewareResponseHandler } from 'astro';
import { apiUrl } from '../services/api';

export const onRequest: MiddlewareResponseHandler = async (context, next) => {
  if (!context.url.pathname.startsWith('/admin')) {
    return next();
  }

  const token = context.cookies.get('token')?.value;

  if (token) {
    try {
      // Validate token with backend
      const response = await fetch(apiUrl('/auth/me'), {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (response.ok) {
        const data = await response.json();
        context.locals.user = data.data;
        context.locals.token = token;
      }
    } catch (e) {
      // Token invalid, continue as guest
    }
  }

  return next();
};
