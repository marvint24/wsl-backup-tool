<script lang="ts">
  import {shutdownWindow} from './store'
  import {EventsEmit} from '../wailsjs/runtime/runtime.js'
  import Modal from './Modal.svelte'

  function close(){
    EventsEmit("WslShutdownCanceled")
    $shutdownWindow=null
  }

  function shutdown(){
    EventsEmit("WslShutdownConfirmed")
    $shutdownWindow=null
  }
  
</script>
  
  
  
<Modal heading="WSL Shutdown needed!">
  <div slot="content">
    <div class="text">To continue the current action, WSL needs to be shutdown.<br>Do you want to continue? (Potential data loss)</div>
  </div>
  <div slot="actions">
    <div class="mbtn" on:click="{shutdown}" on:keydown>OK</div>
    <div class="mbtn" on:click="{close}" on:keydown>Cancel</div>
  </div>
</Modal>


<style>
  .mbtn{
    font-size: 20px;
    background-color: var(--green);
    border-radius: 10px;
    padding: 4px 8px 4px 8px;
    margin: 0 10px 0 0;
    cursor: pointer;
  }
  .text{
    color: var(--dark);
    font-size: 20px;
    padding: 5px 25px 25px 25px; 
    text-align: left; 
  }
</style>