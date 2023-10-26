<script lang="ts">
    import Table from '$lib/components/Table.svelte'
</script>

<Table columns={['a', 'b', 'c']} data={[{
    a: 'a2',
    b: 'b3',
    c: 'c1'
},
{
    a: 'a1',
    b: 'b2',
    c: 'c3'
},
{
    a: 'a3',
    b: 'b1',
    c: 'c2'
}
]} />