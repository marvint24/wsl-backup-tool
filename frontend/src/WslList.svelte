<script lang="ts">
    import { distros } from "./store" // selectedDistro,
    import DistroRow from "./DistroRow.svelte"


</script>

<section>
  <table>
    <thead>
      <tr>
        <td>Default</td>
        <td>Name</td>
        <td>Status</td>
        <td>WSL-Version</td>
        <td></td>
      </tr>
    </thead>
    <tbody>
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
    </tbody>
  </table>
</section>

<style>
tbody{
  max-height: 400px;
  overflow: auto;
}
thead, tbody { 
  display: block;
}
table{
    width: 95%;
    margin-left: auto;
    margin-right: auto;
    border-spacing: 0 10px; 
}
thead>tr>td{
  width: 20%;
}
td{
    text-align: left;
    color: var(--white);
    font-size: 20px;
    width: 20%
}
tr td:first-child , tr td:nth-child(4){
    text-align: center;
}
</style>