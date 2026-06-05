<script lang="ts">
  import {CreateBackupFile,GetSettings} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro} from './store'
  import {removeIllegalCharacters} from './utils'
  import Modal from './Modal.svelte'

  let backupPath: string 
  GetSettings().then((result)=>{
    backupPath=JSON.parse(result).replaceBackupPath
  })

  var backupFilename: string
  let now=new Date().toLocaleString()
  backupFilename=removeIllegalCharacters(`${now}.vhdx`)

  function close(){
    $selectedWindow=null
  }

  function createBackupFile(){
    backupFilename=removeIllegalCharacters(backupFilename)
    CreateBackupFile($selectedDistro,backupFilename,0).then()
    close()
  }
</script>
  
  
  
<Modal heading="Create backup file for {$selectedDistro}">
  <div slot="content">
    <div class="text">Select name</div>
    <input type="text" bind:value={backupFilename}>
  </div>
  <div slot="actions">
    <div class="mbtn" on:click="{createBackupFile}" on:keydown>OK</div>
    <div class="mbtn" on:click="{close}" on:keydown>Cancel</div>
  </div>
</Modal>


<style>
  .mbtn{
    font-size: 20px;
    background-color: var(--green);
    border-radius: 10px;
    padding: 4px 8px 4px 8px;
    margin: 0 10px 0 0;
    cursor: pointer;
  }
  .text{
    color: var(--dark);
    font-size: 20px;
    padding: 5px 0 0 25px; 
    text-align: left; 
  }
  input{
    font-size: 20px;
    margin: 0 20px 20px 20px;
    border-radius: 10px;
    padding: 2px 10px 2px 10px;
    min-width: 600px;
  }
</style>