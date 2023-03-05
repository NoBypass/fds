export default async function fastFetch(url: string) {
    return await fetch(url).then(r => r.json())
}