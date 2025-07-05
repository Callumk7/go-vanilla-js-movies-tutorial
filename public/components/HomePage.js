import { API } from "../services/api.js";
import { MovieItem } from "./MovieItem.js";

export class HomePage extends HTMLElement { // <home-page>
    async render() {
        const topMovies = await API.getTopMovies();
        const randomMovies = await API.getRandomMovies();
        renderMoviesInList(topMovies, document.querySelector("#top-20 ul"));
        renderMoviesInList(randomMovies, document.querySelector("#random ul"));

        function renderMoviesInList(movies, ul) {
            ul.innerHTML = "";
            movies.forEach(movie => {
                const li = document.createElement("li");
                li.appendChild(new MovieItem(movie));
                ul.appendChild(li);
            })
        }
    }

    // This is overriding the super method, it acts as a mount event
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.content.cloneNode(true);
        this.appendChild(content);

        this.render();
    }
}

customElements.define("home-page", HomePage);
