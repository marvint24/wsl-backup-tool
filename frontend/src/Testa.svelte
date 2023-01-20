<script lang="ts">
  import {RestoreDistro} from '../wailsjs/go/main/App.js'


function joo():string{
  console.log("aa")
  RestoreDistro("test2.vhdx","Ubuntu").then()
  return "hi"
}

</script>
  
<main>
  <div on:click="{joo}" on:keypress>hii</div>
</main>

<style>


</style>