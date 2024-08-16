<div id="form">
    <h1>Cadastre-se</h1>
    <div class="form">
        <div>Novo usuário</div>
        <input type="text" autocomplete="username" />
    </div>
    <div class="form">
        <div>Nova senha</div>
        <input type="password" autocomplete="new-password" />
    </div>
    <div class='form'>
        <div>Confirmar a nova senha</div>
        <input type="password" autocomplete="new-password" />
    </div>
    <div id="wrapper">
        <button id="submit">Confirmar</button>
        <a href="/login">Já tenho uma conta</a>
    </div>
</div>

<style>
    * {
        font-size: 1.5rem;
    }

    .form {
        display: flex;
        flex-flow: column;
        gap: 0.5rem;
    }

    #form {
        display: flex;
        flex-flow: column;
        gap: 0.5rem;
    }

    #wrapper {
        display: flex;
        justify-content: space-between;
        align-items: center;
        & > button {
            padding: 0.6rem;
        }
        & > a {
            font-size: 1rem;
            color: white;
            text-decoration: none;
            &:hover {
                transform: scale(1.2);
                transition: all 0.3s ease-in-out;
                text-decoration: underline;
            }
        }
    }

    h1 {
        font-size: 3rem;
        margin: 0.1rem;
        text-align: center;
    }

    #submit {
        box-shadow: 0px 0px 5px black;
        border-radius: 0.3rem;
        border: 0.1px solid black;
        cursor: pointer;
        font-size: 1rem;
        padding: 0.4rem 0.2rem 0.4rem 0.2rem;
        background-color: #82bd69;
        &:hover {
            background-color: #3d692c;
            transition: background-color 350ms ease-in-out;
        }
    }

    input {
        border-radius: 0.2rem;
        /*border: 1px solid darkblue;*/
        border: none;
        box-shadow: 0px 0px 4px black;
        background-color: teal;
    }

    input {
        border-radius: 0.2rem;
        border: white;
    }
    #form {
        display: flex;
        flex-flow: column;
        gap: 1rem;
        margin-bottom: 1rem;
    }

    #fields {
        display: flex;
        flex-flow: column;
        gap: 0.9rem;
    }

    #passwords {
        display: flex;
        flex-flow: column;
        gap: 0.7rem;
    }

    #submit {
        border-radius: 0.3rem;
        font-size: 1rem;
        padding: 0.4rem 0.2rem 0.4rem 0.2rem;
        border: 0;
        background-color: #82bd69;
        &:hover {
            transform: scale(1.2);
            transition: all 0.4s ease-in-out;
            background-color: #81f4c3;
            /*transition: background-color 350ms ease-in-out;*/
        }
    }
</style>
