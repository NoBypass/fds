<script lang="ts">
    import Card from '$lib/components/Card.svelte'
    import Text from '$lib/components/Text.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Button from '$lib/components/Button.svelte'
    import { api } from '$lib/stores/api'
    import { alertStore } from '$lib/stores/alertStore'
    import Spinner from '$lib/components/Spinner.svelte'

    let lastNameChange = 0
    let lastName = ''
    let isLoading = false
    const nameRegex = /^(?![0-9])[a-zA-Z0-9_]{0,16}$/

    let { name, password, remember } = {
        name: {
            val: '',
            error: '',
            resolved: false,
            valid: false
        },
        password: {
            val: '',
            error: '',
            valid: false
        },
        remember: false
    }

    $: {
        name.valid = name.val.length > 0 && name.error === '' && name.resolved
        password.valid = password.val.length > 0 && password.error === ''
    }

    $: {
        if (name.val !== lastName) {
            lastNameChange = Date.now()
            lastName = name.val

            setTimeout(() => {
                if (lastNameChange + 500 < Date.now()) {
                    isLoading = true
                    $api.graphql.query({
                        query: `
                            query ($name: String!) {
                                user(name: $name) {
                                    name
                                }
                            }
                        `,
                        variables: {
                            name: name.val
                        }
                    }).then((res) => {
                        name.error = ''
                        isLoading = false
                        name.resolved = true
                        name.valid = true
                    }).catch((err) => {
                        let msg = err.message
                        if (err.message === 'Failed to fetch') {
                            msg = 'Could not connect to server, please reload the page and try again.'
                        }
                        name.error = msg
                        isLoading = false
                        name.resolved = true
                        alertStore.push(msg, 'danger')
                    })
                }
            }, 500)
        }
    }
</script>

<main class="h-[55vh] grid justify-center">
    <Card tw="place-self-end w-96">
        <div class="grid px-2 pb-2">
            <Text tw="mb-2" b type="h2">Login/Register</Text>
            <Input bind:value={name.val} color={name.error.length > 0 ? 'danger' : name.valid ? 'success' : 'neutral'} regex={nameRegex} placeholder="Minecraft Name">
                <div slot="right">
                    {#if isLoading}
                        <Spinner />
                    {/if}
                </div>
                <div slot="below">
                    {#if name.error.length > 0}
                        <Text size="sm" color="danger" b>{name.error}</Text>
                    {/if}
                </div>
            </Input>
            <Input bind:value={password.val} password placeholder="Password" />

            <Checkbox tw="mb-8 mt-4" id="remember" bind:checked={remember}>
                <Text size="md">Remember me!</Text>
            </Checkbox>

            <div class="grid grid-cols-3">
                <Button href="/">Back</Button>
                <Button disabled={!name.valid && !password.valid} tw="col-start-3" type="primary">Continue</Button>
            </div>
        </div>
    </Card>
</main>