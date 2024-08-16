<script>
    let username = "";
    let password = "";
    let error = "";

    async function doLogin() {
        let res;
        try {
            res = await fetch("http://localhost:8080/login", {
                method: "POST",
                body: JSON.stringify({
                    username,
                    password,
                }),
            });
        } catch (error) {
            console.log("erro na conexão com o server");
            return;
        } finally {
            const json = await res.json();
            result = JSON.stringify(json);
        }
    }
</script>

<div id="form">
    <h1>Fazer Login</h1>
    <div>Usuário</div>
    <input type="text" bind:value={username} autocomplete="username" />
    <div>Senha</div>
    <input
        type="password"
        bind:value={password}
        autocomplete="current-password"
    />
</div>
<div id="buttons">
    <button id="submit" on:click={doLogin} type="button">Confirmar</button>
    <a href="/register">Cadastre-se aqui</a>
</div>

<style>
    h1 {
        font-size: 2rem;
        margin: 0.1rem;
        text-align: center;
    }

    * {
        font-size: 1.2rem;
    }

    #buttons {
        margin-top: 1rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        color: gray;
        & > a {
            color: lightblue;
            text-decoration: none;
            &:hover {
                text-decoration: underline;
            }
        }
    }

    #form {
        display: flex;
        flex-flow: column;
        gap: 0.5rem;
    }

    #submit {
        box-shadow: 0px 0px 5px black;
        border-radius: 0.3rem;
        border: 0.1px solid black;
        cursor: pointer;
        font-size: 1rem;
        padding: 0.4rem 0.2rem 0.4rem 0.2rem;
        background-color: gray;
        &:hover {
            background-color: lightgray;
            transition: background-color 350ms ease-in-out;
        }
    }

    input {
        border-radius: 0.2rem;
        /*border: 1px solid darkblue;*/
        border: none;
        box-shadow: 0px 0px 4px black;
        background-color: white;
        padding: 0.2rem;
    }
</style>
