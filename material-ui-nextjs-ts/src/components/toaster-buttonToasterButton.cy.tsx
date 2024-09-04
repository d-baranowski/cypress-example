import React from 'react'
import ToasterButton from './toaster-button'

describe('<ToasterButton />', () => {
  it(`when I press the button a network request is made and the response get's toasted`, () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<ToasterButton />)
    cy.intercept(
      {
        method: 'POST', // Route all GET requests
        url: 'http://mock-cypress-test.com/greet.v1.GreetService/Greet', // that have a URL that matches '/users/*'
      },
      { fixture: 'helloWorldApi-example-response.json' }
    ).as('helloWorldApi')

    cy.get(`[data-testid="toaster-btn"]`).click();
    cy.get(`[role="status"]`).should('have.text', "Hello, World!");

    cy.wait(`@helloWorldApi`).then((interception) => {
      assert.isNotEmpty(interception.response?.body)
      assert.equal(JSON.stringify({ name: "World" }), JSON.stringify(interception.request?.body))
    })
  })
})
