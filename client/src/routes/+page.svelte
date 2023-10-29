<script lang="ts">
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import Input from '$lib/components/Input.svelte'
    import Text from '$lib/components/Text.svelte'
    import Card from '$lib/components/Card.svelte'
    import type { SvelteComponent } from 'svelte'

    let inputRef: undefined | SvelteComponent | HTMLElement

    $: if (inputRef && typeof inputRef.focus === 'function') {
        inputRef.focus()
    }
</script>

<style>
    .cards {
        & div {
            @apply grid grid-rows-6 h-64;
        }
    }
</style>

<main class="flex flex-col items-center justify-center">
    <section class="w-full lg:gap-32 gap-20 lg:mt-22 mt-14 flex flex-col items-center">
        <Text tw="text-4xl lg:text-6xl text-center">Search for <Text type="s" color="gradient" b>any<br>Player</Text> on the network</Text>
        <Input bind:this={inputRef} placeholder="Search for a Player / Guild / Leaderboard" rounded tw="lg:w-3/5 md:w-4/5 sm:w-11/12">
            <MagnifyIcon slot="left" tw="w-5 h-5 text-neutral-400" />
            <Text slot="right" b type="h6" tw="text-neutral-400">CTRL+K</Text>
        </Input>

        <div class="cards grid text-center md:grid-cols-3 w-full 3xl:px-36 gap-4">
            <Card>
                <Text b type="h3">Country Leaderboards</Text>
                <Text o>...and many more unique ways to compare players.</Text>
            </Card>
            <Card>
                <Text b type="h3">Legal Sessions</Text>
                <Text o>...which you can track using our safe open source app!</Text>
            </Card>
            <Card>
                <Text b type="h3">Ranked with ELO</Text>
                <Text o>...achieved through a complex game analysis tool.</Text>
            </Card>
        </div>
    </section>

    <section class="flex items-center mt-16">
        <Text tw="text-11xl" color="fade-r" b>3M</Text>
        <Text tw="text-lg -translate-x-14 w-104">View the statistics of over 3 million players across
            multiple servers. By registering you can view even more of your
            stats, climb leaderboards and much more!</Text>
    </section>

    <section class="flex my-16 items-center text-right">
        <Text tw="text-lg translate-x-10 w-104">You can choose to delete all your data that does not
            come from public API's at any time you want. Alongside many other features, you also
            get to choose to delay data to the public.</Text>

        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" class="h-28">
            <defs>
                <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="0%">
                    <stop offset="0%" style="stop-opacity: 0" />
                    <stop offset="100%" style="stop-color: #a855f7; stop-opacity: 1" />
                </linearGradient>
            </defs>
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z"
                  stroke="url(#gradient)" />
        </svg>
    </section>
</main>
