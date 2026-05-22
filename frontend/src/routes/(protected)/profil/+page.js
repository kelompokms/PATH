import { PUBLIC_API_URL } from "$env/static/public";

export const load = async ({ fetch }) => {
    const res = await fetch(PUBLIC_API_URL + "user", {
        credentials: "include",
    });

    if (!res.ok) {
        const text = await res.text();
        throw new Error("Gagal mendapatkan pengguna:", text);
        return;
    }

    const json = await res.json();
    return json;
};
