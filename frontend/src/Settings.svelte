<script lang="ts">
  import folder from "./assets/folder.svg";
  import {GetSettings} from '../wailsjs/go/main/App.js'

  let backupPath: string 
  GetSettings().then((result)=>{
    backupPath=JSON.parse(result).BackupPath
  })

  function openFolder(){

  }

  function okay(){

  }

  function cancel(){
    
  }

</script>

<section>
  <div>
    <div id="box">
      <div id="heading">Settings</div>
      <hr/>
      <div class="text">Backup folder</div>
      <div id="row"><input type="text" bind:value={backupPath}><div id="btn" title="Open folder" on:click="{openFolder}" on:keydown><img src="{folder}" alt="folder"></div>
    </div>
    <div id="row2">
      <div id="mbtn" on:click="{okay}" on:keydown>OK</div>
      <div id="mbtn" on:click="{cancel}" on:keydown>Cancel</div>
    </div>
  </div>
</section>

<style>
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
  top: 40px;
  right: -550px;
  display: flex;
  width: 0%;
  flex-direction: row;
}
#box{
  background-color: var(--white);
  border-radius: 15px;
}
#heading{
  color: var(--dark);
  font-size: 20px;
  font-weight: 500;
  padding: 5px 84% 0 20px;  
}
#row{
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
#btn{
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
  margin: 0 0 0px 20px;
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