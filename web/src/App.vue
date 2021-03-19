<template>
  <Header title="gamalytics" />
  <router-view></router-view>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'
import Header from './components/Header.vue'

@Options({
  components: {
    Header
  }
})
export default class App extends Vue {}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.padding {
  padding: 0 20px;
}

nav {
  text-align: center;
  margin: 0 auto;
}
</style>
