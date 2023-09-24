<script lang="ts">
    import './../app.css'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Button from '$lib/components/Button.svelte'
    import SigninModal from '$lib/components/Modals/SigninModal.svelte'
    import ConfirmationModal from '$lib/components/Modals/ConfirmationModal.svelte'
    import SuccessModal from '$lib/components/Modals/SuccessModal.svelte'
    import Avatar from '$lib/components/Avatar.svelte'
    import Dropdown from '$lib/components/Dropdown.svelte'
    import DropdownItem from '$lib/components/DropdownItem.svelte'
    import Shadows from '$lib/components/Shadows.svelte'
    import Alertbox from '$lib/components/Alertbox.svelte'
    import { api } from '$lib/stores/api'

    let showCommandPalette = false
    let showSigninModal = false
    let showSuccessModal = false
    let showConfirmationModal = false
    let token: string | undefined

    $: if (typeof localStorage !== 'undefined' && showSuccessModal != undefined) {
        token = localStorage.getItem('token') || undefined
    }

    const links = [
        'Search', 'Dashboard', 'SkyBlock'
    ]

    const logout = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('self')
        token = undefined
    }

    const submitInfo = async (info: CustomEvent) => {
        const data: {
            name: string
            password: string
            remember: boolean
        } = info.detail
        const res = await $api.graphql.query<{readonly token: string, readonly name: string}>(`mutation {
            signin(name: "${data.name}", password: "${data.password}", remember: ${data.remember}) {
                token, name
            }
        }`, 'signin:'+data.name)

        if (res.token && res.name) {
            localStorage.setItem('token', res.token)
            localStorage.setItem('self', res.name)
            showSuccessModal = true
        } else showConfirmationModal = true
    }

    type Shadow = {
        opacity: number
        color: string
        height: number
        width: number
        offsetX: number
        offsetY: number
    }

    let shadows: undefined | Shadow[]
    for (let i = 0; i < 5; i++) {
        const opacity = Math.random()
        const color = `rgba(${Math.random()*255}, ${Math.random()*255}, ${Math.random()*255}, ${opacity})`
        const height = Math.random()*100
        const width = Math.random()*100
        const offsetX = Math.random()*100
        const offsetY = Math.random()*100
        shadows = [...shadows || [], { opacity, color, height, width, offsetX, offsetY }]
    }
</script>

<style>
    :global(body) {
        @apply bg-slate-950 text-white;
    }
</style>

<Alertbox />
<CommandPalette on:close={() => showCommandPalette = false} open={showCommandPalette}>test <br> test <br> test</CommandPalette>

<div class="z-0 opacity-60">
    <Shadows />
</div>
<ResponsiveContainer tw="bg-white/5">
    <SigninModal on:close={() => showSigninModal = false} open={showSigninModal} on:submit={submitInfo} />
    <nav class="grid grid-rows-none grid-cols-3 w-full h-20">

        <div class="flex items-center gap-2">
            <div class="h-9 w-9 rounded-lg bg-gradient-primary normal-b flex justify-center items-center">
                <Text b type="h2" tw="text-slate-950 -translate-y-0.5">F</Text>
            </div>
            <Text type="h3" b>FDS</Text>
        </div>

        <div class="flex items-center gap-8 justify-self-center">
            <ul class="flex gap-4 border-gray-500/50 rounded-full px-8 py-2.5 border">
                {#each links as link}
                    <li><a href="/{link.toLowerCase()}">{link}</a></li>
                {/each}
            </ul>
        </div>

        <div class="justify-self-end self-center">
            {#if (!token)}
                <Button on:click={() => showSigninModal = true}>Login</Button>
            {:else}
                <Dropdown>
                    <Avatar slot="trigger" />
                    <ul slot="content">
                        <DropdownItem>Settings</DropdownItem>
                        <DropdownItem on:click={logout} color="danger">Logout</DropdownItem>
                    </ul>
                </Dropdown>
            {/if}
        </div>
    </nav>
</ResponsiveContainer>
<div class="z-50 h-px w-screen bg-gradient-to-r from-transparent via-white/30 to-transparent" />

<ResponsiveContainer>
    <SuccessModal open={showSuccessModal} on:close={() => showSuccessModal = false} />
    <ConfirmationModal open={showConfirmationModal} on:close={() => showConfirmationModal = false} />

    <slot />
</ResponsiveContainer>
