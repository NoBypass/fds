<script>
    import './../app.css'
    import MagnifyIcon from '$lib/icons/MagnifyIcon.svelte'
    import CommandPalette from '$lib/components/CommandPalette.svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Text from '$lib/components/Text.svelte'
    import Divider from '$lib/components/Divider.svelte'
    import Button from '$lib/components/Button.svelte'

    let inputRef
    let showCommandPalette = false

    const links = [
        'Search', 'Dashboard', 'SkyBlock'
    ]

    const toggleCommandPalette = () => {
        showCommandPalette = !showCommandPalette
    }

    const closeCommandPalette = () => {
        showCommandPalette = false
    }
</script>

<style>
    :global(body) {
        @apply bg-black text-white;
    }
</style>

<CommandPalette on:close={closeCommandPalette} open={showCommandPalette}>test <br> test <br> test</CommandPalette>

<ResponsiveContainer>
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
            <Button>Login</Button>
        </div>
    </nav>

    <slot />
</ResponsiveContainer>
