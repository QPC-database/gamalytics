<template>
  <router-link class="padding" v-bind:to=link>{{ name }}</router-link>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    link: String,
    name: String
  }
})
</script>

<style scoped>

</style>
