import { redirect } from '@sveltejs/kit';

export async function handle({ event, resolve }){ 
    const sessionId = event.cookies.get("sessionId");
    if (sessionId || event.url.pathname.startsWith("/register") || event.url.pathname.startsWith("/login")){
        return await resolve(event)
    }
    return redirect(301, '/register')
}
