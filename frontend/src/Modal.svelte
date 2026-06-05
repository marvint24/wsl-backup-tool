<script lang="ts">
  export let heading: string = ""
</script>

<section>
  <div>
    <div class="box">
      <div class="heading">{heading}</div>
      <hr/>
      <slot name="content" />
    </div>
    <div class="actions">
      <slot name="actions" />
    </div>
  </div>
</section>

<style>
  section {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: var(--dark-op);
    z-index: 1;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .box {
    background-color: var(--white);
    border-radius: 15px;
  }
  .heading {
    color: var(--dark);
    font-size: 20px;
    font-weight: 500;
    padding: 5px 0 0 25px;
    text-align: left;
    white-space: nowrap;
  }
  hr {
    border: none;
    border-top: 2px solid var(--dark);
    margin-bottom: 15px;
  }
  .actions {
    position: relative;
    top: 10px;
    margin-left: auto;
    margin-right: 0;
    display: flex;
    width: fit-content;
    flex-direction: row;
  }
</style>
