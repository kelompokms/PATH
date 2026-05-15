import { PUBLIC_API_URL } from "$env/static/public";

export async function get(url) {
    const res = await fetch(PUBLIC_API_URL + url, {
        credentials: "include",
    });
    const json = await res.json();
    return json;
}

export async function post(url, data) {
    if (url == undefined) throw new Error("url kosong");
    if (data == undefined) throw new Error("data kosong");

    const res = await fetch(PUBLIC_API_URL + url, {
        credentials: "include",
        method: "POST",
        body: data,
    });
    return res;
}

export async function checkAuth() {
    const res = await fetch(PUBLIC_API_URL + "auth", {
        credentials: "include",
    });
    return res;
}
