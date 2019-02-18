<template>
  <b-table hover :items="obj"></b-table>
</template>

<script>
export default {
  name: 'Table',
  props: {
    obj: Object
  }
}
</script>

<style>
</style>