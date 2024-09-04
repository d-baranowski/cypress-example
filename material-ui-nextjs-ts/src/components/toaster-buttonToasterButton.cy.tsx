import React from 'react'
import ToasterButton from './toaster-button'

describe('<ToasterButton />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<ToasterButton />)
  })
})