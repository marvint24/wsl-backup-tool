<script lang="ts">
  import {GetWslList} from '../wailsjs/go/main/App.js'
  import WslList from './WslList.svelte';

  function getWslDistros(): void {
    GetWslList().then((result)=> {
      console.log(result)
    })
  }

  getWslDistros()
</script>
  <WslList/>
<main>

</main>

<style>



</style>
