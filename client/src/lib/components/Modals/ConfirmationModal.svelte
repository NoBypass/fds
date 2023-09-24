<script lang="ts">
    import Modal from '$lib/components/Modal.svelte'
    import { createEventDispatcher, onMount } from 'svelte'
    import Text from '$lib/components/Text.svelte'
    import Button from '$lib/components/Button.svelte'
    import Discord from '$lib/assets/icons/Discord.svelte'

    export let open = false

    const dispatch = createEventDispatcher()
    const submit = () => {
        dispatch('submit', {
            // TODO
        })
    }
</script>

<Modal preventClose closeX="{false}" open={open}>
    <div class="p-5 text-center grid gap-6">
        <Text b type="h2">Verify that it's you</Text>
        <Text>You need to verify your Minecraft account via Discord</Text>
        <Button
                tw="flex justify-center gap-2"
                href="https://discord.gg/DCgAecjphr"
                on:click="{submit}">
            Verify <Discord tw="text-white h-6 w-6" />
        </Button>
    </div>
</Modal>