describe('HomePage', () => {
    it('should navigate to the Setting page when the Settings button is clicked', () => {
      cy.visit('http://localhost:3000/screens/Settings') // Replace '/' with the actual URL of your homepage
  
      // Find the Settings button and click it
      const button = cy.contains('Log Out')
      button.click()
  
      // Wait for the new page to load and assert that it's the correct URL
      cy.url().should('include', 'http://localhost:3000') // Replace '/Setting' with the actual URL of your Setting page

   
    })
  })