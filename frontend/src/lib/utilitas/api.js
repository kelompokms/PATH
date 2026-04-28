import { PUBLIC_DATABASE_URL } from '$env/static/public';


export async function get(url) {
    const response = await fetch(PUBLIC_DATABASE_URL + url);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
}

export async function post(url, data) {
    const response = await fetch(PUBLIC_DATABASE_URL + url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
}

