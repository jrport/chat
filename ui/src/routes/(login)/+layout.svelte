<script>
    import Typedtitle from "./components/typedtitle.svelte";

    let loaded = false;
    let submitNode;

    function toggleBtns() {
        loaded = true;
    }

    setTimeout(() => toggleBtns(), 3200);
</script>

<svelte:head>
    <link
        href="https://fonts.googleapis.com/css2?family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap"
        rel="stylesheet"
    />
</svelte:head>

<main>
    <div id="form_field">
        <Typedtitle text="Bem vindo ao Whatsapp 2!" typingSpeed={100}
        ></Typedtitle>
        <div
            bind:this={submitNode}
            style:display={loaded ? "flex" : "none"}
            id="button-wrapper"
        >
            <button>[ Entrar ]</button>
            <button>[ Me cadastrar ]</button>
        </div>
    </div>
</main>

<style>
    :global(body) {
        font-family: "Roboto Mono", monospace;
        font-optical-sizing: auto;
        margin: 0;
        background-color: #4d5356;
        color: white;
        font-size: 1rem;
    }
    #button-wrapper {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
    }
    button {
        background-color: transparent;
        border: none;
        border-radius: 0;
        padding: 0;
        color: white;
        &:hover {
            background-color: #667074;
        }
    }
    main {
        /*display: grid;*/
        height: 100dvh;
        display: flex;
        justify-content: center;
        align-items: center;
        /*grid-template-columns: 80% 20%;*/
    }
    #form_field {
        padding: 2rem;
        background-color: #3d4245;
        box-shadow: 0px 10px 10px black;
        grid-column-start: 2;
        border: 0.2rem solid gray;
        display: flex;
        flex-flow: column;
        gap: 0.4rem;
    }
</style>
