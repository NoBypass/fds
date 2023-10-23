<script lang="ts">
    import Text from '$lib/components/Text.svelte'
    import Button from '$lib/components/Button.svelte'
    import Slider from '$lib/components/Slider.svelte'
    import SlidePlaceholder from '$lib/components/Slides/SlidePlaceholder.svelte'
</script>

<main class="w-full items-center flex flex-col">
    <Text tw="text-3xl lg:text-5xl text-center">
        <Text type="s" color="gradient" b>Sessions</Text>, live <Text type="s" color="gradient" b>Tracking</Text>
        <br>
        and much more
    </Text>

    <section class="py-24 text-center grid gap-1.5">
        <Text color="darkened">100% Open-Source and Safe</Text>
        <Button type="primary" tw="pb-3 pt-2 px-10" href="https://github.com/NoBypass/fds" rounded>
            <Text type="h2" b>
                Download v1.0.0 <!-- TODO: Get version from GH API & Direct download latest stable -->
            </Text>
        </Button>
    </section>

    <section class="flex w-full text-center items-center flex-col">
        <Text tw="mb-6" type="h1" b>Features</Text>
        <Slider children={[SlidePlaceholder, SlidePlaceholder, SlidePlaceholder]}></Slider>
    </section>
</main>
