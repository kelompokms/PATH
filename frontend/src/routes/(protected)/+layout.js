import { PUBLIC_API_URL } from "$env/static/public";
import { redirect } from "@sveltejs/kit";

export const load = async ({ fetch }) => {
    const res = await fetch(PUBLIC_API_URL + "auth", {
        credentials: "include",
    });
    console.log(res);
    if (!res.ok) {
        throw redirect(302, "/login");
        return { isAuth: false };
    }
    return { isAuth: true };
};
