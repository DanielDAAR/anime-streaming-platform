import { a as apiUrl } from './chunks/api_DaYyoT3w.mjs';
import 'es-module-lexer';
import './chunks/astro-designed-error-pages_D3ia-UCX.mjs';
import '@astrojs/internal-helpers/path';
import 'cookie';
import { s as sequence } from './chunks/index_DzTkI0dI.mjs';

const onRequest$1 = async (context, next) => {
  if (!context.url.pathname.startsWith("/admin")) {
    return next();
  }
  const token = context.cookies.get("token")?.value;
  if (token) {
    try {
      const response = await fetch(apiUrl("/auth/me"), {
        headers: {
          "Authorization": `Bearer ${token}`
        }
      });
      if (response.ok) {
        const data = await response.json();
        context.locals.user = data.data;
        context.locals.token = token;
      }
    } catch (e) {
    }
  }
  return next();
};

const onRequest = sequence(
	
	onRequest$1
	
);

export { onRequest };
