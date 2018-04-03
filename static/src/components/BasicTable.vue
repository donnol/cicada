<template>
  <div id='btable'>
      {{ header }}
  </div>
</template>

<script>
export default {
  name: 'BasicTable',
  data() {
    return {
      header: 'i am a basic table'
    }
  }
}
</script>

<style scoped>
#btable {
  background-color: #090;
}
</style>
