import "./components/AnimatedLoading.js";
import "./components/YouTubeEmbed.js";
import { API } from "./services/api.js";
import { Router } from "./services/router.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
})

window.app = {
    Router,
    search: (event) => {
        event.preventDefault();

        const q = document.querySelector("input[type=search]").value;
        // TODO: Complete the search
    },
    api: API 
}
