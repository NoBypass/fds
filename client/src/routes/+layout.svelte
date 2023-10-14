<script lang="ts">
    import './../app.css'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Button from '$lib/components/Button.svelte'
    import ConfirmationModal from '$lib/components/Modals/ConfirmationModal.svelte'
    import SuccessModal from '$lib/components/Modals/SuccessModal.svelte'
    import Avatar from '$lib/components/Avatar.svelte'
    import Dropdown from '$lib/components/Dropdown.svelte'
    import DropdownItem from '$lib/components/DropdownItem.svelte'
    import Shadows from '$lib/components/Shadows.svelte'
    import Alertbox from '$lib/components/Alertbox.svelte'
    import { mouseStore } from '$lib/stores/location'
    import GradientLine from '$lib/components/GradientLine.svelte'
    import Discord from '$lib/assets/icons/Discord.svelte'
    import Link from '$lib/components/Link.svelte'

    let showCommandPalette = false
    let showSuccessModal = false
    let showConfirmationModal = false
    let token: string | undefined

    $: if (typeof localStorage !== 'undefined' && showSuccessModal != undefined) {
        token = localStorage.getItem('token') || undefined
    }

    const links = [
        'Home', 'Player', 'Leaderboards', 'Download'
    ]

    const logout = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('self')
        token = undefined
    }
</script>

<style>
    :global(body) {
        @apply bg-slate-950 text-white;
    }
</style>

<svelte:window on:mousemove={(e) => mouseStore.move(e.clientX, e.clientY)}
               on:click={mouseStore.click}/>

<Alertbox />
<CommandPalette on:close={() => showCommandPalette = false} open={showCommandPalette}>test <br> test <br> test</CommandPalette>
<div class="z-0 opacity-40 w-screen overflow-hidden">
    <Shadows />
</div>
<ResponsiveContainer tw="bg-white/5">
    <nav class="grid grid-rows-none grid-cols-3 w-full h-20">

        <div class="flex items-center gap-2">
            <div class="h-9 w-9 rounded-lg bg-gradient-primary normal-b flex justify-center items-center">
                <Text b type="h2" tw="text-slate-950 -translate-y-0.5">F</Text>
            </div>
            <Text type="h3" b>FDS</Text>
        </div>

        <div class="flex items-center gap-8 justify-self-center">
            <ul class="flex border-gray-500/50 rounded-full px-0.5 py-0.5 border">
                {#each links as link}
                    <li class="py-2 px-4 ring-inset ring-white/40 hover:ring-1 rounded-full hover:bg-white/5 transition duration-300">
                        <a href="/{link === links[0] ? '' : link.toLowerCase()}">{link}</a>
                    </li>
                {/each}
            </ul>
        </div>

        <div class="justify-self-end self-center">
            {#if (!token)}
                <Button type="primary" href="/login">Login</Button>
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
<GradientLine />

<ResponsiveContainer>
    <SuccessModal open={showSuccessModal} on:close={() => showSuccessModal = false} />
    <ConfirmationModal open={showConfirmationModal} on:close={() => showConfirmationModal = false} />

    <slot />
</ResponsiveContainer>

<footer class="mt-16 fixed bottom-0">
    <GradientLine />
    <ResponsiveContainer tw="grid grid-cols-9 grid-rows-1 pt-6 pb-6">
        <Text color="darkened" tw="col-start-1 col-end-4">
            <Text type="h3" b>About</Text>
            This is a non-profit community project for all of the Hypixel community to enjoy. We are not affiliated with
            Hypixel or Mojang in any way. If you have questions, suggestions, or concerns, please join our <Link>Discord
            Server</Link> and open a ticket or if you want it to be worked on immediately, open an issue on our
            <Link>GitHub</Link> repository. If you want to contribute to the project, please read the
            <Link>How To Contribute</Link> page.
        </Text>

        <ul class="col-start-6 col-end-8">
            <Text color="darkened" type="h3" b>FDS</Text>
            <li><Link>Discord Server</Link></li>
            <li><Link>/g join FDS Employees</Link></li>
        </ul>

        <ul class="col-start-9">
            <Text color="darkened" type="h3" b>Project</Text>
            <li><Link>How To Contribute</Link></li>
            <li><Link>Roadmap</Link></li>
            <li><Link>GitHub</Link></li>
            <li><Link>YouTube</Link></li>
        </ul>
    </ResponsiveContainer>
</footer>
