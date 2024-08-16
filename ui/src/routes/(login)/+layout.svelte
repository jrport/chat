<script>
    import Logo from "../Logo.svelte";
    import "../../app.css";
</script>

<main>
    <div id="loginForm">
        <Logo />
        <br />
        <slot></slot>
    </div>
</main>

<style>
    main {
        height: 100dvh;
        display: flex;
        justify-content: center;
        align-items: center;
        & > #loginForm {
            padding: 2rem;
            border-radius: 1rem;
            /*background-color: #4c8435;*/
            background-color: #1f2532;
            /*background-color: #108554;*/

            /*border: 1px solid black;*/
        }
    }

    #loginForm {
        /*box-shadow: 10px 7px 2px teal;*/
        box-shadow: 0px 0px 8px black;
    }
</style>
