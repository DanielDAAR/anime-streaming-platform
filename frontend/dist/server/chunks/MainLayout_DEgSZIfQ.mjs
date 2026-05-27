import { T as createComponent, a8 as renderTemplate, a6 as renderSlot, a3 as renderHead, ab as unescapeHTML, H as addAttribute, Q as createAstro } from './astro/server_57n2fm-B.mjs';
import 'kleur/colors';
import 'clsx';
/* empty css                          */

var __freeze = Object.freeze;
var __defProp = Object.defineProperty;
var __template = (cooked, raw) => __freeze(__defProp(cooked, "raw", { value: __freeze(cooked.slice()) }));
var _a;
const $$Astro = createAstro();
const $$MainLayout = createComponent(($$result, $$props, $$slots) => {
  const Astro2 = $$result.createAstro($$Astro, $$props, $$slots);
  Astro2.self = $$MainLayout;
  const {
    title = "AnimeStream - Tu plataforma de anime",
    description = "Descubre los mejores animes, episodios y series en AnimeStream. Plataforma moderna de streaming con contenido embebido.",
    image = "/images/og-default.jpg",
    canonical = Astro2.url.pathname
  } = Astro2.props;
  const siteUrl = Astro2.site?.toString() || "http://localhost:4321";
  return renderTemplate(_a || (_a = __template(['<html lang="es"> <head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta name="description"', '><meta name="keywords" content="anime, streaming, episodios, series anime, ver anime online"><meta name="author" content="AnimeStream"><!-- Open Graph --><meta property="og:type" content="website"><meta property="og:url"', '><meta property="og:title"', '><meta property="og:description"', '><meta property="og:image"', '><!-- Twitter Card --><meta name="twitter:card" content="summary_large_image"><meta name="twitter:title"', '><meta name="twitter:description"', '><meta name="twitter:image"', '><!-- Canonical URL --><link rel="canonical"', '><!-- Favicon --><link rel="icon" type="image/svg+xml" href="/favicon.svg"><!-- Fonts --><link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin><title>', '</title><!-- JSON-LD Structured Data --><script type="application/ld+json">', "<\/script>", "</head> <body> ", "  </body> </html>"])), addAttribute(description, "content"), addAttribute(`${siteUrl}${canonical}`, "content"), addAttribute(title, "content"), addAttribute(description, "content"), addAttribute(`${siteUrl}${image}`, "content"), addAttribute(title, "content"), addAttribute(description, "content"), addAttribute(`${siteUrl}${image}`, "content"), addAttribute(`${siteUrl}${canonical}`, "href"), title, unescapeHTML(JSON.stringify({
    "@context": "https://schema.org",
    "@type": "WebSite",
    "name": "AnimeStream",
    "url": siteUrl,
    "potentialAction": {
      "@type": "SearchAction",
      "target": `${siteUrl}/buscar?q={search_term_string}`,
      "query-input": "required name=search_term_string"
    }
  })), renderHead(), renderSlot($$result, $$slots["default"]));
}, "C:/Users/espec/OneDrive/Desktop/anime-streaming-platform/frontend/src/layouts/MainLayout.astro", void 0);

export { $$MainLayout as $ };
