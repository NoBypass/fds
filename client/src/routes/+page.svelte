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

<main class="flex items-center justify-center">
    <div class="w-full lg:gap-32 gap-20 lg:mt-22 mt-14 flex flex-col items-center">
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
    </div>
</main>
