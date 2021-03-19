<template>
  <div>
    <label :for=id >{{message}}:</label> <br>
    <input :type=type :id=id :name=id>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'

@Options({
  props: {
    id: String,
    type: String,
    message: String
  }
})

export default class InputBox extends Vue {
  id!: string
  type!: string
  message!: string
}
</script>

<style scoped>
  div {
    margin: 20px auto;
  }
</style>
