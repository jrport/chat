<script>
    import Logo from "../Logo.svelte"
    import '../../app.css'
</script>

<main>
    <div id="loginForm">
        <Logo />
        <br>
        <slot></slot>
    </div>
</main>

<style>
main {
    height: 100dvh;
    display: flex;
    justify-content: center;
    align-items:center;
    & > #loginForm {
        padding: 2rem;
        border-radius: 1rem;
        background-color: #4c8435;
        border: 4px solid black;
    }
}

input {
    border-radius: 0.2rem;
    border: white;
}
</style>
