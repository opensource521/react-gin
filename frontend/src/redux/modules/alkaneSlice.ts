import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import api from '../../services/api'
import { IUPAC, SMILES } from '../../types'

interface AlkaneState {
  smiles: SMILES
  iupac: IUPAC
}

const initialState: AlkaneState = {
  smiles: '',
  iupac: '',
}

export const getIUPACFromSMILES = createAsyncThunk(
  'alkanes/getIUPACFromSMILES',
  async (smiles: SMILES) => {
    const response = await api.getIUPACFromSMILES(smiles)
    return {
      smiles,
      iupac: response.data.result,
    }
  }
)

const alkaneSlice = createSlice({
  name: 'alkane',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(getIUPACFromSMILES.fulfilled, (state, action) => {
      state.smiles = action.payload.smiles
      state.iupac = action.payload.iupac
    })
  },
})

export default alkaneSlice.reducer
