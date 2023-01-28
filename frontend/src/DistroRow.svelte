<script lang="ts">
import check from './assets/check.svg'
import terminal from './assets/terminal.svg'
import backup from './assets/backup.svg'
import list from './assets/list.svg'
import stop from './assets/stop.svg'
import {selectedWindow,selectedDistro} from './store'
import {TerminateWsl,LaunchDistro} from '../wailsjs/go/main/App.js'
export var distroRow

function launchDistro(){
    LaunchDistro(distroRow.Name)
}

function createBackup(){
    $selectedDistro=distroRow.Name
    $selectedWindow="BackupCreate"
}

function manageBackups(){
    $selectedDistro=distroRow.Name
    $selectedWindow="BackupList"
}

function terminate(){
  TerminateWsl(distroRow.Name).then()
}
    
</script>



<tr>
    {#if distroRow.Default_}
        <td><img src="{check}" alt="check"></td>
    {:else}
        <td></td>
    {/if}
    <td>{distroRow.Name}</td>
    <td>{distroRow.Status}</td>
    <td>{distroRow.Wsl_version}</td>
    <td>
        <div class="btns">
            <img title="Open in terminal" on:click={launchDistro} on:keydown src="{terminal}" alt="terminal">
            <img title="Backup" on:click={createBackup} on:keydown src="{backup}" alt="backup">
            <img title="Manage backups" on:click={manageBackups} on:keydown src="{list}" alt="list">
            {#if distroRow.Status==="Running"}
                <img title="Terminate" on:click={terminate} on:keydown src="{stop}" alt="stop">
            {/if}
        </div>
    </td>
</tr>


<style>
td>div>img{
    cursor: pointer;
}
.btns{
  margin-left: 20px;
  user-select: none;
}
.btns>img{
  margin-left: 15px;
}
img{
    height: 25px;
    filter: invert(88%) sepia(8%) saturate(1480%) hue-rotate(178deg) brightness(106%) contrast(104%);
    padding-top: 5px;
}
tr{
    background-color: var(--gray);
    border-radius: 50px;
}
td{
    text-align: left;
    color: var(--white);
    font-size: 20px;
    height: 44px;
}
tr>td:first-child , tr>td:nth-child(4){
    text-align: center;
}
tr>td:last-child{
    border-left-width: 2px;
    border-right-width:0;
    border-top-width:0;
    border-bottom-width:0;
    border-width: 2px,0px,0px,0px; 
    border-color: var(--green-light);
    border-style: solid;
}
td:first-child {
    border-top-left-radius: 15px; 
    border-bottom-left-radius: 15px;
}
td:last-child {
    border-bottom-right-radius: 15px; 
    border-top-right-radius: 15px; 
}
</style>