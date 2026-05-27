import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, $ as maybeRenderHead, H as addAttribute, e as Fragment } from '../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$MainLayout } from '../chunks/MainLayout_DEgSZIfQ.mjs';
import { a as $$Navbar, $ as $$Footer } from '../chunks/Footer_CgIpEFuh.mjs';
import { $ as $$AnimeCard } from '../chunks/AnimeCard_BW7lxz4B.mjs';
import { a as apiUrl } from '../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../renderers.mjs';

const $$Astro = createAstro();
const prerender = false;
const $$Buscar = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Buscar;
  const query = Astro2.url.searchParams.get("q") || "";
  let animes = [];
  let total = 0;
  if (query) {
    try {
      const res = await fetch(apiUrl(`/animes?search=${encodeURIComponent(query)}&limit=50`));
      if (res.ok) {
        const data = await res.json();
        animes = data.data || [];
        total = data.meta?.total || 0;
      }
    } catch (e) {
      console.error("Error searching:", e);
    }
  }
  return renderTemplate`${renderComponent($$result, "MainLayout", $$MainLayout, { "title": query ? `Resultados para "${query}" | AnimeStream` : "Buscar Anime | AnimeStream", "description": "Busca tus animes favoritos en AnimeStream.", "canonical": `/buscar?q=${query}` }, { "default": async ($$result2) => renderTemplate` ${renderComponent($$result2, "Navbar", $$Navbar, {})} ${maybeRenderHead()}<section class="section"> <div class="container"> <h1 class="title is-2">Buscar Anime</h1> <form method="GET" action="/buscar" class="mb-6"> <div class="field has-addons"> <div class="control is-expanded search-box"> <input type="text" name="q" class="input is-large" placeholder="Escribe el nombre del anime..."${addAttribute(query, "value")} autofocus> </div> <div class="control"> <button type="submit" class="button is-primary is-large">
Buscar
</button> </div> </div> </form> ${query && renderTemplate`${renderComponent($$result2, "Fragment", Fragment, {}, { "default": async ($$result3) => renderTemplate` <p class="has-text-grey mb-4"> ${total > 0 ? `Se encontraron ${total} resultados` : "No se encontraron resultados"} </p> ${animes.length > 0 && renderTemplate`<div class="columns is-multiline"> ${animes.map((anime) => renderTemplate`<div class="column is-3-desktop is-4-tablet is-6-mobile"> ${renderComponent($$result3, "AnimeCard", $$AnimeCard, { "anime": anime })} </div>`)} </div>`}` })}`} ${!query && renderTemplate`<div class="has-text-centered py-6"> <p class="is-size-4 has-text-grey">Escribe algo para comenzar a buscar</p> </div>`} </div> </section> ${renderComponent($$result2, "Footer", $$Footer, {})} ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/buscar.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/buscar.astro";
const $$url = "/buscar";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Buscar,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
