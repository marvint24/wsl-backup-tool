<script lang="ts">
  import {EventsOn} from '../wailsjs/runtime/runtime.js'
  import WslList from './WslList.svelte'
  import {distros,selectedWindow,settings,shutdownWindow} from './store'
  import sliders from './assets/sliders.svg'
  import Console from './Console.svelte'
  import {GetSettings} from '../wailsjs/go/main/App.js'

  import BackupCreate from './BackupCreate.svelte'
  import BackupList from './BackupList.svelte'
  import Settings from './Settings.svelte'

  import WslShutdown from './WslShutdown.svelte'

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
  EventsOn("SettingsChanged",()=>{
    if($settings.ShowConsole){
      initResizeVars()
    }
  })

  let resizer
  let element
  function initResizeVars(){
    resizer=document.getElementById("console-resizer")
    element=document.getElementById("console")
    resizer.addEventListener('mousedown', initResize, false)
  }

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
    if($settings.ShowConsole){
      initResizeVars()
    }
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

  EventsOn("openWslShutdown",()=>{
    $shutdownWindow=true
    console.log("openWslShutdownasdfsd")
  })
</script>
  
<main>
  {#if $shutdownWindow}
    <WslShutdown/>
  {/if}
  {#if selectedWindow}
    <svelte:component this={imported[$selectedWindow]} />
  {/if}
  <header class="topbar">
    <h1 class="topbar-title">Distributions</h1>
    <button class="settings-btn" title="Settings" on:click="{openSettings}">
      <img src="{sliders}" alt="">
      <span>Settings</span>
    </button>
  </header>
  <WslList/>
  {#if $settings.ShowConsole}
    <section class="console-section">
      <div id="console-resizer"></div>
      <div id="console"><Console/></div>
    </section>
  {/if}
</main>

<style>

main {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.topbar {
  flex: 0 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid var(--border);
}
.topbar-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: -0.01em;
  color: var(--text);
}

.settings-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: var(--surface-2);
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
  padding: 8px 14px 8px 11px;
  color: var(--text-dim);
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.12s ease, color 0.12s ease, border-color 0.12s ease;
}
.settings-btn:hover {
  background: var(--surface-3);
  color: var(--text);
  border-color: var(--border-strong);
}
.settings-btn img {
  height: 17px;
  width: 17px;
  filter: invert(88%) sepia(8%) saturate(1480%) hue-rotate(178deg) brightness(106%) contrast(104%);
}

/* Console */
.console-section {
  flex: 0 0 auto;
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
}
#console-resizer {
  background-color: var(--border);
  width: 100%;
  height: 4px;
  transition: background-color 0.12s ease;
}
#console-resizer:hover {
  cursor: ns-resize;
  background-color: var(--accent);
}
#console {
  width: 100%;
  overflow: auto;
  background-color: #0a0d12;
  border-top: 1px solid var(--border);
  height: 75px;
}
</style>
