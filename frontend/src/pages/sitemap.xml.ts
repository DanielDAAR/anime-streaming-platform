import type { APIRoute } from 'astro';
import { apiUrl } from '../services/api';

export const GET: APIRoute = async () => {
  const siteUrl = 'http://localhost:4321';

  // Fetch all animes for dynamic URLs
  let animes: any[] = [];
  try {
    const res = await fetch(apiUrl('/animes?limit=1000'));
    if (res.ok) {
      const data = await res.json();
      animes = data.data || [];
    }
  } catch (e) {
    console.error('Error fetching animes for sitemap:', e);
  }

  const staticUrls = [
    '',
    '/catalogo',
    '/buscar',
    '/login',
    '/registro',
  ];

  const animeUrls = animes.map(anime => `/anime/${anime.slug}`);
  const episodeUrls = animes.flatMap(anime => 
    Array.from({ length: anime.episodesCount || 0 }, (_, i) => 
      `/anime/${anime.slug}/episodio/${i + 1}`
    )
  );

  const allUrls = [...staticUrls, ...animeUrls, ...episodeUrls];

  const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${allUrls.map(url => `  <url>
    <loc>${siteUrl}${url}</loc>
    <lastmod>${new Date().toISOString().split('T')[0]}</lastmod>
    <changefreq>daily</changefreq>
    <priority>${url === '' ? '1.0' : url.startsWith('/anime/') ? '0.8' : '0.6'}</priority>
  </url>`).join('\n')}
</urlset>`;

  return new Response(sitemap, {
    headers: {
      'Content-Type': 'application/xml',
    },
  });
};
