<script lang="ts">
  import {GetWslList} from '../wailsjs/go/main/App.js'
  import WslList from './WslList.svelte'
  import {distros} from './store'
  import refresh from './assets/refresh.svg'
  import sliders from './assets/sliders.svg'

  function getWslDistros(): void {
    GetWslList().then((result)=> {
      let jsonObj=JSON.parse(result)
      $distros=jsonObj.map((obj)=>{
        obj.uuid=crypto.randomUUID()
        return obj
      })      
    })
  }

  function openSettings(): void {

  }

  var spinClass=""
  function refreshWslList(): void {
    spinClass="spin"
    getWslDistros()
    setTimeout(()=>{spinClass=""},1000)
  }
 
  getWslDistros()

</script>
  
<main>
  <section id="header">
    <div id="Settings" title="Settings"><img src="{sliders}" alt="sliders"><p>Settings</p>  </div>
    <div id="Refresh" title="Refresh" on:click="{refreshWslList}" on:keydown><img class="{spinClass}" src="{refresh}" alt="refresh"></div>
  </section>
  <WslList/>
</main>

<style>
#header{
  margin-bottom: 50px;
  display: flex;
  justify-content: space-between;
}
img{
    height: 25px;
    filter: invert(88%) sepia(8%) saturate(1480%) hue-rotate(178deg) brightness(106%) contrast(104%);
}
#Settings{
  border-radius: 0 0 15px 0;
}
#Settings>p{
  color: var(--white);
  font-size: 20px;
  margin: 10px 0 10px 5px;
  font-weight: 500;
}
#Refresh{
  border-radius: 0 0 0 15px;
}
#header>div{
  display: flex;
  flex-direction: row;
  align-items: center;;
  padding-left: 15px;
  padding-right: 15px;
  background-color: var(--green-light);
}
#header>div:hover{
  cursor: pointer;
}

@keyframes spin {
  from {transform: rotate(0);}
  to   {transform: rotate(180deg);}
}
.spin{
  animation-name: spin;
  animation-duration: 1s;
  animation-iteration-count:1;
}


</style>
