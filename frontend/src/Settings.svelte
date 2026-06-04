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

<section class="modal-overlay">
  <div class="modal">
    <div class="modal-header">
      <h2 class="modal-title">Settings</h2>
    </div>
    <div class="modal-body">
      <form id="SettingsForm">
        <div class="field">
          <label class="field-label" for="BackupPath-Setting">Backup folder</label>
          <div class="path-row">
            <input id="BackupPath-Setting" class="input" type="text" bind:value={tempSettings.BackupPath} on:change={checkPath}>
            <button type="button" class="icon-btn path-btn" title="Browse for folder" on:click="{openFolder}"><img src="{folder}" alt="folder"></button>
          </div>
        </div>

        <div class="field">
          <label class="field-label" for="RefreshInterval-Setting">Refresh interval (seconds)</label>
          <input id="RefreshInterval-Setting" class="input input-narrow" title="in seconds" min="1" type="number" bind:value={tempSettings.RefreshInterval}>
        </div>

        <div class="toggle-row">
          <div>
            <div class="toggle-label">Show console</div>
            <div class="toggle-hint">Display the live log panel at the bottom of the window</div>
          </div>
          <label class="switch">
            <input id="ShowConsole-Setting" type="checkbox" bind:checked={tempSettings.ShowConsole}>
            <span class="slider"></span>
          </label>
        </div>
      </form>
    </div>
    <div class="modal-footer">
      <button class="btn btn-ghost" on:click="{close}">Cancel</button>
      <button class="btn btn-primary" on:click="{okay}">Save</button>
    </div>
  </div>
</section>

<style>
form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}
.field {
  display: flex;
  flex-direction: column;
}
.path-row {
  display: flex;
  gap: 8px;
  align-items: stretch;
}
.path-btn {
  flex: 0 0 auto;
  width: 40px;
  height: auto;
  align-self: stretch;
  border: 1px solid var(--border-strong);
  background: var(--surface-2);
}
.path-btn:hover {
  background: var(--surface-3);
}
.input-narrow {
  max-width: 140px;
}

/* Toggle row */
.toggle-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 14px 16px;
  background: var(--surface-2);
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
}
.toggle-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--text);
  text-align: left;
}
.toggle-hint {
  font-size: 12.5px;
  color: var(--text-dim);
  margin-top: 2px;
  text-align: left;
}

/* Switch */
.switch {
  position: relative;
  display: inline-block;
  width: 42px;
  height: 24px;
  flex: 0 0 auto;
}
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}
.slider {
  position: absolute;
  cursor: pointer;
  inset: 0;
  background: var(--surface-3);
  border: 1px solid var(--border-strong);
  border-radius: 999px;
  transition: background-color 0.15s ease, border-color 0.15s ease;
}
.slider::before {
  content: "";
  position: absolute;
  height: 16px;
  width: 16px;
  left: 3px;
  bottom: 3px;
  background: var(--text-dim);
  border-radius: 50%;
  transition: transform 0.15s ease, background-color 0.15s ease;
}
.switch input:checked + .slider {
  background: var(--accent-soft);
  border-color: var(--accent-border);
}
.switch input:checked + .slider::before {
  transform: translateX(18px);
  background: var(--accent);
}
</style>