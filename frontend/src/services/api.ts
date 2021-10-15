import { toast } from 'react-toastify'

import axios, { AxiosInstance } from 'axios'
import { IUPACResponse } from '../types'

const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_SERVER_API_URL,
})

axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error!.response!.data!.msg) {
      toast.error(error.response.data.msg, {
        position: toast.POSITION.TOP_RIGHT,
      })
    } else {
      toast.error('Network error', {
        position: toast.POSITION.TOP_RIGHT,
      })
    }
  }
)

const api = {
  getIUPACFromSMILES: (smiles: string) =>
    axiosInstance.get<IUPACResponse>('/iupac', { params: { smiles } }),
}

export default api
