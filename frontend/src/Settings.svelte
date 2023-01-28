<script lang="ts">
  import folder from "./assets/folder.svg";
  import {SetSettings,SelectFolder,TestPath} from '../wailsjs/go/main/App.js'
  import {selectedWindow,settings} from './store'
  import {EventsEmit} from '../wailsjs/runtime/runtime.js'

  let tempSettings=JSON.parse(JSON.stringify($settings))

  async function checkPath(){
    if(tempSettings.BackupPath.endsWith("\\")){
      tempSettings.BackupPath=tempSettings.BackupPath.slice(0,-1)
    }
    let el=<HTMLInputElement>document.getElementById("BackupPath-Setting")
    await TestPath(tempSettings.BackupPath).then((res)=>{
      if(res){
        el.setCustomValidity("")
      }else{
        el.setCustomValidity("Path not found")
      }
    })
  }

  window.onload=()=>{
    checkPath()
  }

  function openFolder(){
    SelectFolder().then((result)=>{
      if(result){
        tempSettings.BackupPath=result
        checkPath()
      }  
    })
  }

  function close(){
    $selectedWindow=null
  }

  function okay(){
    if((<HTMLFormElement>document.getElementById("SettingsForm")).checkValidity()){
      $settings=tempSettings
      SetSettings(JSON.stringify($settings)).then(()=>{
        EventsEmit("SettingsChanged")
        close()
      })
    }
  }

</script>

<section>
  <div>
    <div class="box">
      <div class="heading">Settings</div>
      <hr/>
      <form id="SettingsForm">
        
        <div class="text">Backup folder</div>
        <div class="row"><input id="BackupPath-Setting" class="long" type="text" bind:value={tempSettings.BackupPath} on:change={checkPath}><div class="btn" title="Open folder" on:click="{openFolder}" on:keydown><img src="{folder}" alt="folder"></div></div>
        <div class="text">Refresh interval</div>
        <div class="row"><input title="in seconds" min="1" type="number" bind:value={tempSettings.RefreshInterval}></div>
        <div class="row"><div class="text">Show console</div><input id="ShowConsole-Setting" type="checkbox" bind:checked={tempSettings.ShowConsole}></div>
      </form>
      </div>
    <div class="row2">
      <button class="mbtn" on:click="{okay}" on:keydown>OK</button>
      <button class="mbtn" on:click="{close}" on:keydown>Cancel</button>
    </div>
  </div>
</section>

<style>
#ShowConsole-Setting{
  min-width: 20px;
  height: 20px;
  margin: 4.5px 0 10px 10px;
}
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
.long{
  min-width: 600px;
}
form{
  margin: 0 0 20px 20px;
}
input{
  font-size: 20px;
  border-radius: 10px;
  padding: 2px 10px 2px 10px;
  margin-bottom: 5px;
}
input:invalid{
  border: 2px solid red;
}
.text{
  color: var(--dark);
  font-size: 20px;
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