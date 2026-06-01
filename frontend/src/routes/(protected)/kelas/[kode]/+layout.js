import { PUBLIC_API_URL } from "$env/static/public";
import { get } from "$lib/utils/api";

export const load = async ({ params, fetch }) => {
    const res = await fetch(PUBLIC_API_URL + "class/" + params.kode, {
        credentials: "include",
    });
    const json = await res.json();
    return json;
};
