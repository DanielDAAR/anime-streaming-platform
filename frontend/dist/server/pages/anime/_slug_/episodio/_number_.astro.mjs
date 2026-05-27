import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, H as addAttribute, $ as maybeRenderHead } from '../../../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$MainLayout } from '../../../../chunks/MainLayout_DEgSZIfQ.mjs';
import { $ as $$Footer, a as $$Navbar } from '../../../../chunks/Footer_CgIpEFuh.mjs';
import { a as apiUrl } from '../../../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../../../renderers.mjs';

var __freeze = Object.freeze;
var __defProp = Object.defineProperty;
var __template = (cooked, raw) => __freeze(__defProp(cooked, "raw", { value: __freeze(cooked.slice()) }));
var _a;
const $$Astro = createAstro();
const prerender = false;
const $$number = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$number;
  const { slug, number } = Astro2.params;
  let anime = null;
  let episode = null;
  let allEpisodes = [];
  if (slug && number) {
    try {
      const [animeRes, episodeRes, episodesRes] = await Promise.all([
        fetch(apiUrl(`/animes/${slug}`)),
        fetch(apiUrl(`/animes/${slug}/episodes/${number}`)),
        fetch(apiUrl(`/animes/${slug}/episodes?limit=100`))
      ]);
      if (animeRes.ok) anime = (await animeRes.json()).data;
      if (episodeRes.ok) episode = (await episodeRes.json()).data;
      if (episodesRes.ok) allEpisodes = (await episodesRes.json()).data || [];
    } catch (e) {
      console.error("Error fetching episode:", e);
    }
  }
  if (!anime || !episode) {
    return Astro2.redirect("/404");
  }
  const currentIndex = allEpisodes.findIndex((ep) => ep.id === episode.id);
  const prevEpisode = currentIndex > 0 ? allEpisodes[currentIndex - 1] : null;
  const nextEpisode = currentIndex < allEpisodes.length - 1 ? allEpisodes[currentIndex + 1] : null;
  const activeServer = episode.servers.find((s) => s.active) || episode.servers[0];
  return renderTemplate`${renderComponent($$result, "MainLayout", $$MainLayout, { "title": `${anime.title} - Episodio ${episode.number} | AnimeStream`, "description": `Ver ${anime.title} episodio ${episode.number}: ${episode.title}`, "canonical": `/anime/${anime.slug}/episodio/${episode.number}` }, { "default": async ($$result2) => renderTemplate(_a || (_a = __template([" ", " ", '<section class="section pt-4"> <div class="container"> <!-- Breadcrumb --> <nav class="breadcrumb has-succeeds-separator mb-4" aria-label="breadcrumbs"> <ul> <li><a href="/">Inicio</a></li> <li><a', ">", '</a></li> <li class="is-active"><a href="#" aria-current="page">Episodio ', '</a></li> </ul> </nav> <h1 class="title is-4 mb-4">', " - Episodio ", ": ", '</h1> <!-- Player --> <div class="player-container" id="player-container"> <iframe id="video-player"', ' allowfullscreen sandbox="allow-scripts allow-same-origin allow-presentation"></iframe> </div> <!-- Server Selector --> <div class="server-selector mt-4"> <span class="has-text-grey mr-3">Servidores:</span> ', ' </div> <!-- Episode Info --> <div class="box mt-5" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <h3 class="title is-5">Informaci\xF3n del Episodio</h3> ', ' <div class="tags mt-3"> ', ' <span class="tag">\u{1F4FA} Episodio ', " de ", '</span> </div> </div> <!-- Navigation --> <div class="level mt-5"> <div class="level-left"> ', ' </div> <div class="level-right"> ', ' </div> </div> <!-- Episode List --> <h3 class="title is-5 mt-6">Todos los Episodios</h3> <div class="columns is-multiline"> ', " </div> </div> </section> <script>\n    function changeServer(btn) {\n      const url = btn.getAttribute('data-url');\n      const player = document.getElementById('video-player');\n      if (player && url) {\n        player.src = url;\n      }\n      // Update active state\n      document.querySelectorAll('.server-btn').forEach(b => b.classList.remove('active'));\n      btn.classList.add('active');\n    }\n  <\/script> ", " "])), renderComponent($$result2, "Navbar", $$Navbar, {}), maybeRenderHead(), addAttribute(`/anime/${anime.slug}`, "href"), anime.title, episode.number, anime.title, episode.number, episode.title, addAttribute(activeServer?.url, "src"), episode.servers.map((server, idx) => renderTemplate`<button${addAttribute(`server-btn ${server.id === activeServer?.id ? "active" : ""}`, "class")}${addAttribute(server.url, "data-url")} onclick="changeServer(this)"> ${server.name} <span class="quality">(${server.quality})</span> </button>`), episode.description && renderTemplate`<p>${episode.description}</p>`, episode.duration && renderTemplate`<span class="tag">⏱️ ${episode.duration} min</span>`, episode.number, allEpisodes.length, prevEpisode && renderTemplate`<a${addAttribute(`/anime/${anime.slug}/episodio/${prevEpisode.number}`, "href")} class="button is-primary"> <span class="icon">←</span> <span>Ep. Anterior</span> </a>`, nextEpisode && renderTemplate`<a${addAttribute(`/anime/${anime.slug}/episodio/${nextEpisode.number}`, "href")} class="button is-primary"> <span>Ep. Siguiente</span> <span class="icon">→</span> </a>`, allEpisodes.map((ep) => renderTemplate`<div class="column is-3"> <a${addAttribute(`/anime/${anime.slug}/episodio/${ep.number}`, "href")}${addAttribute(`button is-fullwidth ${ep.id === episode.id ? "is-primary" : "is-dark"}`, "class")}>
Ep. ${ep.number} </a> </div>`), renderComponent($$result2, "Footer", $$Footer, {})) })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/anime/[slug]/episodio/[number].astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/anime/[slug]/episodio/[number].astro";
const $$url = "/anime/[slug]/episodio/[number]";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$number,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
