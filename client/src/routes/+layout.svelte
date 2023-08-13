<script lang="ts">
    import './../app.css'
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Divider from '$lib/components/Divider.svelte'
    import Button from '$lib/components/Button.svelte'
    import SigninModal from '$lib/components/Modals/SigninModal.svelte'
    import type { SigninInfo } from '$lib/types/signin'
    import { signin } from '$lib/api/signin'
    import ConfirmationModal from '$lib/components/Modals/ConfirmationModal.svelte'
    import SuccessModal from '$lib/components/Modals/SuccessModal.svelte'

    let inputRef
    let showCommandPalette = false
    let showSigninModal = false
    let showSuccessModal = false
    let showConfirmationModal = false

    const links = [
        'Search', 'Dashboard', 'SkyBlock'
    ]

    const toggleCommandPalette = () => {
        showCommandPalette = !showCommandPalette
    }

    const hideCommandPalette = () => {
        showCommandPalette = false
    }

    const openSigninModal = () => {
        showSigninModal = true
    }

    const hideSigninModal = () => {
        showSigninModal = false
    }

    const hideSuccessModal = () => {
        showSuccessModal = false
    }

    const hideConfirmationModal = () => {
        showConfirmationModal = true
    }

    const submitInfo = async (info: CustomEvent) => {
        const res = await signin(info.detail)
        let token
        if (typeof res !== 'string') token = res.data.token

        if (token) {
            localStorage.setItem('token', token)
            console.log(token)
            showSuccessModal = true
        } else {
            showConfirmationModal = true
        }
    }
</script>

<style>
    :global(body) {
        @apply bg-black text-white;
    }
</style>

<CommandPalette on:close={hideCommandPalette} open={showCommandPalette}>test <br> test <br> test</CommandPalette>

<ResponsiveContainer>
    <SigninModal on:close={hideSigninModal} open={showSigninModal} on:submit={submitInfo} />
    <SuccessModal open={showSuccessModal} on:close={hideSuccessModal} />
    <ConfirmationModal open={showConfirmationModal} on:close={hideConfirmationModal} />

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
