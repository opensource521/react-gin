import axios, { AxiosInstance } from 'axios'
import { IUPACResponse } from '../types'

const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_SERVER_API_URL,
})

const api = {
  getIUPACFromSMILES: (smiles: string) =>
    axiosInstance.get<IUPACResponse>('/iupac', { params: { smiles } }),
}

export default api
