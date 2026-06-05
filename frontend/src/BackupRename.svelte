<script lang="ts">
  import {RenameBackupFile} from '../wailsjs/go/main/App.js'
  import {selectedDistro,backupRenameWindow} from './store'
  import {removeIllegalCharacters} from './utils'
  import Modal from './Modal.svelte'

  var oldName=$backupRenameWindow.name
  var backupFilename=$backupRenameWindow.name

  function close(){
    $backupRenameWindow.name=null
  }

  function renameBackupFile(){
    backupFilename=removeIllegalCharacters(backupFilename)
    RenameBackupFile(oldName,backupFilename,$selectedDistro).then(close)
  }
</script>
  
  
  
<Modal heading="Rename backup file for {$selectedDistro}">
  <div slot="content">
    <div class="text">New name</div>
    <input type="text" bind:value={backupFilename}>
  </div>
  <div slot="actions">
    <div class="mbtn" on:click="{renameBackupFile}" on:keydown>OK</div>
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
    padding: 5px 0 0 25px; 
    text-align: left; 
  }
  input{
    font-size: 20px;
    margin: 0 20px 20px 20px;
    border-radius: 10px;
    padding: 2px 10px 2px 10px;
    min-width: 600px;
  }
</style>