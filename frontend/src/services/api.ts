import axios, { AxiosInstance } from 'axios'
import { IUPACResponse } from '../types'

const Axios: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_SERVER_API_URL,
})

const api = {
  getIUPACFromSMILES: (smiles: string) =>
    Axios.get<IUPACResponse>('/iupac', { params: { smiles } }),
}

export default api
