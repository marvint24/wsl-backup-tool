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
  color: var(--text-dim);
  font-family: var(--font-mono);
  font-size: 12.5px;
  line-height: 1.5;
  height: 100%;
  overflow-wrap: break-word;
  padding: 6px 0;
}
p{
  text-align: left;
  padding: 1px 10px;
  margin: 0;
}
p:hover{
  background: rgba(148, 163, 184, 0.06);
  color: var(--text);
}

</style>