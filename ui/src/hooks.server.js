import { redirect } from "@sveltejs/kit";

export async function handle({ event, resolve }) {
	const sessionId = event.cookies.get("sessionId")
	isAuthRoute = (
		(event.url.pathname.startsWith('/login')) ||
		(event.url.pathname.startsWith('/register'))
	)

	if (!(sessionId) && !(isAuthRoute)) {
		return new redirect(303, '/login')
	}

	if (sessionId) {
		try {
			const response = await fetch('http://localhost:8080/validate', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({ session }),
			});

			if (!response.ok) {
				// If the session is invalid, clear the cookie and redirect to login
				event.cookies.delete('session', { path: '/' });
				return new Response.redirect(`${event.url.origin}/login`, 303);
			}

			// If the session is valid, you can add the user data to the event.locals
			const userData = await response.json();
			event.locals.user = userData;
		} catch (error) {
			console.error('Session validation error:', error);
			// Handle any errors (e.g., network issues) here
			// You might want to clear the session and redirect to login in case of errors
			event.cookies.delete('session', { path: '/' });
			return new Response.redirect(`${event.url.origin}/login`, 303);
		}
	}

	const response = await resolve(event);
	return response;
}


const response = await resolve(event); return response;
}
