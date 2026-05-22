import { PUBLIC_API_URL } from "$env/static/public";

export const load = async ({ params, fetch }) => {
    const res = await fetch(
        PUBLIC_API_URL + "class/" + params.kode + "/murid",
        {
            credentials: "include",
        },
    );

    const json = await res.json();

    return {
        murid: json,
    };
};
