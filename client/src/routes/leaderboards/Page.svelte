<script>
import Button from '$lib/components/Button.svelte'
import { createEventDispatcher } from 'svelte'

export let active = false
export let disabled = false

const dispatch = createEventDispatcher()
</script>

<li class="rounded-lg transition-colors duration-150 {active ? 'bg-purple-500/30 bg-radial' : 'bg-transparent'}">
    <Button {disabled} on:click={(e) => dispatch('click', e)} type="transparent" tw="grid place-content-center h-9 w-9">
        <slot />
    </Button>
</li>