<script lang="ts">
  import {SetBackupPath,LaunchDistro} from '../wailsjs/go/main/App.js'


function joo(){
  console.log("hi")
  LaunchDistro(0,"Ubuntu").then()
}

</script>
  
<main>
  <div on:click="{joo}" on:keypress>hii</div>
</main>

<style>


</style>
