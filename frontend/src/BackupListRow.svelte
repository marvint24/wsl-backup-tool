<script lang="ts">
import {selectedDistro,backupRenameWindow,selectedWindow} from './store'
import {RestoreDistro} from '../wailsjs/go/main/App.js'
import cursor from './assets/cursor.svg'
import restore from './assets/restore.svg'

export var backupFile

function close(){
    $selectedWindow=null
}


function backupRename(){
    $backupRenameWindow={"name":backupFile.Name}
}

function restoreBackup(){
    RestoreDistro(backupFile.Name,$selectedDistro).then()
    close()
}

    
</script>



<tr class="row">
  <td class="col-name">{backupFile.Name}</td>
  <td class="col-date">{backupFile.ModDate}</td>
  <td>
    <div class="btns">
        <button class="icon-btn" title="Rename backup" on:click={backupRename}><img src="{cursor}" alt="rename"></button>
        <button class="icon-btn" title="Restore backup" on:click={restoreBackup}><img src="{restore}" alt="restore"></button>
    </div>
  </td>
</tr>



<style>
.row td {
    background-color: var(--surface-2);
    color: var(--text);
    font-size: 14px;
    height: 48px;
    text-align: left;
    padding-left: 16px;
    border-top: 1px solid var(--border);
    border-bottom: 1px solid var(--border);
    transition: background-color 0.12s ease;
}
.row:hover td {
    background-color: var(--surface-3);
}
.col-name {
    font-weight: 600;
}
.col-date {
    color: var(--text-dim);
    white-space: nowrap;
}
td:first-child {
    border-top-left-radius: var(--radius);
    border-bottom-left-radius: var(--radius);
    border-left: 1px solid var(--border);
}
td:last-child {
    border-top-right-radius: var(--radius);
    border-bottom-right-radius: var(--radius);
    border-right: 1px solid var(--border);
    padding-left: 0;
}
.btns {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 4px;
    padding-right: 10px;
    user-select: none;
}
</style>