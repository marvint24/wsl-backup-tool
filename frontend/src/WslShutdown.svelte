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
  
  
  
<section>
  <div>
    <div class="box">
      <div class="heading">WSL Shutdown needed!</div>
      <hr/>
      <div class="text">To continue the current action, WSL needs to be shutdown.<br>Do you want to continue? (Potential data loss)</div>
    </div>
    <div class="row2">
      <div class="mbtn" on:click="{shutdown}" on:keydown>OK</div>
      <div class="mbtn" on:click="{close}" on:keydown>Cancel</div>
    </div>
  </div>
</section>


<style>
  .mbtn{
    font-size: 20px;
    background-color: var(--green);
    border-radius: 10px;
    padding: 4px 8px 4px 8px;
    margin: 0 10px 0 0;
    cursor: pointer;
  }
  .row2{
    position: relative;
    top: 10px;
    right: -420px;
    display: flex;
    width: 0%;
    flex-direction: row;
  }
  .box{
    background-color: var(--white);
    border-radius: 15px;
  }
  .heading{
    color: var(--dark);
    font-size: 20px;
    font-weight: 500;
    padding: 5px 0 0 25px;  
    white-space: nowrap;
    text-align: left;
  }
  .text{
    color: var(--dark);
    font-size: 20px;
    padding: 5px 25px 25px 25px; 
    text-align: left; 
  }
  hr{
    border: none;
    border-top: 2px solid var(--dark);
    margin-bottom: 15px;
  }
  section {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color:var(--dark-op);
    z-index: 1;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  </style>