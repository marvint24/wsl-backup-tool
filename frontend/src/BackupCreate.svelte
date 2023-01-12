<script lang="ts">
  import folder from "./assets/folder.svg"
  import {CreateBackupFile} from '../wailsjs/go/main/App.js'
 
      
  </script>
  
  
  
  <section>

  </section>
  
  
  <style>
  td>div>img{
      cursor: pointer;
  }
  #btns{
    margin-left: 20px;
    user-select: none;
  }
  #btns>img{
    margin-left: 15px;
  }
  img{
      height: 25px;
      filter: invert(88%) sepia(8%) saturate(1480%) hue-rotate(178deg) brightness(106%) contrast(104%);
      padding-top: 5px;
  }
  tr{
      background-color: var(--gray);
      border-radius: 50px;
  }
  td{
      text-align: left;
      color: var(--white);
      font-size: 20px;
      height: 44px;
  }
  tr>td:first-child , tr>td:nth-child(4){
      text-align: center;
  }
  tr>td:last-child{
      border-left-width: 2px;
      border-right-width:0;
      border-top-width:0;
      border-bottom-width:0;
      border-width: 2px,0px,0px,0px; 
      border-color: var(--green-light);
      border-style: solid;
  }
  td:first-child {
      border-top-left-radius: 15px; 
      border-bottom-left-radius: 15px;
  }
  td:last-child {
      border-bottom-right-radius: 15px; 
      border-top-right-radius: 15px; 
  }
  </style>