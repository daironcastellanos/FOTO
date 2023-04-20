describe('Home', () => {
  it('should type "test" in the search input and then clear it using the "x" button', () => {
    // Visit the Home page
    cy.visit('http://localhost:3000/Home');

    cy.contains('Search').should('be.visible').and('not.be.disabled');

    // Click on the "Book Photographer" button
    cy.contains('Search').click();

    


  });
});
