<script lang="ts">
    export let tw = ''
    export let size: '2xl' | 'xl' | 'lg' | 'md' | 'sm' | 'xs' = '2xl'

    const sizes = {
        '2xl': '3xl:px-128 2xl:px-96 xl:px-40 md:px-24 sm:px-12 px-6',
        'xl': '3xl:px-96 2xl:px-64 xl:px-32 md:px-16 sm:px-8 px-4',
        'lg': '3xl:px-64 2xl:px-48 xl:px-24 md:px-12 sm:px-6 px-3',
        'md': '3xl:px-48 2xl:px-32 xl:px-16 md:px-8 sm:px-4 px-2',
        'sm': '3xl:px-32 2xl:px-24 xl:px-12 md:px-6 sm:px-3 px-1',
        'xs': '3xl:px-24 2xl:px-16 xl:px-8 md:px-4 sm:px-2 px-1',
    }
</script>

<div class="{sizes[size]} relative {tw}">
    <slot />
</div>