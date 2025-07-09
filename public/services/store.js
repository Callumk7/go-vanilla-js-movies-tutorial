const Store = {
    jwt: null,
    get loggedIn() {
        return this.jwt !== null;
    }
}

// When this module loads.. we can check to see if we have a value 
// in local storage.
// Infinite loop? No - we are not using the proxy here. We are using 
// a pure object.

if (localStorage.getItem("jwt")) {
    Store.jwt = localStorage.getItem("jwt")
}

const proxiedStore = new Proxy(Store, {
    // Think of this as a listener for changing a value. We use this API, 
    // and can add functionality. In this case, we set the item in the localStorage
    
    set: (target, prop, value) => {
        switch (prop) {
            case "jwt":
                target[prop] = value;
                if (value === null) {
                    localStorage.removeItem("jwt");
                } else {
                    localStorage.setItem("jwt", value);
                }
        }
        return true;
    }
});

export default proxiedStore;
