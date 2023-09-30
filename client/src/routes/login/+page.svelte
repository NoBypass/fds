<script>
    import Card from '$lib/components/Card.svelte'
    import Text from '$lib/components/Text.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Button from '$lib/components/Button.svelte'

    const form = {
        name: '',
        password: '',
        remember: false
    }
</script>

<main class="h-[55vh] grid justify-center">
    <Card tw="place-self-end w-96">
        <div class="grid gap-6 px-2 pb-2">
            <Text tw="mb-2" b type="h2">Login/Register</Text>
            <Input bind:value={form.name} placeholder="Minecraft Name..." />
            <Input bind:value={form.password} password placeholder="Password..." />

            <Checkbox id="remember" bind:checked={form.remember}>
                <Text size="md">Remember me!</Text>
            </Checkbox>

            <div class="grid grid-cols-3">
                <Button href="/">Back</Button>
                <Button tw="col-start-3" type="primary">Continue</Button>
            </div>
        </div>
    </Card>
</main>