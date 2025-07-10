import API from "../services/api.js";
import { CollectionPage } from "./CollectionPage.js";

export class WatchlistPage extends CollectionPage {
    constructor() {
        super(API.getWatchlist, "Favorite Movies")
    }
}
customElements.define("watchlist-page", WatchlistPage)
