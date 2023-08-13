<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    export let color: 'neutral' | 'primary' | 'danger' = 'neutral'

    const colors = {
        neutral: 'hover:bg-neutral-600/30 text-white',
        primary: 'hover:bg-purple-600/20 text-white',
        danger: 'hover:bg-rose-600/10 text-rose-500',
    }

    const dispatch = createEventDispatcher()
    const click = (event) => dispatch('click', event)
</script>

<li class="{colors[color]} cursor-pointer px-2 py-1 rounded-md">
    <button class="border-none bg-transparent" on:click={click}>
        <slot />
    </button>
</li>