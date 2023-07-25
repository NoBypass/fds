<script lang="ts">
    export let vertical: boolean = true
    export let tw = ''

    $: length = vertical
        ? tw.includes('h-') ? '' : 'h-full'
        : tw.includes('w-') ? '' :  'w-full'
</script>

<div class="bg-neutral-300 {!vertical ? 'h-px' : 'w-px'} {length} {tw}"></div>