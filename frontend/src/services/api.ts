import { toast } from 'react-toastify'

import axios, { AxiosInstance } from 'axios'
import { IUPACResponse } from '../types'

const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_SERVER_API_URL,
})

axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    const msg = error.response?.data?.msg || 'Failed to connect to server'
    toast.error(msg, {
      position: toast.POSITION.TOP_RIGHT,
    })
  }
)

const api = {
  getIUPACFromSMILES: (smiles: string) =>
    axiosInstance.get<IUPACResponse>('/iupac', { params: { smiles } }),
}

export default api
