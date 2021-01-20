# Todo App Project Frontend

Frontend App has app ui and called backend api

## Requirements
[yarn](https://yarnpkg.com/)

[pact](https://github.com/pact-foundation/jest-pact)


## Install, Build And Tests


Project Setup
```
$ yarn install
```

Compile development
```
$ yarn dev
```

Lints and fixes files
```
$ yarn run lint
```

### Test Process

Run the unit test
```
$ yarn test:unit
```

Run the consumer test
```
$ yarn test:contract
```

Publish the consumer test
```
$ yarn test:pact:publish
```

## Directory Structure

```
.
├── client                 
│   ├── ApiService.js         - Using for backend-api call with axios.     
├── components         
│   ├── AddTodo.vue           - Todo input and button defined here.
│   ├── TodoList.vue          - This file is calling other components 
│   ├── TodoListItem.vue      - This compenent has delete and marked item properties          
├── pact
│   ├── pacts
│     ├── frontend-todo-backend-todo.json  - This file has contract log between frontend and backend project          
├── pages
│   ├── index.vue   - Main vue file
├── test
│   ├── contract
│     ├── ApiService.pact.interactions.js  - File has request interactions   
│     ├── ApiService.spec.js               - File has contract testing
│   ├── unit
│     ├── AddTodo.spec.js                  - This file has button and input unit tests 
│     ├── TodoList.spec.js                 - This file has general unit test for todolist
│     ├── TodoListItem.spec.js             - This file has marked , deleted and others unit tests
├── Dockerfile         
├── jest.config.js     - Jest configuration file
├── nuxt.config.js     - Nuxt configuration file
├── package.json
```

## Deployment Process
On Gitlab Pipeline
  - build
  - lint
  - unit test
  - contract-test
  - pact publish
  - package
  - deploy




