<script lang="ts">
import check from './assets/check.svg'
import terminal from './assets/terminal.svg'
import backup from './assets/backup.svg'
import list from './assets/list.svg'
import stop from './assets/stop.svg'
import {selectedWindow,selectedDistro} from './store'
import {TerminateWsl,LaunchDistro} from '../wailsjs/go/main/App.js'
export var distroRow

function launchDistro(){
    LaunchDistro(distroRow.Name)
}

function createBackup(){
    $selectedDistro=distroRow.Name
    $selectedWindow="BackupCreate"
}

function manageBackups(){
    $selectedDistro=distroRow.Name
    $selectedWindow="BackupList"
}

function terminate(){
  TerminateWsl(distroRow.Name).then()
}
    
</script>



<tr class="row">
    <td class="col-default">
        {#if distroRow.Default_}
            <span class="default-mark" title="Default distribution"><img src="{check}" alt="default"></span>
        {/if}
    </td>
    <td class="col-name">{distroRow.Name}</td>
    <td>
        <span class="status" class:running={distroRow.Status==="Running"}>
            <span class="dot"></span>{distroRow.Status}
        </span>
    </td>
    <td class="col-version">{distroRow.Wsl_version}</td>
    <td>
        <div class="btns">
            <button class="icon-btn" title="Open in terminal" on:click={launchDistro}><img src="{terminal}" alt="terminal"></button>
            <button class="icon-btn" title="Backup" on:click={createBackup}><img src="{backup}" alt="backup"></button>
            <button class="icon-btn" title="Manage backups" on:click={manageBackups}><img src="{list}" alt="list"></button>
            {#if distroRow.Status==="Running"}
                <button class="icon-btn danger" title="Terminate" on:click={terminate}><img src="{stop}" alt="stop"></button>
            {/if}
        </div>
    </td>
</tr>


<style>
.row td {
    background-color: var(--surface-2);
    color: var(--text);
    font-size: 15px;
    height: 52px;
    text-align: left;
    border-top: 1px solid var(--border);
    border-bottom: 1px solid var(--border);
    transition: background-color 0.12s ease;
}
.row:hover td {
    background-color: var(--surface-3);
}
.col-name {
    font-weight: 600;
    padding-left: 18px;
}
.col-default,
.col-version {
    text-align: center !important;
    color: var(--text-dim);
}
td:first-child {
    border-top-left-radius: var(--radius);
    border-bottom-left-radius: var(--radius);
    border-left: 1px solid var(--border);
}
td:last-child {
    border-top-right-radius: var(--radius);
    border-bottom-right-radius: var(--radius);
    border-right: 1px solid var(--border);
}

.default-mark {
    display: inline-flex;
}
.default-mark img {
    height: 16px;
    width: 16px;
    /* tint toward the teal accent */
    filter: invert(72%) sepia(38%) saturate(720%) hue-rotate(116deg) brightness(95%) contrast(88%);
}

/* Status badge */
.status {
    display: inline-flex;
    align-items: center;
    gap: 7px;
    font-size: 13px;
    font-weight: 600;
    color: var(--text-dim);
    padding: 4px 11px 4px 9px;
    border-radius: 999px;
    background: rgba(148, 163, 184, 0.1);
}
.status .dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: var(--text-faint);
}
.status.running {
    color: var(--ok);
    background: rgba(52, 211, 153, 0.12);
}
.status.running .dot {
    background: var(--ok);
    box-shadow: 0 0 0 3px rgba(52, 211, 153, 0.18);
}

.btns {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 4px;
    padding-right: 10px;
    user-select: none;
}
</style>