<div id="form">
    <h1>Cadastre-se</h1>
    <div id="fields">
        <div>
            <div>Novo usuário</div>
            <input type="text" autocomplete="username">
        </div>
        <div id="passwords">
            <div>
                <div>Nova senha</div>
                <input type="password" autocomplete="password">
            </div>
            <div id="passwordField">
                <div>Confirmar a nova senha</div>
                <input type="password" autocomplete="password">
            </div>
        </div>
    </div>
</div>

<style>
h1 {
    margin: 0.1rem;
    text-align: center;
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
    &:hover{
        background-color: #3d692c;
        transition: background-color 350ms ease-in-out;
    }
}

</style>
