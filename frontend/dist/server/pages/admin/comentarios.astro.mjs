import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, $ as maybeRenderHead, H as addAttribute } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$AdminLayout } from '../../chunks/AdminLayout_DJ9OW8Ss.mjs';
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$Astro = createAstro();
const prerender = false;
const $$Comentarios = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Comentarios;
  let comments = [];
  let total = 0;
  const page = parseInt(Astro2.url.searchParams.get("page") || "1");
  const token = Astro2.cookies.get("token")?.value;
  try {
    const res = await fetch(apiUrl(`/comments?page=${page}&limit=20`), {
      headers: token ? { "Authorization": `Bearer ${token}` } : {}
    });
    if (res.ok) {
      const data = await res.json();
      comments = data.data || [];
      total = data.meta?.total || 0;
    }
  } catch (e) {
    console.error("Error fetching comments:", e);
  }
  return renderTemplate`${renderComponent($$result, "AdminLayout", $$AdminLayout, { "title": "Moderaci\xF3n de Comentarios" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="level"> <div class="level-left"> <h2 class="title is-4">💬 Moderación de Comentarios</h2> </div> <div class="level-right"> <span class="tag is-info">${total} comentarios totales</span> </div> </div> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> ${comments.length > 0 ? renderTemplate`<div class="comments-list"> ${comments.map((comment) => renderTemplate`<div class="comment-item mb-4"${addAttribute(comment.id, "data-comment-id")}> <div class="is-flex is-justify-content-space-between is-align-items-start"> <div> <div class="comment-header"> <span class="has-text-weight-bold has-text-primary">${comment.user?.username || "Unknown"}</span> <span class="has-text-grey is-size-7 ml-2"> ${new Date(comment.createdAt).toLocaleString("es-ES")} </span> ${comment.isDeleted && renderTemplate`<span class="tag is-danger is-small ml-2">Eliminado</span>`} </div> <p class="mt-2">${comment.content}</p> <div class="tags mt-2"> <span class="tag is-small">👍 ${comment.likes}</span> <span class="tag is-small">🎬 ${comment.animeId}</span> </div> </div> <div class="buttons are-small"> ${!comment.isDeleted && renderTemplate`<button class="button is-danger delete-comment-btn"${addAttribute(comment.id, "data-id")}>
🗑️ Eliminar
</button>`} </div> </div> <hr style="background: #2d2d44; margin: 1rem 0;"> </div>`)} </div>` : renderTemplate`<div class="has-text-centered py-5"> <p class="has-text-grey">No hay comentarios para moderar.</p> </div>`} </div>  ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/comentarios.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/comentarios.astro";
const $$url = "/admin/comentarios";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Comentarios,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
