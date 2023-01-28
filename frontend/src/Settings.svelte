<script lang="ts">
  import folder from "./assets/folder.svg";
  import {GetSettings,SetSettings,SelectFolder,TestPath} from '../wailsjs/go/main/App.js'
  import {selectedWindow} from './store'

  var settings={
    BackupPath:"",
    RefreshInterval:2
  }

  GetSettings().then((result)=>{
    settings=JSON.parse(result)
    checkPath()
  })

  async function checkPath(){
    if(settings.BackupPath.endsWith("\\")){
      settings.BackupPath=settings.BackupPath.slice(0,-1)
    }
    let el=<HTMLInputElement>document.getElementById("BackupPath-Setting")
    await TestPath(settings.BackupPath).then((res)=>{
      if(res){
        el.setCustomValidity("")
      }else{
        el.setCustomValidity("Path not found")
      }
    })
  }

  function openFolder(){
    SelectFolder().then((result)=>{
      if(result){
        settings.BackupPath=result
      }  
    })
  }



  function close(){
    $selectedWindow=null
  }

  async function okay(){
    console.log(settings)
    if((<HTMLFormElement>document.getElementById("SettingsForm")).checkValidity()){
      SetSettings(JSON.stringify(settings))
      close()
    }
  }

</script>

<section>
  <div>
    <div class="box">
      <form id="SettingsForm">
        <div class="heading">Settings</div>
        <hr/>
        <div class="text">Backup folder</div>
        <div class="row"><input id="BackupPath-Setting" type="text" bind:value={settings.BackupPath} on:change={checkPath}><div class="btn" title="Open folder" on:click="{openFolder}" on:keydown><img src="{folder}" alt="folder"></div></div>
        <div class="text">Refresh interval</div>
        <div class="row"><input id="RefreshInterval-Setting" title="in seconds" min="1" type="number" bind:value={settings.RefreshInterval}></div>
      </form>
      </div>
    <div class="row2">
      <button class="mbtn" on:click="{okay}" on:keydown>OK</button>
      <button class="mbtn" on:click="{close}" on:keydown>Cancel</button>
    </div>
  </div>
</section>

<style>
.mbtn{
  font-size: 20px;
  color: var(--white);
  border-width: 0;
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
input:invalid{
  border: 2px solid red;
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