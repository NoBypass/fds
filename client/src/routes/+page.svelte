<script lang="ts">
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import Input from '$lib/components/Input.svelte'
    import Text from '$lib/components/Text.svelte'
    import Card from '$lib/components/Card.svelte'
    import Index from '$lib/assets/images/index.png'
    import LB from '$lib/assets/images/lb.png'
    import Session from '$lib/assets/images/session.png'
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

    img {
        mask-image: linear-gradient(to top, rgba(0, 0, 0, 1.0) 50%, transparent 100%);
        @apply -z-10 w-full absolute bottom-0 left-0;
    }
</style>

<main class="flex flex-col items-center justify-center">
    <section class="w-full lg:gap-32 gap-20 lg:mt-22 mt-14 flex flex-col items-center">
        <Text tw="text-4xl lg:text-6xl text-center">Search for <Text type="s" color="gradient" b>any<br>Player</Text> on the network</Text>
        <Input bind:this={inputRef} placeholder="Search for a Player" rounded tw="lg:w-3/5 md:w-4/5 sm:w-11/12">
            <MagnifyIcon slot="left" tw="w-5 h-5 text-neutral-400" />
            <Text slot="right" b type="h6" tw="text-neutral-400">CTRL+K</Text>
        </Input>

        <div class="cards grid text-center md:grid-cols-3 w-full 3xl:px-20 gap-4">
            <Card>
                <Text b type="h3">Country Leaderboards</Text>
                <Text o>...and many more unique ways to compare players.</Text>
                <img src={LB} alt="Leaderboard showcase">
            </Card>
            <Card>
                <Text b type="h3">Legal Sessions</Text>
                <Text o>...which you can track using our safe open source app!</Text>
                <img src={Session} alt="Legal session showcase">
            </Card>
            <Card>
                <Text b type="h3">Ranked with ELO</Text>
                <Text o>...achieved through a complex game analysis tool.</Text>
                <img src={Index} alt="ELO showcase">
            </Card>
        </div>
    </section>

    <section class="grid grid-rows-2 grid-cols-1 md:grid-cols-2 md:grid-rows-1 items-center mt-16 w-full">
        <div class="grid w-full h-full relative">
            <Text tw="text-11xl place-self-center md:place-self-end" color="primary" b>3M</Text>
            <div class="bg-gradient-fade-b md:bg-gradient-fade-r h-full w-full absolute md:w-64 md:place-self-end" />
        </div>
        <Text tw="text-lg md:-translate-x-12 md:translate-y-0 -translate-y-32 text-center md:text-left max-w-lg">View the statistics of over 3 million players across
            multiple servers. By registering you can view even more of your
            stats, climb leaderboards and much more!</Text>
    </section>

    <section class="grid grid-rows-2 grid-cols-1 md:grid-cols-2 md:grid-rows-1 items-center mt-16">
        <Text tw="text-lg md:translate-x-10 md:translate-y-0 -translate-y-14 z-10 text-center md:text-right row-start-2 md:row-start-1 max-w-lg">You can choose to delete all your data that does not
            come from public API's at any time you want. Alongside many other features, you also
            get to choose to delay data to the public.</Text>
        <div class="grid w-full h-full relative z-0 row-start-1 md:col-start-2">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" class="h-32 text-purple-500 place-self-center md:justify-self-start md:self-center" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
            </svg>
            <div class="bg-gradient-fade-b md:bg-gradient-fade-r h-full w-full md:w-36 absolute md:rotate-180 translate-x-2" />
        </div>
    </section>
</main>
