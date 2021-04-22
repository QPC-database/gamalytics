<template>
  <div>
    <label :for=id >{{message}}:</label> <br>
    <input :type=type v-bind:value="value" v-on:input="$emit('input', $event.target.value)">
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    id: String,
    type: String,
    message: String,
    value: String
  }
})

</script>

<style scoped>
  div {
    margin: 20px auto;
  }
</style>
