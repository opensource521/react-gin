describe('Test App', () => {
  it('App should work', () => {
    cy.visit('http://localhost:3000')

    cy.contains('Welcome to World Of Alkane!')
  })

  it('App should convert', () => {
    cy.get('input[name="smiles"]').type('CC(C)CC')

    cy.contains('Convert').click()

    cy.contains('2-methylbutane')
  })

  it('App should validate', () => {
    cy.get('input[name="smiles"]').clear()

    cy.contains('Convert').click()

    cy.contains('SMILES is required')
  })
})
