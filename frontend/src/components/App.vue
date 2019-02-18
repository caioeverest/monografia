<template>
  <div id="app">
    <Orquestrador class="wrapper"/>
  </div>
</template>

<script>
import Orquestrador from './layout/Orquestrador.vue';

export default {
  name: 'app',
  components: {
    Orquestrador
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Roboto');
#app {
  font-family: Roboto, Arial, sans-serif;
}

.wrapper {
    display: flex;
    width: 100%;
    align-items: stretch;
    perspective: 1500px;
}

body {
    font-family: Roboto, Arial, sans-serif !important;
    background: #f3f3f3 !important;
}

p {
    font-family: Roboto, Arial, sans-serif;
    font-size: 1.1em;
    font-weight: 300;
    line-height: 1.7em;
    color: #999;
}

.line {
    width: 100%;
    height: 1px;
    border-bottom: 1px dashed #ddd;
    margin: 40px 0;
}

.tipografia-tab {
    font-size: 16px;
}
</style>
