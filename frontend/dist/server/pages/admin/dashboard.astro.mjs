import { T as createComponent, $ as maybeRenderHead, H as addAttribute, a8 as renderTemplate, Q as createAstro, a1 as renderComponent } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$AdminLayout } from '../../chunks/AdminLayout_DJ9OW8Ss.mjs';
import 'clsx';
/* empty css                                        */
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$Astro$1 = createAstro();
const $$StatsCard = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro$1, $$props, $$slots);
  Astro2.self = $$StatsCard;
  const { value, label, icon, color = "primary" } = Astro2.props;
  const colorMap = {
    primary: "#6c5ce7",
    success: "#00b894",
    warning: "#fdcb6e",
    danger: "#ff7675",
    info: "#00cec9"
  };
  return renderTemplate`${maybeRenderHead()}<div class="stat-card"${addAttribute(`border-top: 3px solid ${colorMap[color] || color}`, "style")} data-astro-cid-264jsjhh> <div class="stat-icon is-size-2 mb-2" data-astro-cid-264jsjhh>${icon}</div> <div class="stat-value" data-astro-cid-264jsjhh>${value}</div> <div class="stat-label" data-astro-cid-264jsjhh>${label}</div> </div> `;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/admin/StatsCard.astro", void 0);

const $$Astro = createAstro();
const prerender = false;
const $$Dashboard = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Dashboard;
  let stats = { latestAnimes: [], userStats: { total: 0, totalAdmins: 0, totalUsers: 0 } };
  const token = Astro2.cookies.get("token")?.value;
  try {
    const res = await fetch(apiUrl("/dashboard/stats"), {
      headers: token ? { "Authorization": `Bearer ${token}` } : {}
    });
    if (res.ok) {
      const data = await res.json();
      stats = data.data;
    }
  } catch (e) {
    console.error("Error fetching dashboard:", e);
  }
  return renderTemplate`${renderComponent($$result, "AdminLayout", $$AdminLayout, { "title": "Dashboard" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="columns is-multiline"> <div class="column is-3"> ${renderComponent($$result2, "StatsCard", $$StatsCard, { "value": stats.userStats?.total || 0, "label": "Total Usuarios", "icon": "\u{1F465}", "color": "primary" })} </div> <div class="column is-3"> ${renderComponent($$result2, "StatsCard", $$StatsCard, { "value": stats.userStats?.totalAdmins || 0, "label": "Administradores", "icon": "\u{1F464}", "color": "warning" })} </div> <div class="column is-3"> ${renderComponent($$result2, "StatsCard", $$StatsCard, { "value": stats.userStats?.totalUsers || 0, "label": "Usuarios Regulares", "icon": "\u{1F464}", "color": "success" })} </div> <div class="column is-3"> ${renderComponent($$result2, "StatsCard", $$StatsCard, { "value": stats.latestAnimes?.length || 0, "label": "Animes Recientes", "icon": "\u{1F4FA}", "color": "info" })} </div> </div> <div class="columns mt-5"> <div class="column is-8"> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <h3 class="title is-5">Animes Recientes</h3> ${stats.latestAnimes?.length > 0 ? renderTemplate`<div class="table-container"> <table class="table is-fullwidth is-hoverable"> <thead> <tr> <th>Título</th> <th>Estado</th> <th>Rating</th> <th>Episodios</th> </tr> </thead> <tbody> ${stats.latestAnimes.map((anime) => renderTemplate`<tr> <td> <a${addAttribute(`/anime/${anime.slug}`, "href")} class="has-text-primary">${anime.title}</a> </td> <td> <span${addAttribute(`tag ${anime.status === "ongoing" ? "is-success" : "is-primary"}`, "class")}> ${anime.status} </span> </td> <td>⭐ ${anime.rating.toFixed(1)}</td> <td>${anime.episodesCount}</td> </tr>`)} </tbody> </table> </div>` : renderTemplate`<p class="has-text-grey">No hay animes registrados aún.</p>`} </div> </div> <div class="column is-4"> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <h3 class="title is-5">⚡ Acciones Rápidas</h3> <div class="buttons are-medium"> <a href="/admin/animes" class="button is-primary is-fullwidth"> <span class="icon">🎬</span> <span>Gestionar Animes</span> </a> <a href="/admin/episodios" class="button is-info is-fullwidth"> <span class="icon">📺</span> <span>Gestionar Episodios</span> </a> <a href="/admin/comentarios" class="button is-warning is-fullwidth"> <span class="icon">💬</span> <span>Moderar Comentarios</span> </a> <a href="/admin/usuarios" class="button is-success is-fullwidth"> <span class="icon">👥</span> <span>Gestionar Usuarios</span> </a> </div> </div> </div> </div> ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/dashboard.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/dashboard.astro";
const $$url = "/admin/dashboard";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Dashboard,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
