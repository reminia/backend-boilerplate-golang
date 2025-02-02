# Backend Boilerplate

Use it for setting up a take-home interview assignment or a portfolio project.

## Contents

- [Backend service](https://github.com/DevSkillsHQ/backend-boilerplate-golang/tree/main/app) - a Golang service with a `/ping` endpoint. Extend with your code.
- [API test suite](https://github.com/DevSkillsHQ/backend-boilerplate-golang/blob/main/cypress/integration/backend.spec.js) - a Cypress test suite. Extend with your tests.
- [Pipeline](https://github.com/DevSkillsHQ/backend-boilerplate-golang/blob/main/.github/workflows/tests.yml) - a test Runner that executes the Cypress tests on push to a branch other than `master`/`main`.

## Tech Stack

- Golang
- Cypress
- GitHub Actions

## Cypress test

1. npm install. Install cypress dep. 
2. npm run start. Start the go server.
3. npm test. Test go endpoints with Cypress.

---

Made by [DevSkills](https://devskills.co). 

Did you find this repo useful? **Give us a shout on [Twitter](https://twitter.com/DevSkillsHQ) / [LinkedIn](https://www.linkedin.com/company/devskills)**.
