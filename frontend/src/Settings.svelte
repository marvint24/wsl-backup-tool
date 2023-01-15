<script lang="ts">
  import folder from "./assets/folder.svg";
  import {GetSettings,SetBackupPath,SelectFolder} from '../wailsjs/go/main/App.js'
  import {selectedWindow} from './store'

  let backupPath: string 
  GetSettings().then((result)=>{
    backupPath=JSON.parse(result).BackupPath
  })

  function openFolder(){
    SelectFolder().then((result)=>{
      if(result){
        backupPath=result
      }  
    })
  }

  function close(){
    $selectedWindow=null
  }

  function okay(){
    SetBackupPath(backupPath)
    close()
  }


</script>

<section>
  <div>
    <div class="box">
      <div class="heading">Settings</div>
      <hr/>
      <div class="text">Backup folder</div>
      <div class="row"><input type="text" bind:value={backupPath}><div class="btn" title="Open folder" on:click="{openFolder}" on:keydown><img src="{folder}" alt="folder"></div></div>
    </div>
    <div class="row2">
      <div class="mbtn" on:click="{okay}" on:keydown>OK</div>
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
  right: -560px;
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
  padding: 5px 0 0 20px; 
  text-align: left; 
}
.row{
  display: flex;
  flex-direction: row;
}
hr{
  border: none;
  border-top: 2px solid var(--dark);
  margin-bottom: 15px;
}
img{
  padding: 3px 0 0 0;
  height: 22px;
}
.btn{
  text-align: center;
  border-width: 1px;
  border: solid;
  border-color: var(--dark);
  background-color: white;
  border-radius: 8px;
  width: 34px;
  height: 28px;
  cursor: pointer;
  margin-right: 20px;
  margin-left: 3px;
}
input{
  font-size: 20px;
  margin: 0 0 20px 20px;
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