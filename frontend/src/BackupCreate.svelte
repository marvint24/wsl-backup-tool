<script lang="ts">
  import {CreateBackupFile,GetSettings} from '../wailsjs/go/main/App.js'
  import {selectedWindow,selectedDistro,refresh} from './store'

  function refreshDistos(){
    let i=1
    let interval=setInterval(()=>{
        if(i==20){
            clearInterval(interval)
        }
        $refresh=true
        i++
    },1500)
}
 
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
  backupFilename=`${now}.vhdx`

  function close(){
    $selectedWindow=null
  }

  function createBackupFile(){
    CreateBackupFile($selectedDistro,backupFilename).then()
    refreshDistos()
    close()
  }
</script>
  
  
  
<section>
  <div>
    <div class="box">
      <div class="heading">Create backup file for {$selectedDistro}</div>
      <hr/>
      <div class="text">Select name</div>
      <input type="text" bind:value={backupFilename}>
    </div>
    <div class="row2">
      <div class="mbtn" on:click="{createBackupFile}" on:keydown>OK</div>
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
    right: -520px;
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
    padding: 5px 0 0 25px;  
    white-space: nowrap;
    text-align: left;
  }
  .text{
    color: var(--dark);
    font-size: 20px;
    padding: 5px 0 0 25px; 
    text-align: left; 
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