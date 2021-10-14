import { Formik, Form, Field, FormikErrors, ErrorMessage } from 'formik'
import { Label, Button } from 'reactstrap'
import clsx from 'clsx'

import { InputFormValues, InputFormProps } from '../types'

function InputForm({ smiles, onSubmit }: InputFormProps) {
  const initialValues: InputFormValues = { smiles }

  function handleValidate(values: InputFormValues) {
    const errors: FormikErrors<InputFormValues> = {}

    if (!values.smiles) errors.smiles = 'SMILES is required'
    else if (!/^[C()]+$/.test(values.smiles))
      errors.smiles = "SMILES must consists of only 'C' and round brackets"
    else if (values.smiles[0] !== 'C')
      errors.smiles = "SMILES must start with 'C'"

    return errors
  }

  function handleSubmit(values: InputFormValues) {
    onSubmit(values.smiles)
  }

  return (
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
    />
  )
}

export default InputForm
