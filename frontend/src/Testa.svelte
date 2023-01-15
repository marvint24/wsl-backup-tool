<script lang="ts">
  import {SetBackupPath} from '../wailsjs/go/main/App.js'


function joo():string{
  console.log("aa")
  SetBackupPath("aaaa").then()
  return "hi"
}

</script>
  
<main>
  <div on:click="{joo}" on:keypress>hii</div>
</main>

<style>


</style>