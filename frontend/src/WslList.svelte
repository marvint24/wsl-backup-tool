<script lang="ts">
    import { distros } from "./store" // selectedDistro,
    import DistroRow from "./DistroRow.svelte"


</script>

<section class="list-scroll">
  <table>
      <thead>
        <tr class="col-head">
            <th>Default</th>
            <th>Name</th>
            <th>Status</th>
            <th>WSL Version</th>
            <th></th>
        </tr>
      </thead>
      <tbody>
        {#each $distros as distroRow (distroRow.uuid)}
            <DistroRow distroRow={distroRow}/>
        {/each}
      </tbody>
  </table>
</section>

<style>
.list-scroll {
  flex: 1 1 auto;
  overflow: auto;
  padding: 12px 16px;
  min-height: 0;
}
table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 8px;
}
.col-head th {
  text-align: left;
  color: var(--text-faint);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  padding: 0 18px 4px 18px;
}
.col-head th:first-child,
.col-head th:nth-child(4) {
  text-align: center;
}
</style>