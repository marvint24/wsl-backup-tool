<script lang="ts">
  import {shutdownWindow} from './store'
  import {EventsEmit} from '../wailsjs/runtime/runtime.js'

  function close(){
    EventsEmit("WslShutdownCanceled")
    $shutdownWindow=null
  }

  function shutdown(){
    EventsEmit("WslShutdownConfirmed")
    $shutdownWindow=null
  }
  
</script>
  
  
  
<section class="modal-overlay shutdown-overlay">
  <div class="modal">
    <div class="modal-header">
      <h2 class="modal-title">WSL shutdown required</h2>
    </div>
    <div class="modal-body">
      <p class="text">To continue the current action, WSL needs to be shut down.</p>
      <p class="warning">⚠ All running distributions will stop &mdash; unsaved work may be lost.</p>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" on:click="{close}">Cancel</button>
      <button class="btn btn-danger" on:click="{shutdown}">Shut down WSL</button>
    </div>
  </div>
</section>


<style>
  .shutdown-overlay {
    z-index: 30;
  }
  .text {
    color: var(--text);
    font-size: 15px;
    margin: 0 0 12px 0;
    line-height: 1.5;
  }
  .warning {
    color: var(--danger);
    background: var(--danger-soft);
    border: 1px solid rgba(248, 113, 113, 0.25);
    border-radius: var(--radius-sm);
    padding: 11px 14px;
    font-size: 13.5px;
    font-weight: 500;
    margin: 0;
    line-height: 1.45;
  }
</style>