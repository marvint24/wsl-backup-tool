<script lang="ts">
  import {RenameBackupFile} from '../wailsjs/go/main/App.js'
  import {selectedDistro,backupRenameWindow} from './store'

  var backupFilename=$backupRenameWindow.name

  function close(){
    $backupRenameWindow.name=null
  }

  function renameBackupFile(){
    // RenameBackupFile()
  }
</script>
  
  
  
<section>
  <div>
    <div class="box">
      <div class="heading">Rename backup file for {$selectedDistro}</div>
      <hr/>
      <div class="text">New name</div>
      <input type="text" bind:value={backupFilename}>
    </div>
    <div class="row2">
      <div class="mbtn" on:click="{renameBackupFile}" on:keydown>OK</div>
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
    right: -520px;
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
    text-align: left;
    white-space: nowrap
  }
  hr{
    border: none;
    border-top: 2px solid var(--dark);
    margin-bottom: 15px;
  }
  input{
    font-size: 20px;
    margin: 0 20px 20px 20px;
    border-radius: 10px;
    padding: 2px 10px 2px 10px;
    min-width: 600px;
  }
  .text{
    color: var(--dark);
    font-size: 20px;
    padding: 5px 0 0 25px; 
    text-align: left; 
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