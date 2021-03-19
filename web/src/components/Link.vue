<template>
  <router-link class="padding" v-bind:to=link>{{ name }}</router-link>
</template>

<script lang="ts">
import { Vue, Options } from 'vue-class-component'

@Options({
  props: {
    link: String,
    name: String
  }
})

export default class Link extends Vue {
  link!: string
  name!: string
}
</script>

<style scoped>

</style>
