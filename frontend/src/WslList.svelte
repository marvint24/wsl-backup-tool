<script lang="ts">
    import { distros } from "./store" // selectedDistro,
    import DistroRow from "./DistroRow.svelte"


</script>

<div>
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
</div>

<style>
section {
  overflow: auto;
  max-height: calc(100vh - 97px);
}
table{
  width: 97%;
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

::-webkit-scrollbar {
  width: 10px;
}

::-webkit-scrollbar-track {
  background: hwb(0 0% 100% / 0);
}

::-webkit-scrollbar-thumb {
  background: var(--white);
  border-radius: 5px;
}
div{
  padding: 0 8px 0 8px;
}
</style>