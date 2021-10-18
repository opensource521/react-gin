import React from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { Row, Col } from 'reactstrap'
import { ToastContainer } from 'react-toastify'

import InputForm from './components/InputForm'
import { RootState } from './redux/store'
import { getIUPACFromSMILES } from './redux/modules/alkaneSlice'
import { SMILES } from './types'

function App() {
  const smiles = useSelector((state: RootState) => state.alkane.smiles)
  const iupac = useSelector((state: RootState) => state.alkane.iupac)
  const dispatch = useDispatch()

  function handleSubmit(smiles: SMILES) {
    dispatch(getIUPACFromSMILES(smiles))
  }

  return (
    <div className="container-fluid">
      <div className="p-5 mb-4 bg-light">
        <h1 className="display-3">Welcome to World Of Alkane!</h1>
        <p className="lead">
          This is a simple app that converts SMILES (Simplified Molecular-Input
          Line-Entry System) format into IUPAC (International Union of Pure and
          Applied Chemistry) format
        </p>
      </div>
      <main className="p-5">
        <Row className="justify-content-center">
          <Col lg={6}>
            <div className="w-50">
              <InputForm smiles={smiles} onSubmit={handleSubmit} />
            </div>
            <h3 className="mt-3">IUPAC: {iupac}</h3>
          </Col>
        </Row>
      </main>
      <ToastContainer />
    </div>
  )
}

export default App
