import { T as createComponent, $ as maybeRenderHead, H as addAttribute, a8 as renderTemplate, Q as createAstro, a1 as renderComponent, a6 as renderSlot } from './astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$MainLayout } from './MainLayout_DEgSZIfQ.mjs';
import 'clsx';
/* empty css                          */

const $$Astro$1 = createAstro();
const $$AdminSidebar = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro$1, $$props, $$slots);
  Astro2.self = $$AdminSidebar;
  const currentPath = Astro2.url.pathname;
  const menuItems = [
    { href: "/admin/dashboard", icon: "\u{1F4CA}", label: "Dashboard" },
    { href: "/admin/animes", icon: "\u{1F3AC}", label: "Gesti\xF3n Animes" },
    { href: "/admin/episodios", icon: "\u{1F4FA}", label: "Gesti\xF3n Episodios" },
    { href: "/admin/comentarios", icon: "\u{1F4AC}", label: "Moderaci\xF3n" },
    { href: "/admin/usuarios", icon: "\u{1F465}", label: "Usuarios" }
  ];
  return renderTemplate`${maybeRenderHead()}<aside class="admin-sidebar"> <div class="has-text-centered mb-5"> <span class="logo-text is-size-4">🎬 AnimeStream</span> <p class="is-size-7 has-text-grey">Panel Admin</p> </div> <p class="menu-label">General</p> <ul class="menu-list"> ${menuItems.map((item) => renderTemplate`<li> <a${addAttribute(item.href, "href")}${addAttribute(currentPath === item.href ? "is-active" : "", "class")}> <span class="icon">${item.icon}</span> <span>${item.label}</span> </a> </li>`)} </ul> <p class="menu-label mt-5">Sistema</p> <ul class="menu-list"> <li> <a href="/"> <span class="icon">🏠</span> <span>Ver Sitio</span> </a> </li> <li> <a href="#" id="admin-logout"> <span class="icon">🚪</span> <span>Cerrar Sesión</span> </a> </li> </ul> </aside> `;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/components/admin/AdminSidebar.astro", void 0);

const $$Astro = createAstro();
const $$AdminLayout = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$AdminLayout;
  const user = Astro2.locals.user;
  if (!user || user.role !== "admin") {
    return Astro2.redirect("/login");
  }
  const { title = "Panel de Administraci\xF3n" } = Astro2.props;
  return renderTemplate`${renderComponent($$result, "MainLayout", $$MainLayout, { "title": `${title} | Admin`, "description": "Panel de administraci\xF3n de AnimeStream", "data-astro-cid-2kanml4j": true }, { "default": ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="admin-layout" data-astro-cid-2kanml4j> <div class="columns is-gapless" data-astro-cid-2kanml4j> <div class="column is-2-desktop is-3-tablet admin-sidebar-wrapper" data-astro-cid-2kanml4j> ${renderComponent($$result2, "AdminSidebar", $$AdminSidebar, { "data-astro-cid-2kanml4j": true })} </div> <div class="column admin-content" data-astro-cid-2kanml4j> <header class="admin-header" data-astro-cid-2kanml4j> <div class="container-fluid" data-astro-cid-2kanml4j> <div class="level" data-astro-cid-2kanml4j> <div class="level-left" data-astro-cid-2kanml4j> <h1 class="title is-4" data-astro-cid-2kanml4j>${title}</h1> </div> <div class="level-right" data-astro-cid-2kanml4j> <div class="user-menu" data-astro-cid-2kanml4j> <span class="icon" data-astro-cid-2kanml4j>👤</span> <span id="admin-username" data-astro-cid-2kanml4j>Admin</span> </div> </div> </div> </div> </header> <main class="admin-main" data-astro-cid-2kanml4j> ${renderSlot($$result2, $$slots["default"])} </main> </div> </div> </div> ` })}  `;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/layouts/AdminLayout.astro", void 0);

export { $$AdminLayout as $ };
