<script lang="ts">
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import Input from '$lib/components/Input.svelte'
    import Text from '$lib/components/Text.svelte'
    import Card from '$lib/components/Card.svelte'

    let inputRef: undefined | HTMLElement

    $: if (inputRef && typeof inputRef.focus === 'function') {
        inputRef.focus()
    }
</script>

<main class="flex items-center justify-center">
    <div class="w-full gap-32 mt-32 flex flex-col items-center">
        <Text tw="text-6xl text-center">Search for <Text type="s" color="gradient" b>any<br>Player</Text> on the network</Text>
        <Input bind:this={inputRef} placeholder="Search for a Player / Guild / Leaderboard..." rounded tw="lg:w-3/5 md:w-4/5 sm:w-11/12">
            <MagnifyIcon slot="left" tw="w-5 h-5 text-neutral-400" />
            <Text slot="right" b type="h6" tw="text-neutral-400">CTRL+K</Text>
        </Input>

        <div>
            <Card></Card>
        </div>
    </div>
</main>
