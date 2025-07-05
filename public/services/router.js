import { routes } from "./routes.js";

export const Router = {
    init: () => {
        window.addEventListener("popstate", () => {
            Router.go(location.pathname, false);
        });

        // Start at initial route
        Router.go(location.pathname + location.search)

        // Enhance current links in the document
        document.querySelectorAll("a.navlink").forEach((a) => {
            a.addEventListener("click", (event) => {
                event.preventDefault();
                const href = a.getAttribute("href");
                Router.go(href);
            })
        })
    },
    go: (route, addToHistory = true) => { // eg. login forms -> fix the history
        if (addToHistory) {
            history.pushState(null, "", route);
        };

        let pageElement = null;
        const routePath = route.includes("?") ? route.split("?")[0] : route;

        for (const r of routes) {
            if (typeof r.path === "string" && r.path === routePath) {
                pageElement = new r.component();
                break;
            } else if (r.path instanceof RegExp) {
                const match = r.path.exec(routePath);
                if (match) {
                    pageElement = new r.component();
                    const params = match.slice(1);
                    pageElement.params = params;
                    break;
                }
            }
        }


        if (pageElement == null) {
            pageElement = document.createElement("h1");
            pageElement.textContent = "Page not found";

        }           

        function updatePage() {
            document.querySelector("main").innerHTML = "";
            document.querySelector("main").appendChild(pageElement);
        }

        const oldPage = document.querySelector("main").firstElementChild;
        if (oldPage) {
            oldPage.style.viewTransitionName = "old";
        }
        pageElement.style.viewTransitionName = "new";

        if (!document.startViewTransition) {
            updatePage();
        } else {
            document.startViewTransition(() => {
                updatePage();
            })
        }
    },
}
