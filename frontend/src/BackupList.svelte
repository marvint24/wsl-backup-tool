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
  
  
  
<section class="window">
  {#if $backupRenameWindow.name != null}
    <BackupRename/>
  {/if}
  <div>
    <div class="box">
      <div class="heading">Manage backup files for {$selectedDistro}</div>
      <hr/>
      <section>
        {#if backupFiles.length!=0}
        <div class="scroll">
          <table>
              <tr>
                <td>Name</td>
                <td>Created</td>
                <td><div class="btn" title="Open folder" on:click="{openFolder}" on:keydown><img src="{folder}" alt="folder"></div></td>
              </tr>
                {#each backupFiles as item (item.uuid)}
                  <BackupListRow backupFile={item}/>
                {/each}
          </table>
        </div>
        {:else}
          <p>There are no files {":("}</p>
        {/if}
      </section>

    </div>
    <div class="row2">
      <div class="mbtn" on:click="{close}" on:keydown>Close</div>
    </div>
  </div>
</section>


<style>
section {
  padding: 0 5px 0 5px;
}
.scroll{
  overflow: auto;
  max-height: calc(100vh - 152px);
}
::-webkit-scrollbar {
  width: 10px;
}
::-webkit-scrollbar-track {
  background: hwb(0 0% 100% / 0);
}
::-webkit-scrollbar-thumb {
  background: var(--dark);
  border-radius: 5px;
}

img{
  padding: 3px 0 0 0;
  height: 22px;
}
.btn{
  border-width: 1px;
  border: solid;
  border-color: var(--dark);
  background-color: white;
  border-radius: 8px;
  width: 34px;
  height: 28px;
  cursor: pointer;
}
p{
  color: var(--dark);
  font-size: 20px;
  font-weight: 500;
  white-space: nowrap 
}
table{
  width: 95%;
  margin-left: auto;
  margin-right: auto;
  border-spacing: 0 10px;
  margin-bottom: 20px;
}
td{
    text-align: left;
    color: var(--dark);
    font-size: 20px;
    padding-left: 25px;
}
td:nth-child(3){
  display: flex;
  justify-content: right;
  text-align: center;
}
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
  margin-left: auto; 
  margin-right: 80px;
  display: flex;
  width: 0%;
  flex-direction: row;
}
.box{
  background-color: var(--white);
  border-radius: 15px;
  width: 95vw;
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
.window {
  position: absolute;
  width: calc(100% - 10px);
  height: 100%;
  background-color:var(--dark-op);
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}
  </style>