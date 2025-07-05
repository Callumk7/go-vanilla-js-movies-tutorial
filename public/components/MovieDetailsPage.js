import {API} from "../services/api.js"

export class MovieDetailsPage extends HTMLElement {
    id = null;
    movie = null;

    async render() {
        try {
            this.movie = await API.getMovieByID(this.id);
        } catch (error) {
            alert("This movie does not exist, from a component") // TODO: Replace this alert
            return
        }
        const template = document.getElementById("template-movie-details");
        const content = template.content.cloneNode(true);
        this.appendChild(content);

        this.querySelector("h2").textContent = this.movie.title;
        this.querySelector("h3").textContent = this.movie.tagline;
    }

    connectedCallback() {
        this.id = 14; // TODO: Update to use props
        this.render();
    }
}

customElements.define("movie-details-page", MovieDetailsPage);
