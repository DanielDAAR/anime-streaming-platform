import { T as createComponent, $ as maybeRenderHead, H as addAttribute, a8 as renderTemplate, Q as createAstro } from './astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import 'clsx';

const $$Astro = createAstro();
const $$AnimeCard = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$AnimeCard;
  const { anime } = Astro2.props;
  const statusLabels = {
    ongoing: "En Emisi\xF3n",
    completed: "Finalizado",
    upcoming: "Pr\xF3ximamente",
    cancelled: "Cancelado"
  };
  return renderTemplate`${maybeRenderHead()}<a${addAttribute(`/anime/${anime.slug}`, "href")} class="anime-card"${addAttribute(anime.id, "data-anime-id")}> <div class="card-image"> <img${addAttribute(anime.images.poster, "src")}${addAttribute(anime.title, "alt")} loading="lazy"> <span class="rating-badge">⭐ ${anime.rating.toFixed(1)}</span> <span${addAttribute(`status-badge ${anime.status}`, "class")}>${statusLabels[anime.status]}</span> </div> <div class="card-content"> <h3 class="title">${anime.title}</h3> <p class="is-size-7 has-text-grey">${anime.episodesCount} episodios • ${anime.year}</p> <div class="genres"> ${anime.genres.slice(0, 3).map((genre) => renderTemplate`<span class="tag">${genre}</span>`)} </div> </div> </a>`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/AnimeCard.astro", void 0);

export { $$AnimeCard as $ };
