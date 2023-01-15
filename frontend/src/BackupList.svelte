<script lang="ts">
  import {GetBackupFiles} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro} from './store'
  import BackupListRow from './BackupListRow.svelte'

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
    })
  }
 
  listBackupFiles()
</script>
  
  
  
<section id="window">
  <div>
    <div id="box">
      <div id="heading">Manage backup TAR-files for {$selectedDistro}</div>
      <hr/>
      <section>
        <table>
            <tr>
                <td>Name</td>
                <td>Created</td>
                <td></td>
            </tr>
            {#each backupFiles as item (item.uuid)}
                <BackupListRow backupFile={item}/>
            {/each}
        </table>
      </section>

    </div>
    <div id="row2">
      <div id="mbtn" on:click="{close}" on:keydown>Close</div>
    </div>
  </div>
</section>


<style>
table{
  width: 95%;
  margin-left: auto;
  margin-right: auto;
  border-spacing: 0 10px;
}
td{
    text-align: left;
    color: var(--dark);
    font-size: 20px;
    padding-left: 25px;
}
#mbtn{
  font-size: 20px;
  background-color: var(--green);
  border-radius: 10px;
  padding: 4px 8px 4px 8px;
  margin: 0 10px 0 0;
  cursor: pointer;
}
#row2{
  position: relative;
  top: 10px;
  left: 91%;
  display: flex;
  width: 0%;
  flex-direction: row;
}
#box{
  background-color: var(--white);
  border-radius: 15px;
  width: 95vw;
  height: 89vh;
}
#heading{
  color: var(--dark);
  font-size: 20px;
  font-weight: 500;
  padding: 5px 47% 0 20px; 
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
  padding: 5px 78% 0 10px;  
  margin: 0 0 0 15px;
}
#window {
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