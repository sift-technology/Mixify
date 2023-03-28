describe('Open Website', () => {
  it('website opens', () => {
    cy.visit('http://localhost:63021')
  })
})

describe('Title', () => {
  it('title in tab is "Mixify"', () => {
    cy.visit('http://localhost:63021')

    cy.title().should('equal', "Mixify")
    
  })
})

// this following test no longer works because we removed the chosen value info:

// describe('Answer Buttons', () => {
//   it('clicking answer buttons displays response', () => {
//     cy.visit('http://localhost:63021')

//     cy.get('mat-button-toggle[type=button]').eq(0).click()
//     cy.get('#response1').should('have.text', 'Chosen value is 1')

//     cy.get('mat-button-toggle[type=button]').eq(6).click()
//     cy.get('#response2').should('have.text', 'Chosen value is 3')

//     cy.get('input[type=range]').as('range').invoke('val', 25).trigger('change')
//     cy.get('#response3').should('have.text', 'Chosen value is 25')

//     cy.get('mat-button-toggle[type=button]').eq(11).click()
//     cy.get('#response4').should('have.text', 'Chosen value is 4')
//   })
// })

describe('Submit Button', () => {
  it('clicking "submit" navigates to a new url', () => {
    cy.visit('http://localhost:63021')

    cy.contains('Submit').click()

    cy.url().should('include', '/results')
  })
})

describe('Survey Functionality', () => {
  it('user is able to click through survey', () => {
    cy.visit('http://localhost:63021')

    cy.get('mat-button-toggle[type=button]').eq(0).click()
    cy.get('mat-button-toggle[type=button]').eq(4).click()
    cy.get('input[type=range]').as('range').invoke('val', 5).trigger('change')
    cy.get('mat-button-toggle[type=button]').eq(8).click()
    cy.get('mat-button-toggle[type=button]').eq(12).click()
    cy.get('mat-button-toggle[type=button]').eq(16).click()

    cy.contains('Submit').click()

    cy.url().should('include', '/results')
  })
})

describe('Open Spotify From Results', () => {
  it('clicking song name takes user to Spotify website', () => {
    cy.visit('http://localhost:63021/results')

    cy.contains('Song Title').invoke('attr', 'target', "_self").click()

    cy.url()
  })
})