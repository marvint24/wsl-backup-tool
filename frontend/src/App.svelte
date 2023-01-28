<script lang="ts">
  import {EventsOn} from '../wailsjs/runtime/runtime.js'
  import WslList from './WslList.svelte'
  import {distros,selectedWindow} from './store'
  import sliders from './assets/sliders.svg'

  // import Testa from './Testa.svelte'

  import BackupCreate from './BackupCreate.svelte'
  import BackupList from './BackupList.svelte'
  import Settings from './Settings.svelte'

  const imported = {
    BackupCreate,
    BackupList,
    Settings
  }

  function openSettings(): void {
    $selectedWindow="Settings"
  }

  //Global event listeners
  EventsOn("openSettings",()=>{
    openSettings()
  })

  EventsOn("refreshWSL",(result)=>{
    let jsonObj=JSON.parse(result)
    $distros=jsonObj.map((obj)=>{
      obj.uuid=crypto.randomUUID()
      return obj
    })      
  })

</script>
  
<main>
  {#if selectedWindow}
    <svelte:component this={imported[$selectedWindow]} />
  {/if}
  <section class="header">
    <div class="Settings" title="Settings" on:click="{openSettings}" on:keydown><img src="{sliders}" alt="sliders"><p>Settings</p>  </div>
  </section>
  <WslList/>
  <!-- <Testa/> -->
</main>

<style>
.header{
  margin-bottom: 50px;
  display: flex;
  justify-content: space-between;
}
img{
    height: 25px;
    filter: invert(88%) sepia(8%) saturate(1480%) hue-rotate(178deg) brightness(106%) contrast(104%);
}
.Settings{
  border-radius: 0 0 15px 0;
}
.Settings>p{
  color: var(--white);
  font-size: 20px;
  margin: 10px 0 10px 5px;
  font-weight: 500;
}
.header>div{
  display: flex;
  flex-direction: row;
  align-items: center;;
  padding-left: 15px;
  padding-right: 15px;
  background-color: var(--green-light);
}
.header>div:hover{
  cursor: pointer;
}
</style>
