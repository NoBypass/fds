<script lang="ts">
    import Text from '$lib/components/Text.svelte'
    import Link from '$lib/components/Link.svelte'
    import Card from '$lib/components/Card.svelte'
</script>

<div class="flex gap-2 mb-16">
    <Card tw="rounded-2xl">Lines of Code: <Text type="s" b>32,694</Text></Card>
    <Card tw="rounded-2xl">Commits: <Text type="s" b>202</Text></Card>
    <Card tw="rounded-2xl">Contributors: <Text type="s" b>2</Text></Card>
    <Card tw="rounded-2xl">Version: <Text type="s" b>0.0.0</Text></Card>
    <Card tw="rounded-2xl">Last Release: <Text type="s" b>never</Text></Card>
</div>

<Text type="h2" b>How can I contribute?</Text>
<br>
<Text>
    Contributing can be everything from reporting a bug over suggesting a feature to writing code. If you want to do any
    of those, you can do that on the Project's <Link href="https://github.com/NoBypass/fds">GitHub Page</Link>.
    If you want to suggest a feature or create a bug report, you can do that by creating an issue
    <Link href="https://github.com/NoBypass/fds/issues/new/choose">here</Link>. If you're interested in contributing by
    writing code, please read the <Link>How To Contribute Guide</Link>.
</Text>

<Text tw="mt-16" type="h2" b>Is it save to use the app?</Text>
<br>
<Text>
    The short answer is: Yes absolutely. But you don't have to take my word for it, all the code for the app, server
    and even the website are all open-source which means that the code can be viewed by everybody. If you yourself
    don't understand the code and you want to be sure, you can view the "Discussions" tab or simply ask somebody who
    knows it! The file you download to install the app is automatically generated from that code by GitHub, which
    means you're guaranteed to download the executable from that code and not just some random installer.
</Text>




