<script lang="ts">
  import {GetBackupFiles,OpenBackupFolder} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro,backupRenameWindow} from './store'
  import BackupListRow from './BackupListRow.svelte'
  import BackupRename from './BackupRename.svelte'
  import folder from "./assets/folder.svg";

  function close(){
    $selectedWindow=null
  }

  var backupFiles=[]
  function listBackupFiles(){
    GetBackupFiles($selectedDistro).then((result)=>{
      let jsonObj=JSON.parse(result)
      backupFiles=jsonObj.map((obj)=>{
        obj.uuid=crypto.randomUUID()
        return obj
      }) 
      backupFiles=backupFiles.sort((a,b)=>{
        if(a.ModDateInt>b.ModDateInt){
          return -1
        }else{
          return 1
        }
      }) 
    })
  }

  $:{
    if($backupRenameWindow.name==null){
      listBackupFiles()
    }
  }

  function openFolder(){
    OpenBackupFolder($selectedDistro)
  }

</script>
  
  
  
<section class="modal-overlay">
  {#if $backupRenameWindow.name != null}
    <BackupRename/>
  {/if}
  <div class="modal modal--wide">
    <div class="modal-header header-row">
      <h2 class="modal-title">Manage backups for <strong>{$selectedDistro}</strong></h2>
      <button class="icon-btn" title="Open folder" on:click="{openFolder}"><img src="{folder}" alt="folder"></button>
    </div>
    <div class="modal-body">
      {#if backupFiles.length!=0}
        <table class="backup-table">
            <thead>
              <tr class="col-head">
                <th>Name</th>
                <th>Created</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {#each backupFiles as item (item.uuid)}
                <BackupListRow backupFile={item}/>
              {/each}
            </tbody>
        </table>
      {:else}
        <p class="empty">No backup files yet.</p>
      {/if}
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" on:click="{close}">Close</button>
    </div>
  </div>
</section>


<style>
.header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
}
.empty {
  color: var(--text-dim);
  font-size: 15px;
  text-align: center;
  padding: 28px 0;
  margin: 0;
}
.backup-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 8px;
}
.col-head th {
  text-align: left;
  color: var(--text-faint);
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  padding: 0 16px 2px 16px;
}
</style>