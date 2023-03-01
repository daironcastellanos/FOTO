describe('GET /api/users/get', () => {
    it('returns a list of users', () => {
      cy.request('/api/users/get')
        .its('status')
        .should('equal', 200);
  
      cy.request('/api/users/get')
        .its('body')
        .should('have.length', 10);
    });
  });

  