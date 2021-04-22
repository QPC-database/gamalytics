import axios from 'axios'
import { IUser } from '@/types'

export function getUser (userid: string) {
  return axios.get(`user/${encodeURI(userid)}`).catch((err) => {
    throw (err)
  })
}

export function newUser (user: IUser) {
  return axios.post('user/new', user).catch((err) => {
    throw (err)
  })
}
