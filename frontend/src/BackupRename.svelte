<script lang="ts">
  import {RenameBackupFile} from '../wailsjs/go/main/App.js'
  import {selectedDistro,backupRenameWindow} from './store'

  var oldName=$backupRenameWindow.name
  var backupFilename=$backupRenameWindow.name

  function close(){
    $backupRenameWindow.name=null
  }

  function removeIllegalCharacters(){
    backupFilename=backupFilename.replaceAll("<","")
    backupFilename=backupFilename.replaceAll(">","")
    backupFilename=backupFilename.replaceAll(":","_")
    backupFilename=backupFilename.replaceAll("\"","")
    backupFilename=backupFilename.replaceAll("/","")
    backupFilename=backupFilename.replaceAll("\\","")
    backupFilename=backupFilename.replaceAll("|","")
    backupFilename=backupFilename.replaceAll("?","")
    backupFilename=backupFilename.replaceAll("*","")
  }

  function renameBackupFile(){
    removeIllegalCharacters()
    RenameBackupFile(oldName,backupFilename,$selectedDistro).then(close)
  }
</script>
  
  
  
<section class="modal-overlay rename-overlay">
  <div class="modal">
    <div class="modal-header">
      <h2 class="modal-title">Rename backup for <strong>{$selectedDistro}</strong></h2>
    </div>
    <div class="modal-body">
      <label class="field-label" for="rename-name">New name</label>
      <input id="rename-name" class="input" type="text" bind:value={backupFilename}>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" on:click="{close}">Cancel</button>
      <button class="btn btn-primary" on:click="{renameBackupFile}">Rename</button>
    </div>
  </div>
</section>


<style>
  /* sits above the backup-list dialog */
  .rename-overlay {
    z-index: 20;
  }
</style>