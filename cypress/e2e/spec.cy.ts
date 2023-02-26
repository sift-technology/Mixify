describe('Open Website', () => {
  it('website opens', () => {
    cy.visit('http://localhost:4200')
  })
})

describe('Title', () => {
  it('title in tab is "Mixify"', () => {
    cy.visit('http://localhost:4200')

    cy.title().should('equal', "Mixify")
    
  })
})

describe('Answer Buttons', () => {
  it('clicking answer buttons displays response', () => {
    cy.visit('http://localhost:4200')

    cy.get('mat-button-toggle[type=button]').eq(0).click()
    cy.get('#response1').should('have.text', 'Chosen value is 1')

    cy.get('mat-button-toggle[type=button]').eq(6).click()
    cy.get('#response2').should('have.text', 'Chosen value is 3')

    cy.get('input[type=range]').as('range').invoke('val', 25).trigger('change')
    cy.get('#response3').should('have.text', 'Chosen value is 25')

    cy.get('mat-button-toggle[type=button]').eq(11).click()
    cy.get('#response4').should('have.text', 'Chosen value is 4')
  })
})

describe('Submit Button', () => {
  it('clicking "submit" navigates to a new url', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Submit').click()

    cy.url().should('include', '/results')
  })
})