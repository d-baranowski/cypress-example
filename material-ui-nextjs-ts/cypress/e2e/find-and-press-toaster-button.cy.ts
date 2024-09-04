describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000')
    cy.get(`a[href="/about"]`).click()
    cy.get(`[data-testid="toaster-btn"]`).click();
    cy.get(`[role="status"]`).should('have.text', "Hello, World!");
  })
})
