<script lang="ts">
    import './../app.css'
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Divider from '$lib/components/Divider.svelte'
    import Button from '$lib/components/Button.svelte'
    import SigninModal from '$lib/components/Modals/SigninModal.svelte'
    import { signin } from '$lib/services/signin'
    import ConfirmationModal from '$lib/components/Modals/ConfirmationModal.svelte'
    import SuccessModal from '$lib/components/Modals/SuccessModal.svelte'
    import Avatar from '$lib/components/Avatar.svelte'
    import Dropdown from '$lib/components/Dropdown.svelte'
    import DropdownItem from '$lib/components/DropdownItem.svelte'

    let inputRef
    let showCommandPalette = false
    let showSigninModal = false
    let showSuccessModal = false
    let showConfirmationModal = false
    let token: string | undefined = undefined

    $: if (typeof localStorage !== 'undefined') {
        token = localStorage.getItem('token')
    }

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
        token = localStorage.getItem('token')
    }

    const hideConfirmationModal = () => {
        showConfirmationModal = true
    }

    const submitInfo = async (info: CustomEvent) => {
        const res = await signin(info.detail)
        let token
        let self
        if (typeof res !== 'string') {
            token = res.data.token
            self = res.data.username
        }

        if (token) {
            localStorage.setItem('token', token)
            localStorage.setItem('self', self)
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
            {#if (!token)}
                <Button on:click={openSigninModal}>Login</Button>
            {:else}
                <Dropdown>
                    <Avatar slot="trigger" />
                    <ul slot="content">
                        <DropdownItem>Settings</DropdownItem>
                        <DropdownItem color="danger">Logout</DropdownItem>
                    </ul>
                </Dropdown>
            {/if}
        </div>
    </nav>

    <slot />
</ResponsiveContainer>
