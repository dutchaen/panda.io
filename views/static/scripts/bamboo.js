
class AdminUser {
    constructor(json) {
        this.username = json["username"];
        this.name = json["name"];
        this.email = json["email"];
        this.bio = json["bio"];
        this.profile_photo_b64 = json["profile_photo_b64"];
        this.primary_color = json["primary_color"];
        this.secondary_color = json["secondary_color"];
        this.links = Link.fromArray(json["links"]);
    }
}

class PublicUser {
    constructor(json) {
        this.id = json["id"];
        this.username = json["username"];
        this.name = json["name"];
        this.bio = json["bio"];
        this.is_bitcoin_baller = json["is_bitcoin_baller"];
        this.profile_photo_b64 = json["profile_photo_b64"];
        this.primary_color = json["primary_color"];
        this.secondary_color = json["secondary_color"];
        this.links = Link.fromArray(json["links"]);
        this.created_at = json["created_at"];
    }
}

class Link {
    constructor(json) {
        this.name = json["name"];
        this.url = json["url"];
    }

    static fromArray(arr) {
        const posts = [];
        for (var i = 0; i < arr.length; i++) {
            posts.push(new Link(arr[i]));
        }
        return posts;
    }
}

async function login( username, password ) {
    const response = await fetch("/api/login", {
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({"username":username,"password":password})
    });

    const json = await response.json();
    if (response.status !== 200) {
        console.log(json);
        throw new Error(JSON.stringify(json));
        return;
    }


    await sleep(2000);
}

async function register( name, email, username, password ) {
    const response = await fetch("/api/register", {
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            "name": name, 
            "email": email,
            "username":username,
            "password":password
        })
    });

    const json = await response.json();
    if (response.status !== 200) {
        console.log(json);
        throw new Error(JSON.stringify(json));
        return;
    }

    await sleep(2000);
}

async function get_user( username ) {
    const response = await fetch(`/api/user/@${username}`, {
        method: "GET",
        credentials: "include",
    });

    const json = await response.json();
    if (response.status !== 200) {
        console.log(json);
        return undefined;
    }

    return new PublicUser(json);
}

async function get_me() {
    const response = await fetch("/api/account/@me", {
        method: "GET",
        credentials: "include",
    });

    const json = await response.json();
    if (response.status !== 200) {
        console.log(json);
        return undefined;
    }

    return new AdminUser(json);
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function set_me( admin_user ) {

    const response = await fetch("/api/account/@me", {
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(admin_user)
    });

    const json = await response.json();
    if (response.status !== 200) {
        console.log(json);
    }

    await sleep(2000);
}