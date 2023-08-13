<script lang="ts">
    export let tw = ''
    export let progress = 1

    let strokeDasharray = ''
    let strokeDashoffset = 0
    $: strokeDasharray = `${progress * 100} ${100 - progress * 100}`
    $: strokeDashoffset = 25 - progress * 25
</script>

<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="{tw}">
    <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
            stroke-dasharray="{strokeDasharray}"
            stroke-dashoffset="{strokeDashoffset}"
    />
</svg>