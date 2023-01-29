<script lang="ts">
  import {EventsOn} from '../wailsjs/runtime/runtime.js'

  let logs: string[]
  logs = []
  EventsOn("consoleLog",(data)=>{
      console.log(logs)
      logs.push(data)
      logs=logs
  })
</script>

<section>
  {#each logs as log,index (index)}
    <p>{log}</p>
  {/each}
</section>


<style>
section{
  color: white;
  font-family: Consolas;
  font-size: 14px;
  height: 100%;
  overflow-wrap: break-word;
}
p{
  position: left;
  text-align: left;
  padding: 0 5px 0 5px;
  margin: 0 0 0 0;
}

</style>