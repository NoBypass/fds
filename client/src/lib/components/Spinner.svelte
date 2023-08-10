<script lang="ts">
    export let tw = ''
    export let size: 'sm' | 'md' | 'lg' = 'md'
    export let color: 'primary' | 'warning' = 'primary'

    const colors = {
        primary: 'border-b-purple-500',
        warning: 'border-b-amber-500',
    }

    const sizes = {
        sm: 'w-4 h-4',
        md: 'w-6 h-6',
        lg: 'w-8 h-8',
    }
</script>

<div class="relative flex {sizes[size]} {tw}">
    <i class="absolute w-full h-full rounded-full animate-spinner-ease border-solid border-t-transparent border-l-transparent border-r-transparent border-3 {colors[color]}"></i>
    <i class="absolute w-full h-full rounded-full animate-spinner-linear opacity-75 border-dotted border-t-transparent border-l-transparent border-r-transparent border-3 {colors[color]}"></i>
</div>