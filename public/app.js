import "./components/AnimatedLoading.js";
import "./components/YouTubeEmbed.js";
import { API } from "./services/api.js";
import { Router } from "./services/router.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
})

window.app = {
    Router,
    showError: (message = "There was an error", goToHome = true) => {
        document.getElementById("alert-modal").showModal();
        document.querySelector("#alert-modal p").textContent = message;
        if (goToHome) app.Router.go("/");
    },
    closeError: () => {
        document.getElementById("alert-modal").close();
    },
    search: (event) => {
        event.preventDefault();

        const q = document.querySelector("input[type=search]").value;
        // TODO: Complete the search
    },
    api: API 
}
