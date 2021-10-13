import { createSlice } from '@reduxjs/toolkit'

interface AlkaneState {
  smiles: string
  iupac: string
}

// Define the initial state using that type
const initialState: AlkaneState = {
  smiles: 'CC(C)CC',
  iupac: '2-methylbutane',
}

const alkaneSlice = createSlice({
  name: 'alkane',
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {},
})

export default alkaneSlice.reducer
