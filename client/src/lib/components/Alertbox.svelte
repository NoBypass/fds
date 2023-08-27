<script lang="ts">
    import Close from '$lib/assets/icons/Close.svelte'
    import { alerts } from '$lib/stores/alertStore'

    const colors = {
        primary: 'bg-purple-500',
        secondary: 'bg-blue-500',
        warning: 'bg-yellow-500',
        danger: 'bg-red-500',
        success: 'bg-green-500',
        info: 'bg-neutral-800',
    }
</script>

<div class="absolute w-full grid z-50 h-0">
    {#each $alerts as alert}
        <div style="opacity: {alert.stage}; transition: opacity 0.5s ease-in-out;"
             class="rounded-xl place-self-end max-w-[350px] min-w-[200px] p-3 grid grid-cols-10 mt-4 mr-4 {colors[alert.color]}">
            {#if alert.icon}
                <div class="flex items-center justify-center rounded-full">
                    Icon
                </div>
            {/if}
            <p class="px-4 self-center text-white col-start-2 col-end-9">{alert.message}</p>
            <Close tw="opacity-70 w-5/6 text-white place-self-end self-center col-start-10" />
        </div>
    {/each}
</div>
