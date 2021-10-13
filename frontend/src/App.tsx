import React from 'react'
import { useSelector } from 'react-redux'
import { RootState } from './redux/store'
import './App.css'

function App() {
  const smiles = useSelector((state: RootState) => state.alkane.smiles)
  const iupac = useSelector((state: RootState) => state.alkane.iupac)

  return (
    <div className="App">
      <h1>Smiles: {smiles}</h1>
      <h1>IUPAC: {iupac}</h1>
    </div>
  )
}

export default App
