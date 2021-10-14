export type IUPAC = string
export type SMILES = string
export interface InputFormValues {
  smiles: string
}
export interface InputFormProps {
  smiles: SMILES
  onSubmit: (smiles: SMILES) => void
}
export interface IUPACResponse {
  result: IUPAC
}
