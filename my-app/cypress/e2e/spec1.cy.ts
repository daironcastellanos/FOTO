describe('ScrollingView', ()=>{
  it('displays images', ()=>{
      cy.visit('/scrolling-view');

      cy.get('.grid').should('exist');
      cy.get('.img').should('exist');
      cy.get('.img').should('have.length', 50);

  });
});

export {};