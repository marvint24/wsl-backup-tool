<script lang="ts">
    import { distros } from "./store" // selectedDistro,
    import DistroRow from "./DistroRow.svelte"


</script>

<section>
    <table>
        <tr>
            <td>Default</td>
            <td>Name</td>
            <td>Status</td>
            <td>WSL-Version</td>
            <td></td>
        </tr>
        {#each $distros as distroRow (distroRow.uuid)}
            <DistroRow distroRow={distroRow}/>
        {/each}
    </table>
</section>

<style>
table{
    width: 95%;
    margin-left: auto;
    margin-right: auto;
    border-spacing: 0 10px;
}
td{
    text-align: left;
    color: var(--white);
    font-size: 20px;
}
tr td:first-child , tr td:nth-child(4){
    text-align: center;
}
</style>