import { API } from "./services/api.js";

window.app = {
    search: (event) => {
        event.preventDefault();

        const q = document.querySelector("input[type=search]").value;
        // TODO: Complete the search
    },
    api: API 
}
