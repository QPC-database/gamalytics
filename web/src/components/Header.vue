<template>
  <header>
    <h4>{{ title }}</h4>
    <nav>
      <Link :link="link.path" :name="link.name" :key="link.path" v-for="link in links" />
    </nav>
  </header>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'
import Link from '@/components/Link.vue'
import Router from '@/router'

@Options({
  props: {
    title: String
  },
  components: {
    Link
  }
})

export default class Header extends Vue {
  title!: string
  links = Router.getRoutes()
}
</script>

<style scoped>
header {
  padding-bottom: 50px;
}
</style>
