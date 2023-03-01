import React from 'react'
import ScrollingView from './ScrollView'

describe('<ScrollingView />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<ScrollingView />)
  })
})