import{e as d}from"./api.BhIwWj_B.js";import"./auth.h-v8t-lg.js";import"./hoisted.jgWgEjQ_.js";window.openEpisodeModal=()=>{document.getElementById("episode-modal")?.classList.add("is-active")};window.closeEpisodeModal=()=>{document.getElementById("episode-modal")?.classList.remove("is-active"),document.getElementById("episode-form")?.reset()};window.addServerInput=()=>{const e=document.getElementById("servers-list"),t=document.createElement("div");t.className="server-input box mb-2",t.style.background="#252542",t.innerHTML=`
        <div class="columns">
          <div class="column is-4">
            <input class="input is-small" placeholder="Nombre" data-server-name />
          </div>
          <div class="column is-5">
            <input class="input is-small" placeholder="URL embed" data-server-url />
          </div>
          <div class="column is-3">
            <div class="select is-small is-fullwidth">
              <select data-server-quality>
                <option value="720p">720p</option>
                <option value="1080p">1080p</option>
                <option value="480p">480p</option>
                <option value="360p">360p</option>
              </select>
            </div>
          </div>
        </div>
      `,e?.appendChild(t)};window.saveEpisode=async()=>{const e=[];document.querySelectorAll(".server-input").forEach(o=>{const i=o.querySelector("[data-server-name]")?.value,a=o.querySelector("[data-server-url]")?.value,s=o.querySelector("[data-server-quality]")?.value;i&&a&&e.push({name:i,url:a,quality:s,active:e.length===0})});const t={animeId:document.getElementById("episode-anime-id")?.value,number:parseInt(document.getElementById("episode-number")?.value),title:document.getElementById("episode-title")?.value,description:document.getElementById("episode-description")?.value,duration:parseInt(document.getElementById("episode-duration")?.value)||0,servers:e};try{await d.create(t),window.closeEpisodeModal(),window.location.reload()}catch(o){alert("Error al crear episodio"),console.error(o)}};document.querySelectorAll(".delete-episode-btn").forEach(e=>{e.addEventListener("click",async()=>{if(!confirm("¿Eliminar este episodio?"))return;const t=e.getAttribute("data-id");try{await d.delete(t),window.location.reload()}catch{alert("Error al eliminar")}})});
