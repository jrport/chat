export async function handle({ event, resolve }) {
    const sessionId = event.cookies.get("sessionId")
    if (!sessionId && !( event.url.pathname == '/login' || event.url.pathname == '/register' )){
        return new Response("Redirect", {headers: {Location: "/login"}, status: 303})
    };
    
    let response = await resolve(event)
    return response
}
