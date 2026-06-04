<script lang="ts">
  import {CreateBackupFile,GetSettings} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro} from './store'

  let backupPath: string 
  GetSettings().then((result)=>{
    backupPath=JSON.parse(result).replaceBackupPath
  })

  var backupFilename: string
  let now=new Date().toLocaleString()
  backupFilename=`${now}.vhdx`
  removeIllegalCharacters()

  function close(){
    $selectedWindow=null
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

  function createBackupFile(){
    removeIllegalCharacters()
    CreateBackupFile($selectedDistro,backupFilename,0).then()
    close()
  }
</script>
  
  
  
<section class="modal-overlay">
  <div class="modal">
    <div class="modal-header">
      <h2 class="modal-title">Create backup for <strong>{$selectedDistro}</strong></h2>
    </div>
    <div class="modal-body">
      <label class="field-label" for="backup-name">Backup name</label>
      <input id="backup-name" class="input" type="text" bind:value={backupFilename}>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" on:click="{close}">Cancel</button>
      <button class="btn btn-primary" on:click="{createBackupFile}">Create backup</button>
    </div>
  </div>
</section>