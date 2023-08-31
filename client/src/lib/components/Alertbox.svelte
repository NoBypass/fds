<script lang="ts">
    import Close from '$lib/assets/icons/Close.svelte'
    import { alerts, alertStore } from '$lib/stores/alertStore'
    import Button from '$lib/components/Button.svelte'

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
        <div style="opacity: {alert.stage};
                    transform: translateY({alert.stage * 20}px);"
             class="rounded-xl place-self-end max-w-[350px] min-w-[200px] p-3 grid grid-cols-10 mt-4 mr-4 {colors[alert.color]}">
            {#if alert.icon}
                <div class="flex items-center justify-center rounded-full">
                    Icon
                </div>
            {/if}
            <p class="px-4 self-center text-white col-start-2 col-end-9">{alert.message}</p>
            <Button tw="place-self-end self-center col-start-10 w-2/3 h-full p-0"
                 color="transparent" on:click={() => alertStore.remove(alert)}>
                <Close tw="w-full text-white opacity-70 hover:opacity-90 duration-150" />
            </Button>
        </div>
    {/each}
</div>
