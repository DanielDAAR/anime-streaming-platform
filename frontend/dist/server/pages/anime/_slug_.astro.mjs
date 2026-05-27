import { T as createComponent, $ as maybeRenderHead, H as addAttribute, a8 as renderTemplate, Q as createAstro, a1 as renderComponent, ab as unescapeHTML } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$MainLayout } from '../../chunks/MainLayout_DEgSZIfQ.mjs';
import { a as $$Navbar, $ as $$Footer } from '../../chunks/Footer_CgIpEFuh.mjs';
import 'clsx';
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$Astro$2 = createAstro();
const $$EpisodeCard = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro$2, $$props, $$slots);
  Astro2.self = $$EpisodeCard;
  const { episode, animeSlug } = Astro2.props;
  return renderTemplate`${maybeRenderHead()}<a${addAttribute(`/anime/${animeSlug}/episodio/${episode.number}`, "href")} class="episode-card"> <div class="is-flex is-align-items-center"> <span class="episode-number">#${episode.number}</span> <div class="is-flex-grow-1"> <h4 class="title is-6 mb-1">${episode.title}</h4> ${episode.duration && renderTemplate`<p class="is-size-7 has-text-grey">${episode.duration} min</p>`} </div> <span class="icon has-text-primary">▶</span> </div> </a>`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/EpisodeCard.astro", void 0);

const $$Astro$1 = createAstro();
const $$CommentSection = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro$1, $$props, $$slots);
  Astro2.self = $$CommentSection;
  const { animeId, comments } = Astro2.props;
  return renderTemplate`${maybeRenderHead()}<div class="comment-section"${addAttribute(animeId, "data-anime-id")}> <h3 class="title is-4 mb-4">Comentarios</h3> <div class="comment-form mb-5"> <div class="field"> <div class="control"> <textarea class="textarea" id="comment-content" placeholder="Escribe tu comentario..." rows="3"></textarea> </div> </div> <div class="field"> <button class="button is-primary" id="submit-comment"> <span>Comentar</span> </button> </div> </div> <div id="comments-list"> ${comments.length === 0 ? renderTemplate`<div class="has-text-centered py-5"> <p class="has-text-grey">No hay comentarios aún. ¡Sé el primero!</p> </div>` : comments.map((comment) => renderTemplate`<div class="comment-item"${addAttribute(comment.id, "data-comment-id")}> <div class="comment-header"> <div class="avatar"> ${comment.user?.username?.charAt(0).toUpperCase() || "?"} </div> <div> <span class="username">${comment.user?.username || "Usuario"}</span> <span class="date">${new Date(comment.createdAt).toLocaleDateString("es-ES")}</span> </div> </div> <p class="comment-content">${comment.content}</p> <div class="comment-actions"> <button class="like-btn"${addAttribute(comment.id, "data-comment-id")}> <span>👍</span> <span class="likes-count">${comment.likes}</span> </button> <button class="reply-btn"${addAttribute(comment.id, "data-comment-id")}> <span>💬</span> <span>Responder</span> </button> </div> ${comment.replies && comment.replies.length > 0 && renderTemplate`<div class="replies"> ${comment.replies.map((reply) => renderTemplate`<div class="reply-item comment-item"${addAttribute(reply.id, "data-comment-id")}> <div class="comment-header"> <div class="avatar"> ${reply.user?.username?.charAt(0).toUpperCase() || "?"} </div> <div> <span class="username">${reply.user?.username || "Usuario"}</span> <span class="date">${new Date(reply.createdAt).toLocaleDateString("es-ES")}</span> </div> </div> <p class="comment-content">${reply.content}</p> </div>`)} </div>`} </div>`)} </div> </div> `;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/CommentSection.astro", void 0);

var __freeze = Object.freeze;
var __defProp = Object.defineProperty;
var __template = (cooked, raw) => __freeze(__defProp(cooked, "raw", { value: __freeze(cooked.slice()) }));
var _a;
const $$Astro = createAstro();
const prerender = false;
const $$slug = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$slug;
  const { slug } = Astro2.params;
  let anime = null;
  let episodes = [];
  let comments = [];
  if (!slug) {
    return Astro2.redirect("/404");
  }
  try {
    const animeRes = await fetch(apiUrl(`/animes/${slug}`));
    if (animeRes.ok) {
      const animeData = await animeRes.json();
      anime = animeData.data;
    }
    if (anime) {
      const episodesRes = await fetch(apiUrl(`/animes/${anime.slug}/episodes?limit=100`));
      if (episodesRes.ok) {
        const epData = await episodesRes.json();
        episodes = epData.data || [];
      }
      const commentsRes = await fetch(apiUrl(`/comments/${anime.id}?limit=50`));
      if (commentsRes.ok) {
        const commData = await commentsRes.json();
        comments = commData.data || [];
      }
    }
  } catch (e) {
    console.error("Error fetching anime detail:", e);
  }
  if (!anime) {
    return Astro2.redirect("/404");
  }
  const statusLabels = {
    ongoing: "En Emisi\xF3n",
    completed: "Finalizado",
    upcoming: "Pr\xF3ximamente",
    cancelled: "Cancelado"
  };
  return renderTemplate`${renderComponent($$result, "MainLayout", $$MainLayout, { "title": `${anime.title} | AnimeStream`, "description": anime.description.substring(0, 160), "image": anime.images.poster, "canonical": `/anime/${anime.slug}` }, { "default": async ($$result2) => renderTemplate` ${renderComponent($$result2, "Navbar", $$Navbar, {})}  ${maybeRenderHead()}<section class="hero-section"${addAttribute(`background: linear-gradient(135deg, #1a1a2e 0%, #0f0f1a 100%), url(${anime.images.banner || anime.images.poster}) center/cover;`, "style")}> <div class="container"> <div class="columns is-vcentered"> <div class="column is-3"> <figure class="image is-2by3"> <img${addAttribute(anime.images.poster, "src")}${addAttribute(anime.title, "alt")} class="has-radius" style="border-radius: 12px; box-shadow: 0 8px 24px rgba(0,0,0,0.5);"> </figure> </div> <div class="column is-9"> <div class="tags mb-3"> <span${addAttribute(`tag is-medium ${anime.status === "ongoing" ? "is-success" : anime.status === "completed" ? "is-primary" : "is-warning"}`, "class")}> ${statusLabels[anime.status]} </span> <span class="tag is-medium">⭐ ${anime.rating.toFixed(1)}</span> <span class="tag is-medium">${anime.episodesCount} episodios</span> <span class="tag is-medium">${anime.year}</span> </div> <h1 class="title is-2">${anime.title}</h1> <p class="subtitle">${anime.description}</p> <div class="tags"> ${anime.genres.map((genre) => renderTemplate`<span class="tag is-primary is-light">${genre}</span>`)} </div> ${anime.studio && renderTemplate`<p class="mt-3"><strong>Studio:</strong> ${anime.studio}</p>`} </div> </div> </div> </section>  <section class="section"> <div class="container"> <h2 class="title is-3">Episodios</h2> ${episodes.length > 0 ? renderTemplate`<div class="columns is-multiline"> ${episodes.map((episode) => renderTemplate`<div class="column is-6"> ${renderComponent($$result2, "EpisodeCard", $$EpisodeCard, { "episode": episode, "animeSlug": anime.slug })} </div>`)} </div>` : renderTemplate`<div class="has-text-centered py-6"> <p class="has-text-grey">No hay episodios disponibles aún.</p> </div>`} </div> </section>  <section class="section" style="background: #1a1a2e;"> <div class="container"> ${renderComponent($$result2, "CommentSection", $$CommentSection, { "animeId": anime.id, "comments": comments })} </div> </section>   ${renderComponent($$result2, "Footer", $$Footer, {})} `, "head": async ($$result2) => renderTemplate(_a || (_a = __template(['<script type="application/ld+json">', "<\/script>"])), unescapeHTML(JSON.stringify({
    "@context": "https://schema.org",
    "@type": "TVSeries",
    "name": anime.title,
    "description": anime.description,
    "image": anime.images.poster,
    "genre": anime.genres,
    "aggregateRating": {
      "@type": "AggregateRating",
      "ratingValue": anime.rating,
      "bestRating": 10
    },
    "numberOfEpisodes": anime.episodesCount
  }))) })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/anime/[slug].astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/anime/[slug].astro";
const $$url = "/anime/[slug]";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$slug,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
