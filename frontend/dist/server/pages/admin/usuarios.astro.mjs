import { T as createComponent, a1 as renderComponent, a8 as renderTemplate, Q as createAstro, $ as maybeRenderHead, H as addAttribute } from '../../chunks/astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import { $ as $$AdminLayout } from '../../chunks/AdminLayout_DJ9OW8Ss.mjs';
import { a as apiUrl } from '../../chunks/api_DaYyoT3w.mjs';
export { renderers } from '../../renderers.mjs';

const $$Astro = createAstro();
const prerender = false;
const $$Usuarios = createComponent(async ($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$Usuarios;
  let users = [];
  let total = 0;
  const page = parseInt(Astro2.url.searchParams.get("page") || "1");
  const token = Astro2.cookies.get("token")?.value;
  try {
    const res = await fetch(apiUrl(`/users?page=${page}&limit=20`), {
      headers: token ? { "Authorization": `Bearer ${token}` } : {}
    });
    if (res.ok) {
      const data = await res.json();
      users = data.data || [];
      total = data.meta?.total || 0;
    }
  } catch (e) {
    console.error("Error fetching users:", e);
  }
  return renderTemplate`${renderComponent($$result, "AdminLayout", $$AdminLayout, { "title": "Gesti\xF3n de Usuarios" }, { "default": async ($$result2) => renderTemplate` ${maybeRenderHead()}<div class="level"> <div class="level-left"> <h2 class="title is-4">👥 Gestión de Usuarios</h2> </div> <div class="level-right"> <span class="tag is-info">${total} usuarios</span> </div> </div> <div class="box" style="background: #1a1a2e; border: 1px solid #2d2d44;"> <div class="table-container"> <table class="table is-fullwidth is-hoverable"> <thead> <tr> <th>Usuario</th> <th>Email</th> <th>Rol</th> <th>Estado</th> <th>Registro</th> <th>Acciones</th> </tr> </thead> <tbody> ${users.map((user) => renderTemplate`<tr${addAttribute(user.id, "data-user-id")}> <td> <div class="is-flex is-align-items-center"> <div class="avatar mr-2" style="width: 32px; height: 32px; border-radius: 50%; background: #6c5ce7; display: flex; align-items: center; justify-content: center; color: white; font-weight: bold;"> ${user.username.charAt(0).toUpperCase()} </div> <span class="has-text-weight-medium">${user.username}</span> </div> </td> <td>${user.email}</td> <td> <span${addAttribute(`tag ${user.role === "admin" ? "is-danger" : "is-info"}`, "class")}> ${user.role} </span> </td> <td> <span${addAttribute(`tag ${user.isActive ? "is-success" : "is-warning"}`, "class")}> ${user.isActive ? "Activo" : "Inactivo"} </span> </td> <td>${new Date(user.createdAt).toLocaleDateString("es-ES")}</td> <td> <div class="buttons are-small"> <button class="button is-warning toggle-role-btn"${addAttribute(user.id, "data-id")}${addAttribute(user.role, "data-role")}> ${user.role === "admin" ? "\u{1F464} User" : "\u2699\uFE0F Admin"} </button> <button class="button is-info toggle-active-btn"${addAttribute(user.id, "data-id")}> ${user.isActive ? "\u{1F534} Desactivar" : "\u{1F7E2} Activar"} </button> </div> </td> </tr>`)} </tbody> </table> </div> ${users.length === 0 && renderTemplate`<div class="has-text-centered py-5"> <p class="has-text-grey">No hay usuarios registrados.</p> </div>`} </div>  ` })}`;
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/usuarios.astro", void 0);

const $$file = "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/pages/admin/usuarios.astro";
const $$url = "/admin/usuarios";

const _page = /*#__PURE__*/Object.freeze(/*#__PURE__*/Object.defineProperty({
  __proto__: null,
  default: $$Usuarios,
  file: $$file,
  prerender,
  url: $$url
}, Symbol.toStringTag, { value: 'Module' }));

const page = () => _page;

export { page };
