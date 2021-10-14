import React from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { Label, Button, Row, Col } from 'reactstrap'
import { Formik, Form, Field, FormikErrors, ErrorMessage } from 'formik'
import clsx from 'clsx'

import { RootState } from './redux/store'
import { getIUPACFromSMILES } from './redux/modules/alkaneSlice'
import { SMILEFormValues } from './types'

function App() {
  const smiles = useSelector((state: RootState) => state.alkane.smiles)
  const iupac = useSelector((state: RootState) => state.alkane.iupac)
  const dispatch = useDispatch()

  const initialValues: SMILEFormValues = { smiles }

  function handleValidate(values: SMILEFormValues) {
    const errors: FormikErrors<SMILEFormValues> = {}

    if (!values.smiles) errors.smiles = 'SMILES is required'
    else if (!/^[C()]+$/.test(values.smiles))
      errors.smiles = "SMILES must consists of only 'C' and round brackets"
    else if (values.smiles[0] !== 'C')
      errors.smiles = "SMILES must start with 'C'"

    return errors
  }

  function handleSubmit(values: SMILEFormValues) {
    dispatch(getIUPACFromSMILES(values.smiles))
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
      <main>
        <Row className="justify-content-center">
          <Col lg={3} md={6}>
            <Formik
              initialValues={initialValues}
              validate={handleValidate}
              onSubmit={handleSubmit}
              render={({ errors, touched }) => (
                <Form>
                  <Label>SMILES *</Label>
                  <Field
                    name="smiles"
                    placeholder="Input SMILES"
                    className={clsx('form-control', {
                      'is-invalid': errors.smiles && touched.smiles,
                    })}
                  />
                  <ErrorMessage
                    name="smiles"
                    component="div"
                    className="invalid-feedback"
                  />
                  <Button type="submit" color="primary" className="mt-2">
                    Convert
                  </Button>
                </Form>
              )}
            ></Formik>
            <h3 className="mt-3">IUPAC: {iupac}</h3>
          </Col>
        </Row>
      </main>
    </div>
  )
}

export default App
