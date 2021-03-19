<template>
  <p>Please log in</p>
  <form>
    <InputBox id="email" type="email" message="Email" />

    <InputBox id="password" type="password" message="Password" />

    <input type="submit" value="Submit">
  </form>
</template>

<script>
import { Options, Vue } from 'vue-class-component'
import InputBox from '@/components/InputBox'

@Options({
  name: 'LogIn',
  components: {
    InputBox
  }
})

export default class LogIn extends Vue {}

</script>

<style scoped>

</style>
