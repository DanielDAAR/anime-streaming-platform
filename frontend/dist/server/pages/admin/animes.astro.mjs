import { T as createComponent, $ as maybeRenderHead, H as addAttribute, a8 as renderTemplate, a1 as renderComponent, Q as createAstro } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$AdminLayout } from '../../chunks/AdminLayout_DJ9OW8Ss.mjs';
import 'clsx';
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$AnimeFormModal = createComponent(async ($$result, $$props, $$slots) => {
  return renderTemplate`${maybeRenderHead()}<div class="modal" id="anime-modal"> <div class="modal-background"></div> <div class="modal-card"> <header class="modal-card-head"> <p class="modal-card-title" id="modal-title">Nuevo Anime</p> <button class="delete" aria-label="close" id="close-modal"></button> </header> <section class="modal-card-body"> <form id="anime-form"> <input type="hidden" id="anime-id"> <div class="field"> <label class="label">Título</label> <div class="control"> <input class="input" type="text" id="anime-title" required> </div> </div> <div class="field"> <label class="label">Slug (URL amigable)</label> <div class="control"> <input class="input" type="text" id="anime-slug" placeholder="se-genera-automaticamente"> </div> </div> <div class="field"> <label class="label">Descripción</label> <div class="control"> <textarea class="textarea" id="anime-description" rows="3" required></textarea> </div> </div> <div class="columns"> <div class="column"> <div class="field"> <label class="label">Estado</label> <div class="control"> <div class="select is-fullwidth"> <select id="anime-status"> <option value="ongoing">En Emisión</option> <option value="completed">Finalizado</option> <option value="upcoming">Próximamente</option> <option value="cancelled">Cancelado</option> </select> </div> </div> </div> </div> <div class="column"> <div class="field"> <label class="label">Año</label> <div class="control"> <input class="input" type="number" id="anime-year" min="1900" max="2100"${addAttribute((/* @__PURE__ */ new Date()).getFullYear(), "value")}> </div> </div> </div> </div> <div class="field"> <label class="label">Géneros (separados por coma)</label> <div class="control"> <input class="input" type="text" id="anime-genres" placeholder="Acción, Aventura, Fantasía" required> </div> </div> <div class="field"> <label class="label">Rating (0-10)</label> <div class="control"> <input class="input" type="number" id="anime-rating" min="0" max="10" step="0.1" value="0"> </div> </div> <div class="field"> <label class="label">URL Poster</label> <div class="control"> <input class="input" type="url" id="anime-poster" required> </div> </div> <div class="field"> <label class="label">URL Banner</label> <div class="control"> <input class="input" type="url" id="anime-banner"> </div> </div> <div class="field"> <label class="label">Studio</label> <div class="control"> <input class="input" type="text" id="anime-studio"> </div> </div> </form> </section> <footer class="modal-card-foot"> <button class="button is-primary" id="save-anime">Guardar</button> <button class="button" id="cancel-modal">Cancelar</button> </footer> </div> </div> `;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/admin/AnimeFormModal.astro", void 0);

const $$Astro = createAstro();
const prerender = false;
const $$Animes = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Animes;
  let animes = [];
  let total = 0;
  const page = parseInt(Astro2.url.searchParams.get("page") || "1");
  try {
    const res = await fetch(apiUrl(`/animes?page=${page}&limit=20`));
    if (res.ok) {
      const data = await res.json();
      animes = data.data || [];
      total = data.meta?.total || 0;
    }
  } catch (e) {
    console.error("Error fetching animes:", e);
  }
  return renderTemplate`${renderComponent($$result, "AdminLayout", $$AdminLayout, { "title": "Gesti\xF3n de Animes" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="level"> <div class="level-left"> <h2 class="title is-4">🎬 Gestión de Animes</h2> </div> <div class="level-right"> <button class="button is-primary" onclick="openAnimeModal()"> <span class="icon">➕</span> <span>Nuevo Anime</span> </button> </div> </div> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <div class="table-container"> <table class="table is-fullwidth is-hoverable"> <thead> <tr> <th>Poster</th> <th>Título</th> <th>Estado</th> <th>Rating</th> <th>Episodios</th> <th>Acciones</th> </tr> </thead> <tbody> ${animes.map((anime) => renderTemplate`<tr${addAttribute(anime.id, "data-anime-id")}> <td> <figure class="image is-48x48"> <img${addAttribute(anime.images.poster, "src")}${addAttribute(anime.title, "alt")} class="is-rounded"> </figure> </td> <td> <a${addAttribute(`/anime/${anime.slug}`, "href")} class="has-text-primary">${anime.title}</a> </td> <td> <span${addAttribute(`tag ${anime.status === "ongoing" ? "is-success" : anime.status === "completed" ? "is-primary" : "is-warning"}`, "class")}> ${anime.status} </span> </td> <td>⭐ ${anime.rating.toFixed(1)}</td> <td>${anime.episodesCount}</td> <td> <div class="buttons are-small"> <button class="button is-info edit-anime-btn"${addAttribute(JSON.stringify(anime), "data-anime")}>
✏️
</button> <button class="button is-danger delete-anime-btn"${addAttribute(anime.id, "data-id")}>
🗑️
</button> </div> </td> </tr>`)} </tbody> </table> </div> ${animes.length === 0 && renderTemplate`<div class="has-text-centered py-5"> <p class="has-text-grey">No hay animes registrados.</p> </div>`} </div> ${renderComponent($$result2, "AnimeFormModal", $$AnimeFormModal, {})}  ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/animes.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/animes.astro";
const $$url = "/admin/animes";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Animes,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
