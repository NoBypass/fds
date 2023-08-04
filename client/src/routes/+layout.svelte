<script>
    import './../app.css'
    import MagnifyIcon from '$lib/icons/MagnifyIcon.svelte'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Divider from '$lib/components/Divider.svelte'
    import Button from '$lib/components/Button.svelte'
    import Modal from '$lib/components/Modal.svelte'
    import loginBg from '$lib/assets/login-bg.png'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'

    let inputRef
    let showCommandPalette = false
    let showSigninModal = false

    const links = [
        'Search', 'Dashboard', 'SkyBlock'
    ]

    const toggleCommandPalette = () => {
        showCommandPalette = !showCommandPalette
    }

    const closeCommandPalette = () => {
        showCommandPalette = false
    }

    const openSigninModal = () => {
        showSigninModal = true
    }

    const hideSigninModal = () => {
        showSigninModal = false
    }
</script>

<style>
    :global(body) {
        @apply bg-black text-white;
    }
</style>

<CommandPalette on:close={closeCommandPalette} open={showCommandPalette}>test <br> test <br> test</CommandPalette>

<ResponsiveContainer>
    <Modal open={showSigninModal} on:close={hideSigninModal} tw="w-2/3 h-160 flex space-between">
        <img src={loginBg}
             alt="background"
             class="w-1/2 h-full object-none z-20"
             style="-webkit-mask-image:-webkit-gradient(linear, left bottom, right bottom, from(rgba(0,0,0,1)), to(rgba(0,0,0,0)));
                    mask-image: linear-gradient(to right, rgba(0,0,0,1), rgba(0,0,0,0));">
        <div class="w-full">
            <div class="p-12 z-30 relative w-full h-full grid grid-cols-1 grid-rows-3">
                <div>
                    <Text type="h1" size="xl" b>Login/Register</Text>
                    <Text tw="mt-4">You can both register a new account and log into an existing account via the same form.</Text>
                </div>

                <div>
                    <Input light rounded placeholder="Minecraft username" tw="mt-12 w-full" />
                    <Input light rounded placeholder="Password" tw="mt-10 w-full" />
                    <Checkbox tw="mt-10"><Text size="md">Remember me!</Text></Checkbox>
                </div>

                <div class="grid w-full grid-cols-2 grid-rows-1 gap-32 mt-12 h-10 place-self-end">
                    <Button on:click={hideSigninModal} rounded color="transparent"><Text b>Cancel</Text></Button>
                    <Button rounded color="neutral"><Text b>Continue</Text></Button>
                </div>
            </div>
            <img src={loginBg} alt="background" class="saturate-150 absolute w-full h-full object-none blur-[100px]  z-10 left-0 top-0">
        </div>
    </Modal>

    <nav class="grid grid-rows-none grid-cols-6 w-full h-20">

        <div class="flex items-center">
            <Text type="h3" b>FDS</Text>
        </div>

        <div class="flex items-center col-span-2">
            <div bind:this={inputRef} class="w-full border transition-all duration-200 border-neutral-700 py-1 px-4 rounded-full hover:border-neutral-400 flex gap-2 items-center">
                <MagnifyIcon tw="w-5 h-5 text-neutral-400" />
                <input type="text"
                       placeholder="Search anything..."
                       class="focus:outline-none w-full bg-black"
                       on:focusin={toggleCommandPalette} />
                <Text b type="h6" tw="text-neutral-400">CTRL+K</Text>
            </div>
        </div>

        <div class="flex col-start-6 items-center gap-8 justify-self-end">
            <ul class="flex gap-4">
                {#each links as link}
                    <li><a href="/{link.toLowerCase()}">{link}</a></li>
                {/each}
            </ul>
            <Divider vertical tw="h-6" />
            <Button on:click={openSigninModal}>Login</Button>
        </div>
    </nav>

    <slot />
</ResponsiveContainer>
