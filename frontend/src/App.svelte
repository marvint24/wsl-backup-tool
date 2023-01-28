<script lang="ts">
  import {EventsOn} from '../wailsjs/runtime/runtime.js'
  import WslList from './WslList.svelte'
  import {distros,selectedWindow,settings} from './store'
  import sliders from './assets/sliders.svg'
  import Console from './Console.svelte'
  import {GetSettings} from '../wailsjs/go/main/App.js'

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

  GetSettings().then((result)=>{
    $settings=JSON.parse(result)
  })

  //Console resizer
  let resizer
  let element
  function initResize() {
    window.addEventListener('mousemove', Resize, false)
    window.addEventListener('mouseup', stopResize, false)
  }
  function Resize(e) {
    let bottomPos=window.innerHeight-e.clientY
    element.style.height = bottomPos + 'px'
  }
  function stopResize() {
    window.removeEventListener('mousemove', Resize, false)
    window.removeEventListener('mouseup', stopResize, false)
  }
  
  window.onload=()=>{
    resizer=document.getElementById("console-resizer")
    element=document.getElementById("console")
    resizer.addEventListener('mousedown', initResize, false)
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
    <div class="Settings" title="Settings" on:click="{openSettings}" on:keydown><img src="{sliders}" alt="sliders"><p>Settings</p></div>
  </section>
  <WslList/>
  <section class="console-section">
    <div id="console-resizer"></div>
    <div id="console"><Console/></div>
  </section>
</main>

<style>

/* Console */
.console-section{
  position: absolute;
  bottom: 0px;
  left: 0;
  width: 100%;
}
#console-resizer{
  background-color: var(--gray);
  width: 100%;
  height: 4px;
}
#console-resizer:hover{
  cursor: ns-resize;
}
#console{
  width: 100%;
  overflow: auto;
  background-color: black;
  height: 75px;
}
::-webkit-scrollbar {
  width: 7px;
}
::-webkit-scrollbar-track {
  background: hwb(0 0% 100%);
}
::-webkit-scrollbar-thumb {
  background: var(--white);
  border-radius: 5px;
}



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
