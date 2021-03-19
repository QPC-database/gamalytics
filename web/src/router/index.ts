import { createRouter, createWebHistory } from 'vue-router'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import Home from '@/views/Home'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import LogIn from '@/views/LogIn'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import Register from '@/views/Register'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Log In',
    component: LogIn
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }
]

const Router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default Router
