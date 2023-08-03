<script lang="ts">
    export let b = false
    export let size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' = 'md'
    export let type: 'p' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' = 'p'
    export let tw = ''

    const sizes = {
        xs: 'text-xs',
        sm: 'text-sm',
        md: 'text-md',
        lg: 'text-lg',
        xl: 'text-xl',
    }

    const className = `${sizes[size]} ${tw} ${b ? 'font-bold' : ''}`
</script>

<style>
    h1 {
        @apply text-4xl;
    }

    h2 {
        @apply text-2xl;
    }

    h3 {
        @apply text-xl;
    }

    h4 {
        @apply text-lg;
    }

    h5 {
        @apply text-sm;
    }

    h6 {
        @apply text-xs;
    }

    p {
        @apply opacity-70;
    }
</style>

{#if (type === 'p')}
    <p class={className}><slot /></p>
{:else if (type === 'h1')}
    <h1 class={className}><slot /></h1>
{:else if (type === 'h2')}
    <h2 class={className}><slot /></h2>
{:else if (type === 'h3')}
    <h3 class={className}><slot /></h3>
{:else if (type === 'h4')}
    <h4 class={className}><slot /></h4>
{:else if (type === 'h5')}
    <h5 class={className}><slot /></h5>
{:else if (type === 'h6')}
    <h6 class={className}><slot /></h6>
{/if}
