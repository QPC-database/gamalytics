<template>
  <p>Please log in</p>
  <form @submit.prevent="submitLogIn">
     <InputBox v-model="email" id="email" type="email" message="Email" />
     <InputBox v-model="password" id="password" type="password" message="Password" />
    <input type="submit" value="Submit">
  </form>
</template>

<script>
import { defineComponent } from 'vue'
import InputBox from '@/components/InputBox'

export default defineComponent({
  name: 'LogIn',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  components: {
    InputBox
  },
  methods: {
    submitLogIn: function (e) {
      console.log(this.email)
    }
  }
})

</script>

<style scoped>

</style>
