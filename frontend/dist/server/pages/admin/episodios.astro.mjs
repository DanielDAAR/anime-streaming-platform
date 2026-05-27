import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, $ as maybeRenderHead, H as addAttribute } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$AdminLayout } from '../../chunks/AdminLayout_DJ9OW8Ss.mjs';
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$Astro = createAstro();
const prerender = false;
const $$Episodios = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Episodios;
  let episodes = [];
  let animes = [];
  parseInt(Astro2.url.searchParams.get("page") || "1");
  try {
    const [epRes, animeRes] = await Promise.all([
      fetch(apiUrl("/episodes/latest?limit=50")),
      fetch(apiUrl("/animes?limit=100"))
    ]);
    if (epRes.ok) {
      const data = await epRes.json();
      episodes = data.data || [];
    }
    if (animeRes.ok) {
      const data = await animeRes.json();
      animes = data.data || [];
    }
  } catch (e) {
    console.error("Error fetching episodes:", e);
  }
  const animeMap = new Map(animes.map((a) => [a.id, a]));
  return renderTemplate`${renderComponent($$result, "AdminLayout", $$AdminLayout, { "title": "Gesti\xF3n de Episodios" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="level"> <div class="level-left"> <h2 class="title is-4">📺 Gestión de Episodios</h2> </div> <div class="level-right"> <button class="button is-primary" onclick="openEpisodeModal()"> <span class="icon">➕</span> <span>Nuevo Episodio</span> </button> </div> </div> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <div class="table-container"> <table class="table is-fullwidth is-hoverable"> <thead> <tr> <th>Anime</th> <th>Episodio</th> <th>Título</th> <th>Servidores</th> <th>Acciones</th> </tr> </thead> <tbody> ${episodes.map((episode) => renderTemplate`<tr${addAttribute(episode.id, "data-episode-id")}> <td>${animeMap.get(episode.animeId)?.title || "Unknown"}</td> <td>#${episode.number}</td> <td>${episode.title}</td> <td> <span class="tag">${episode.servers?.length || 0} servers</span> </td> <td> <div class="buttons are-small"> <button class="button is-danger delete-episode-btn"${addAttribute(episode.id, "data-id")}>
🗑️
</button> </div> </td> </tr>`)} </tbody> </table> </div> ${episodes.length === 0 && renderTemplate`<div class="has-text-centered py-5"> <p class="has-text-grey">No hay episodios registrados.</p> </div>`} </div>  <div class="modal" id="episode-modal"> <div class="modal-background"></div> <div class="modal-card"> <header class="modal-card-head"> <p class="modal-card-title">Nuevo Episodio</p> <button class="delete" onclick="closeEpisodeModal()"></button> </header> <section class="modal-card-body"> <form id="episode-form"> <div class="field"> <label class="label">Anime</label> <div class="control"> <div class="select is-fullwidth"> <select id="episode-anime-id" required> <option value="">Seleccionar anime...</option> ${animes.map((anime) => renderTemplate`<option${addAttribute(anime.id, "value")}>${anime.title}</option>`)} </select> </div> </div> </div> <div class="columns"> <div class="column"> <div class="field"> <label class="label">Número</label> <input class="input" type="number" id="episode-number" min="1" required> </div> </div> <div class="column"> <div class="field"> <label class="label">Duración (min)</label> <input class="input" type="number" id="episode-duration" min="0"> </div> </div> </div> <div class="field"> <label class="label">Título</label> <input class="input" type="text" id="episode-title" required> </div> <div class="field"> <label class="label">Descripción</label> <textarea class="textarea" id="episode-description" rows="2"></textarea> </div> <div class="field"> <label class="label">Servidores</label> <div id="servers-list"> <div class="server-input box mb-2" style="background: #252542;"> <div class="columns"> <div class="column is-4"> <input class="input is-small" placeholder="Nombre" data-server-name> </div> <div class="column is-5"> <input class="input is-small" placeholder="URL embed" data-server-url> </div> <div class="column is-3"> <div class="select is-small is-fullwidth"> <select data-server-quality> <option value="720p">720p</option> <option value="1080p">1080p</option> <option value="480p">480p</option> <option value="360p">360p</option> </select> </div> </div> </div> </div> </div> <button type="button" class="button is-small is-info" onclick="addServerInput()">+ Agregar Servidor</button> </div> </form> </section> <footer class="modal-card-foot"> <button class="button is-primary" onclick="saveEpisode()">Guardar</button> <button class="button" onclick="closeEpisodeModal()">Cancelar</button> </footer> </div> </div>  ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/episodios.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/episodios.astro";
const $$url = "/admin/episodios";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Episodios,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
