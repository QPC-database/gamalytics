import axios from 'axios'
import { IUser } from '@/types'

export function logIn (user: IUser) {
  return axios.post('auth/login', user).catch((err) => {
    throw (err)
  })
}
