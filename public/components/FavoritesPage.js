import API from "../services/api.js";
import { CollectionPage } from "./CollectionPage.js";

export class FavoritesPage extends CollectionPage {
    constructor() {
        super(API.getFavorites, "Favorite Movies")
    }
}
