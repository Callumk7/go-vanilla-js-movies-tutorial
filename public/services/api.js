export const API = {
    baseURL: "/api",
    getTopMovies: async () => {
        return await API.fetch("movies/top")
    },
    getRandomMovies: async () => {
        return await API.fetch("movies/random")
    },
    getMovieByID: async (id) => {
        return await API.fetch(`movies/${id}`)
    },
    searchMovies: async (query, order, genre) => {
        return await API.fetch("movies/search", {query, order, genre})
    },
    getGenres: async () => {
        return await API.fetch("genres")
    },
    register: async (name, email, password) => {
        return await API.send("account/register", {email, password, name});
    },
    login: async (email, password) => {
        return await API.send("account/authenticate", {email, password});
    },
    getFavorites: async () => {
        return await API.fetch("account/favorites")
    },
    getWatchlist: async () => {
        return await API.fetch("account/watchlist")
    },
    saveToCollection: async (movie_id, collection) => {

    },
    send: async (serviceName, data) => {
        const body = JSON.stringify(data);
        try {
            const response = await fetch(`${API.baseURL}/${serviceName}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": app.Store.jwt ? `Bearer ${app.Store.jwt}` : null
                },
                body
            });
            const result = await response.json();
            return result;
        } catch (error) {
            console.error(error)   
        }
    },
    fetch: async (serviceName, args) => {
        try {
            const queryString = args ? new URLSearchParams(args).toString() : "";
            const response = await fetch(`${API.baseURL}/${serviceName}?${queryString}`, {
                headers: {
                    "Authorization": app.Store.jwt ? `Bearer ${app.Store.jwt}` : null
                },
            });
            const result = await response.json();
            return result;
        } catch (error) {
            console.error(error)   
        }
    }
};

export default API;

