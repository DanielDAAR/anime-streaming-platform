import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, $ as maybeRenderHead, H as addAttribute, e as Fragment } from '../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$MainLayout } from '../chunks/MainLayout_DEgSZIfQ.mjs';
import { a as $$Navbar, $ as $$Footer } from '../chunks/Footer_CgIpEFuh.mjs';
import { $ as $$AnimeCard } from '../chunks/AnimeCard_BW7lxz4B.mjs';
import { a as apiUrl } from '../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const prerender = false;
const $$Catalogo = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Catalogo;
  const page = parseInt(Astro2.url.searchParams.get("page") || "1");
  const limit = 20;
  const search = Astro2.url.searchParams.get("search") || "";
  const genre = Astro2.url.searchParams.get("genre") || "";
  const status = Astro2.url.searchParams.get("status") || "";
  let animes = [];
  let total = 0;
  let totalPages = 0;
  try {
    const params = new URLSearchParams();
    params.set("page", String(page));
    params.set("limit", String(limit));
    if (search) params.set("search", search);
    if (genre) params.set("genre", genre);
    if (status) params.set("status", status);
    const res = await fetch(apiUrl(`/animes?${params}`));
    if (res.ok) {
      const data = await res.json();
      animes = data.data || [];
      total = data.meta?.total || 0;
      totalPages = data.meta?.totalPages || 0;
    }
  } catch (e) {
    console.error("Error fetching catalog:", e);
  }
  const genres = ["Acci\xF3n", "Aventura", "Comedia", "Drama", "Fantas\xEDa", "Romance", "Terror", "Sci-Fi", "Deportes"];
  const statuses = [
    { value: "ongoing", label: "En Emisi\xF3n" },
    { value: "completed", label: "Finalizado" },
    { value: "upcoming", label: "Pr\xF3ximamente" }
  ];
  return renderTemplate`${renderComponent($$result, "MainLayout", $$MainLayout, { "title": "Cat\xE1logo de Animes", "description": "Explora nuestro cat\xE1logo completo de animes. Filtra por g\xE9nero, estado y m\xE1s.", "canonical": `/catalogo?page=${page}` }, { "default": async ($$result2) => renderTemplate` ${renderComponent($$result2, "Navbar", $$Navbar, {})} ${maybeRenderHead()}<section class="section"> <div class="container"> <h1 class="title is-2">Catálogo</h1> <!-- Filters --> <div class="box mb-5" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <form method="GET" action="/catalogo"> <div class="columns"> <div class="column is-4"> <div class="field"> <label class="label">Buscar</label> <div class="control search-box"> <input type="text" name="search" class="input" placeholder="Nombre del anime..."${addAttribute(search, "value")}> </div> </div> </div> <div class="column is-3"> <div class="field"> <label class="label">Género</label> <div class="control"> <div class="select is-fullwidth"> <select name="genre"> <option value="">Todos</option> ${genres.map((g) => renderTemplate`<option${addAttribute(g, "value")}${addAttribute(genre === g, "selected")}>${g}</option>`)} </select> </div> </div> </div> </div> <div class="column is-3"> <div class="field"> <label class="label">Estado</label> <div class="control"> <div class="select is-fullwidth"> <select name="status"> <option value="">Todos</option> ${statuses.map((s) => renderTemplate`<option${addAttribute(s.value, "value")}${addAttribute(status === s.value, "selected")}>${s.label}</option>`)} </select> </div> </div> </div> </div> <div class="column is-2"> <div class="field"> <label class="label">&nbsp;</label> <button type="submit" class="button is-primary is-fullwidth">
Filtrar
</button> </div> </div> </div> </form> </div> <!-- Results --> ${animes.length > 0 ? renderTemplate`${renderComponent($$result2, "Fragment", Fragment, {}, { "default": async ($$result3) => renderTemplate` <p class="has-text-grey mb-4">Mostrando ${animes.length} de ${total} resultados</p> <div class="columns is-multiline"> ${animes.map((anime) => renderTemplate`<div class="column is-3-desktop is-4-tablet is-6-mobile"> ${renderComponent($$result3, "AnimeCard", $$AnimeCard, { "anime": anime })} </div>`)} </div>  ${totalPages > 1 && renderTemplate`<nav class="pagination is-centered mt-6" role="navigation"> <a${addAttribute(`pagination-previous ${page <= 1 ? "is-disabled" : ""}`, "class")}${addAttribute(page > 1 ? `/catalogo?page=${page - 1}&search=${search}&genre=${genre}&status=${status}` : "#", "href")}>
← Anterior
</a> <a${addAttribute(`pagination-next ${page >= totalPages ? "is-disabled" : ""}`, "class")}${addAttribute(page < totalPages ? `/catalogo?page=${page + 1}&search=${search}&genre=${genre}&status=${status}` : "#", "href")}>
Siguiente →
</a> <ul class="pagination-list"> ${Array.from({ length: totalPages }, (_, i) => i + 1).map((p) => renderTemplate`<li> <a${addAttribute(`pagination-link ${p === page ? "is-current" : ""}`, "class")}${addAttribute(`/catalogo?page=${p}&search=${search}&genre=${genre}&status=${status}`, "href")}> ${p} </a> </li>`)} </ul> </nav>`}` })}` : renderTemplate`<div class="has-text-centered py-6"> <p class="is-size-4 has-text-grey">No se encontraron animes</p> <a href="/catalogo" class="button is-primary mt-4">Ver todos</a> </div>`} </div> </section> ${renderComponent($$result2, "Footer", $$Footer, {})} ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/catalogo.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/catalogo.astro";
const $$url = "/catalogo";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Catalogo,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
