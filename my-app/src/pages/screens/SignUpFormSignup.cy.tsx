import React from 'react'
import Signup from './SignUpForm'

describe('<Signup />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Signup />)
  })
})