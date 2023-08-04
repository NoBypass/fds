<script>
    import loginBg from '$lib/assets/login-bg.png'
    import Button from '$lib/components/Button.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Modal from '$lib/components/Modal.svelte'
    import Text from '$lib/components/Text.svelte'
    import { createEventDispatcher } from 'svelte'
    import Spinner from '$lib/components/Spinner.svelte'

    export let open = false

    let isValid = false
    const errors = {
        password: '',
        username: ''
    }
    const pwdChecks = [
        {
            name: 'Password must be between 1 and 32 characters long',
            check: () => form.password.length > 0 && form.password.length <= 32
        },
        {
            name: 'Password must contain at least one uppercase letter',
            check: () => /[A-Z]/.test(form.password)
        },
        {
            name: 'Password must contain at least one lowercase letter',
            check: () => /[a-z]/.test(form.password)
        },
        {
            name: 'Password must contain at least one number',
            check: () => /[0-9]/.test(form.password)
        },
        {
            name: 'Password must contain at least one special character',
            check: () => /[!@#$%^&*()\-_=+{}[\]:;"'<>,.?/\\|~`]/.test(form.password)
        },
        {
            name: 'Password must not contain any spaces',
            check: () => !form.password.includes(' ')
        }
    ]
    const usrChecks = [
        {
            name: 'Username must be between 1 and 16 characters long',
            check: () => form.username.length > 0 && form.username.length <= 16
        },
        {
            name: 'Username must only contain letters, numbers, and underscores',
            check: () => /^[a-zA-Z0-9_]+$/.test(form.username)
        }
    ]
    const form = {
        username: '',
        password: '',
        remember: false
    }

    $: {
        isValid = pwdChecks.every(check => check.check()) && usrChecks.every(check => check.check())
        errors.password = pwdChecks.filter(check => !check.check())[0]?.name
        errors.username = usrChecks.filter(check => !check.check())[0]?.name
        if (form.password.length === 0) errors.password = ''
        if (form.username.length === 0) errors.username = ''
        if (!errors.password) errors.password = ''
        if (!errors.username) errors.username = ''
    }

    const handleUsername = (e) => {
        form.username = e.detail.target.value
    }

    const handlePassword = (e) => {
        form.password = e.detail.target.value
    }

    const handleRemember = (e) => {
        form.remember = e.detail.target.checked
    }

    const dispatch = createEventDispatcher()
    const close = () => {
        open = false
        dispatch('close')
    }
</script>

<Modal open={open} on:close={close} tw="w-2/3 h-160 flex space-between">
    <img src={loginBg}
         alt="background"
         class="w-1/2 h-full object-none z-20"
         style="-webkit-mask-image:-webkit-gradient(linear, left bottom, right bottom, from(rgba(0,0,0,1)), to(rgba(0,0,0,0)));
                mask-image: linear-gradient(to right, rgba(0,0,0,1), rgba(0,0,0,0));">
    <div class="w-full">
        <div class="p-12 z-30 relative w-full h-full grid grid-cols-1 grid-rows-3">
            <div>
                <Text type="h1" size="xl" b>Login/Register</Text>
                <Text o tw="mt-4">You can both register a new account and log into an existing account via the same form.</Text>
            </div>

            <div>
                <Input color={errors.username.length !== 0 ? 'error' : form.username.length === 0 ? 'neutral' : 'success'}
                       on:change={handleUsername}
                       light rounded placeholder="Minecraft username" tw="mt-12 w-full">
                    <Spinner slot="right" tw="hidden" />
                </Input>
                <Text tw="mt-2" color="error" b>{errors.username}</Text>

                <Input color={errors.password.length !== 0 ? 'error' : form.password.length === 0 ? 'neutral' : 'success'}
                       on:change={handlePassword} light rounded placeholder="Password" tw="mt-10 w-full" />
                <Text tw="mt-2" color="error" b>{errors.password}</Text>

                <Checkbox on:change={handleRemember} tw="mt-10"><Text size="md">Remember me!</Text></Checkbox>
            </div>

            <div class="grid w-full grid-cols-2 grid-rows-1 gap-32 mt-12 h-10 place-self-end">
                <Button on:click={close} rounded color="transparent"><Text b>Cancel</Text></Button>
                <Button disabled={!isValid} rounded color="neutral"><Text b>Continue</Text></Button>
            </div>
        </div>
        <img src={loginBg} alt="background" class="saturate-150 absolute w-full h-full object-none blur-[100px]  z-10 left-0 top-0">
    </div>
</Modal>