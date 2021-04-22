<template>
  <p>Please fill out the form below to register</p>
  <form>
    <InputBox id="firstname" type="text" message="First name" />

    <InputBox id="lastname" type="text" message="Last name" />

    <InputBox id="email" type="email" message="Email" />

    <InputBox id="password" type="password" message="Password" />

    <InputBox id="confirmPassword" type="password" message="Confirm password" />

    <input type="submit" value="Submit">
  </form>
</template>

<script>
import { Options, Vue } from 'vue-class-component'
import InputBox from '@/components/InputBox'
import { IUser } from '@/types'

@Options({
  name: 'Register',
  components: {
    InputBox
  }
})

export default class Register extends Vue {}
</script>

<style scoped>

</style>
