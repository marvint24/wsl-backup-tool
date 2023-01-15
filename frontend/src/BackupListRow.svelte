<script lang="ts">
import {refresh,selectedDistro,backupRenameWindow} from './store'
import {RenameBackupFile,RestoreDistro} from '../wailsjs/go/main/App.js'
import cursor from './assets/cursor.svg'
import restore from './assets/restore.svg'


export var backupFile

// function refreshDistos(){
//     let i=1
//     let interval=setInterval(()=>{
//         if(i==20){
//             clearInterval(interval)
//         }
//         $refresh=true
//         i++
//     },1500)
// }
 




function backupRename(){
    $backupRenameWindow={"name":backupFile.Name}
}

function restoreBackup(){

}

    
</script>



<tr>
  <td>{backupFile.Name}</td>
  <td>{backupFile.ModDate}</td>
  <td>
    <div class="btns">
        <img title="Rename backup" on:click={backupRename} on:keydown src="{cursor}" alt="terminal">
        <img title="Restore backup" on:click={restoreBackup} on:keydown src="{restore}" alt="backup">
    </div>
  </td>
</tr>



<style>
td>div>img{
    cursor: pointer;
}
.btns{
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
    background-color: var(--dark);
    border-radius: 50px;
}
td{
    text-align: left;
    color: var(--white);
    font-size: 20px;
    height: 44px;
    padding-left: 25px;
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