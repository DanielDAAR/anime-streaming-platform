import { renderers } from './renderers.mjs';
import { c as createExports, s as serverEntrypointModule } from './chunks/_@astrojs-ssr-adapter_BfkZmOwS.mjs';
import { manifest } from './manifest_BZ3-9JJP.mjs';

const _page0 = () => import('./pages/_image.astro.mjs');
const _page1 = () => import('./pages/404.astro.mjs');
const _page2 = () => import('./pages/admin/animes.astro.mjs');
const _page3 = () => import('./pages/admin/comentarios.astro.mjs');
const _page4 = () => import('./pages/admin/dashboard.astro.mjs');
const _page5 = () => import('./pages/admin/episodios.astro.mjs');
const _page6 = () => import('./pages/admin/usuarios.astro.mjs');
const _page7 = () => import('./pages/anime/_slug_/episodio/_number_.astro.mjs');
const _page8 = () => import('./pages/anime/_slug_.astro.mjs');
const _page9 = () => import('./pages/buscar.astro.mjs');
const _page10 = () => import('./pages/catalogo.astro.mjs');
const _page11 = () => import('./pages/login.astro.mjs');
const _page12 = () => import('./pages/registro.astro.mjs');
const _page13 = () => import('./pages/sitemap.xml.astro.mjs');
const _page14 = () => import('./pages/index.astro.mjs');

const pageMap = new Map([
    ["node_modules/astro/dist/assets/endpoint/node.js", _page0],
    ["src/pages/404.astro", _page1],
    ["src/pages/admin/animes.astro", _page2],
    ["src/pages/admin/comentarios.astro", _page3],
    ["src/pages/admin/dashboard.astro", _page4],
    ["src/pages/admin/episodios.astro", _page5],
    ["src/pages/admin/usuarios.astro", _page6],
    ["src/pages/anime/[slug]/episodio/[number].astro", _page7],
    ["src/pages/anime/[slug].astro", _page8],
    ["src/pages/buscar.astro", _page9],
    ["src/pages/catalogo.astro", _page10],
    ["src/pages/login.astro", _page11],
    ["src/pages/registro.astro", _page12],
    ["src/pages/sitemap.xml.ts", _page13],
    ["src/pages/index.astro", _page14]
]);
const serverIslandMap = new Map();
const _manifest = Object.assign(manifest, {
    pageMap,
    serverIslandMap,
    renderers,
    middleware: () => import('./_astro-internal_middleware.mjs')
});
const _args = {
    "mode": "standalone",
    "client": "file:///C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/dist/client/",
    "server": "file:///C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/dist/server/",
    "host": true,
    "port": 4321,
    "assets": "_astro"
};
const _exports = createExports(_manifest, _args);
const handler = _exports['handler'];
const startServer = _exports['startServer'];
const options = _exports['options'];
const _start = 'start';
{
	serverEntrypointModule[_start](_manifest, _args);
}

export { handler, options, pageMap, startServer };
