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
    import Link from '$lib/components/Link.svelte'

    const links = ['Home', 'Info', 'Leaderboards', 'Download']
    // const links = ['Home', 'Player', 'Leaderboards', 'Download']
    let showCommandPalette = false
    let showSuccessModal = false
    let showConfirmationModal = false
    let token: string | undefined
    let mobileMenuOpen = false
    let currentPath = links[0]

    $: if (typeof localStorage !== 'undefined' && showSuccessModal != undefined) {
        token = localStorage.getItem('token') || undefined
    }

    const logout = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('self')
        token = undefined
    }
</script>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&family=Raleway&display=swap');

    :global(body) {
        @apply bg-slate-950 text-white;
        font-family: Inter, sans-serif;
    }
</style>

<svelte:head>
    <title>FDS</title>
</svelte:head>
<svelte:window on:mousemove={(e) => mouseStore.move(e.clientX, e.clientY)}
               on:click={mouseStore.click}/>

<Alertbox />
<nav class="transition-opacity duration-150 absolute w-full h-screen bg-slate-950/60 backdrop-blur-xl z-10 {mobileMenuOpen ? 'opacity-100' : 'opacity-0 hidden'}">
    <ul class="mt-32 text-center flex flex-col gap-8">
        {#each links as link, i}
            <li class="transition-all duration-500 {mobileMenuOpen ? 'opacity-100' : 'opacity-0'}" style="transition-delay: {i*100}ms">
                <a on:click={() => currentPath = link} class="hover:text-white text-white/60 transition duration-150" on:click={() => mobileMenuOpen = false} href="/{link === links[0] ? '' : link.toLowerCase()}">
                    {#if link === currentPath}
                        <Text type="h1" color='gradient'>
                            {link}
                        </Text>
                    {:else}
                        <Text type="h1">
                            {link}
                        </Text>
                    {/if}
                </a>
            </li>
        {/each}
    </ul>
</nav>

<CommandPalette on:close={() => showCommandPalette = false} open={showCommandPalette}>test <br> test <br> test</CommandPalette>
<div class="z-0 opacity-40 w-full">
    <Shadows />
</div>
<ResponsiveContainer tw="z-20 bg-white/5 top-0 sticky backdrop-blur-md">
    <nav class="grid grid-rows-none grid-cols-3 w-full h-20">

        <div class="flex items-center gap-2">
            <div class="h-9 w-9 rounded-lg bg-gradient-primary normal-b flex justify-center items-center">
                <Text b type="h2" tw="text-slate-950 -translate-y-0.5">F</Text>
            </div>
            <Text type="h3" b>FDS</Text>
        </div>

        <div class="items-center gap-8 justify-self-center md:flex hidden">
            <ul class="flex border-gray-500/50 rounded-full px-0.5 py-0.5 border">
                {#each links as link}
                    <li class="py-2 px-4 ring-inset ring-white/40 hover:ring-1 rounded-full hover:bg-white/5 transition duration-300">
                        <a on:click={() => currentPath = link} href="/{link === links[0] ? '' : link.toLowerCase()}">
                            {#if link === currentPath}
                                <Text color='gradient' b>
                                    {link}
                                </Text>
                            {:else}
                                {link}
                            {/if}
                        </a>
                    </li>
                {/each}
            </ul>
        </div>

        <div class="col-start-3 gap-12 justify-self-end self-center flex">
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
            <Button on:click={() => mobileMenuOpen = !mobileMenuOpen} tw="{mobileMenuOpen ? '' : 'gap-1.5'} md:hidden w-10 flex items-center justify-center flex-col">
                <span class="transition-all duration-150 bg-white/60 w-6 h-[3px] block rounded-full {mobileMenuOpen ? 'rotate-45 translate-y-px' : ''}" />
                <span class="transition-all duration-150 bg-white/60 w-6 h-[3px] block rounded-full {mobileMenuOpen ? '-rotate-45 -translate-y-px' : ''}" />
            </Button>
        </div>
    </nav>
    <GradientLine tw="z-20" />
</ResponsiveContainer>

<main class="mt-16">
    <ResponsiveContainer>
        <SuccessModal open={showSuccessModal} on:close={() => showSuccessModal = false} />
        <ConfirmationModal open={showConfirmationModal} on:close={() => showConfirmationModal = false} />

        <slot />
    </ResponsiveContainer>
</main>

<footer class="mt-16 bottom-0">
    <GradientLine />
    <ResponsiveContainer tw="grid grid-cols-9 grid-rows-2 pt-6 pb-6">
        <Text color="darkened" tw="col-span-4 row-span-2">
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

        <ul class="md:col-start-9 col-start-6 md:row-start-1 row-start-2 col-span-4">
            <Text color="darkened" type="h3" b>Project</Text>
            <li><Link>How To Contribute</Link></li>
            <li><Link>Roadmap</Link></li>
            <li><Link>GitHub</Link></li>
            <li><Link>YouTube</Link></li>
        </ul>
    </ResponsiveContainer>
</footer>
