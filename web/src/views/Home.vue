<template>
    <div v-if="loggedIn">
      Welcome {{ user.firstname }}
    </div>
    <div v-else>
      Welcome to Gamalytics, please: <Link link="login" name="log in"/> or <Link link="register" name="register a new account"/>
    </div>
</template>

<script lang="ts">
import { IUser } from '@/types'
import { defineComponent } from 'vue'
import Link from '@/components/Link.vue'

export default defineComponent({
  data () {
    return {
      user: {} as IUser,
      loggedIn: false
    }
  },
  async created () {
    // this.user = (await getUser('antonio.jelic@vediagames.com')).data
  },
  components: {
    Link
  }
})
</script>

<style scoped>

</style>
