<script>
    let username = "";
    let password = "";

    async function doLogin() {
        try {
            const res = await fetch('http://localhost:8080/login', {
                method: 'POST',
                body: JSON.stringify({
                    username,
                    password
                })
            })
        } catch (error) {
            console.log("erro na conexão com o server") 
            return
        } finally {
            const json = await res.json()
            result = JSON.stringify(json)
        }
	}

</script>

<div id="form">
    <div>Usuário</div>
    <input type="text" bind:value={username} autocomplete="username">
    <div>Senha</div>
    <input type="password" bind:value={password} autocomplete="password">
    <div id="registrate">
        <a href="/register">Cadastre-se aqui</a>
    </div>
</div>
<div id="buttons">
    <button id="submit" on:click={doLogin} type="button">Confirmar</button>
</div>


<style>
* {
    font-size: 1rem;
}

#buttons{
    display: flex;
    justify-content: space-between;
}
#passwordField {
    margin-bottom: 0.4rem;
}

#form {
    display: flex;
    flex-flow: column;
    gap: 0.5rem;
}

#registrate {
    margin-bottom: 0.6rem;
    & > a {
        color: darkblue;
        text-decoration: none;
        &:hover {
            text-decoration: underline;
        }
    }
}

#submit {
    border-radius: 0.3rem;
    border: 0.5px solid black;
    font-size: 1rem;
    padding: 0.4rem 0.2rem 0.4rem 0.2rem;
    background-color: #82bd69;
    &:hover{
        background-color: #3d692c;
        transition: background-color 350ms ease-in-out;
    }
}

input {
    border-radius: 0.2rem;
    border: white;
}


</style>
