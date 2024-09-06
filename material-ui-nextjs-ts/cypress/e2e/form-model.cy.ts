describe('Form Model Creation and List View Test', () => {
  it('should create a new form model and check if it appears in the list view', () => {
    // Step 1: Visit the page
    cy.visit('http://localhost:3000')
    cy.get(`a[href="/form"]`).click()

    // Step 2: Click the "New Form Model" button to open the dialog
    cy.contains('button', 'Add Form Model').click();

    // Step 3: Fill in the fields on the "Personal Info" tab
    cy.get('input[name="name"]').type('Test Name');         // Assuming the field name is 'name'
    cy.get('input[name="age"]').type('30');                 // Assuming the field name is 'age'

    cy.get('input[name="dateTime"]').type('09/09/20241200AM'); // Assuming field name is 'dateTime'

    // Select Gender (radio button)
    cy.get('input[name="gender"]').check('male');           // Assuming field name 'gender'

    // Step 4: Switch to the "Preferences" tab
    cy.contains('button', 'Preferences').click();  // Change the tab to 'Preferences'

    // Step 5: Fill in the form fields in the "Preferences" tab
    cy.get('input[name="skills"]').type('JavaScript{downArrow}{enter}'); // Assuming multi-select for skills
    // Close the multi select overlay
    cy.get('form').click({force: true})

    // Fill in the city field after switching tabs
    cy.get('input[name="city"]').type('New York{downArrow}{enter}');          // Assuming the field name is 'city'

    // Interact with the satisfaction slider (adjust if necessary)
    // Hacky solution to move the slider with cypress code
    cy.get('[data-testid="satisfactionLevel"] input[type=range]', { timeout: 60000 }).invoke('val', '90');
    cy.get('[data-testid="satisfactionLevel"] .MuiSlider-thumbColorPrimary').invoke('attr', 'style', 'left: 90%');
    cy.get('[data-testid="satisfactionLevel"] .MuiSlider-track').invoke('attr', 'style', 'width: 90%');


    // Check the "termsAccepted" checkbox
    cy.get('[data-testid="termsAccepted"] input[type="checkbox"]').check();           // Assuming field name 'termsAccepted'

    // Step 6: Switch to the "Tasks" tab
    cy.contains('button', 'Tasks').click();  // Change the tab to 'Tasks'

    // Step 7: Add a new task in the "Todo List"
    cy.contains('button', 'Add Task').click(); // Click the "Add Task" button
    cy.get('input[name="todos.0.task"]').type('Complete Cypress Test'); // Fill in the task

    cy.get(`[data-testid="task 2"]`).find(`[data-testid="DeleteIcon"]`).click();

    // Step 8: Submit the form
    cy.contains('button', 'Submit').click();

    // Step 9: Wait for the form to be saved and the list to refresh
    // Ensure the form is closed
    cy.get('dialog').should('not.exist');

    // Step 11: Verify that the task also appears in the table or somewhere else relevant
    cy.get(`[data-testid="example-list"] tr[data-index="0"] td[data-index="1"]`).should('have.text', "Test Name")
    cy.get(`[data-testid="example-list"] tr[data-index="0"] td[data-index="2"]`).should('have.text', "30")
    cy.get(`[data-testid="example-list"] tr[data-index="0"] td[data-index="3"]`).should('have.text', "New York")
    cy.get(`[data-testid="example-list"] tr[data-index="0"] td[data-index="4"]`).should('have.text', "male")

    // Delete all the rows
    cy.get('[data-testid="example-list"] tbody tr').each(($row) => {
      // Right-click on the first cell of the current row
      const r = cy.wrap($row)
      r.find('td').first().rightclick();
      cy.get(`[role="menu"] [data-testid="DeleteIcon"]`).click();
    });
  });
});
