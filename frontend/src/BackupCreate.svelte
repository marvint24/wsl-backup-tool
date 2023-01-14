<script lang="ts">
  import {CreateBackupFile,GetSettings} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro} from './store'
 
  let backupPath: string 
  GetSettings().then((result)=>{
    backupPath=JSON.parse(result).replaceBackupPath
  })

  var backupFilename: string
  let now=new Date().toLocaleString()
  .replaceAll("<","")
  .replaceAll(">","")
  .replaceAll(":","_")
  .replaceAll("\"","")
  .replaceAll("/","")
  .replaceAll("\\","")
  .replaceAll("|","")
  .replaceAll("?","")
  .replaceAll("*","")
  backupFilename=`${now}.tar`

  function close(){
    $selectedWindow=null
  }

  function createBackupFile(){
    CreateBackupFile($selectedDistro,backupFilename).then()
    close()
  }
</script>
  
  
  
<section>
  <div>
    <div id="box">
      <div id="heading">Create backup TAR-file for {$selectedDistro}</div>
      <hr/>
      <div class="text">Select name</div>
      <input type="text" bind:value={backupFilename}>
    </div>
    <div id="row2">
      <div id="mbtn" on:click="{createBackupFile}" on:keydown>OK</div>
      <div id="mbtn" on:click="{close}" on:keydown>Cancel</div>
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
    top: 10px;
    right: -520px;
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