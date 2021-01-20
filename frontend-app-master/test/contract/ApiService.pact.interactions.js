const { Matchers } = require('@pact-foundation/pact')

const addTodoListItemRequest = {
  uponReceiving: 'a post todo list item',
  method: 'POST',
  path: '/addTodoListItem/FirstTodo'
}

const addTodoListItemResponse = {
  status: 200,
  body: {
    id: Matchers.like(1),
    title: Matchers.like('FirstTodo'),
    completed: Matchers.like(false)
  }
}

const getTodoListItemsRequest = {
  uponReceiving: 'a get request to get todo list items',
  method: 'GET',
  path: '/getTodoList'
}

const getTodoListItemsResponse = {
  status: 200,
  body: [{
    id: Matchers.like(1),
    title: Matchers.like('First Todo Item'),
    completed: Matchers.like(false)
  }, {
    id: Matchers.like(2),
    title: Matchers.like('Second Todo Item'),
    completed: Matchers.like(false)
  }]
}

const updateTodoItemCompletedRequest = {
  uponReceiving: 'a post comleted item',
  method: 'POST',
  path: '/updateItemCompleted/1/true'
}

const updateTodoItemCompletedResponse = {
  status: 200,
  body: {
    id: Matchers.like(1),
    title: Matchers.like('Todo Item'),
    completed: Matchers.like(true)
  }
}

const deleteTodoItemRequest = {
  uponReceiving: 'a post delete item',
  method: 'POST',
  path: '/deleteTodoItem/1'
}

const deleteTodoItemResponse = {
  status: 200,
  body: {
    id: Matchers.like(1),
    title: Matchers.like(''),
    completed: Matchers.like(true)
  }
}

module.exports = {
  getTodoListItems: {
    state: 'i have a todo list',
    uponReceiving: 'a request for getting todo list items',
    withRequest: getTodoListItemsRequest,
    willRespondWith: getTodoListItemsResponse
  },
  addTodoListItemByTitle: {
    state: 'i have not empty product list which matches given barcode',
    uponReceiving: 'a request for getting to added todo item',
    withRequest: addTodoListItemRequest,
    willRespondWith: addTodoListItemResponse
  },
  updateTodoItemCompleted: {
    state: 'i have will update todo item',
    uponReceiving: 'a request for update to item completed value',
    withRequest: updateTodoItemCompletedRequest,
    willRespondWith: updateTodoItemCompletedResponse
  },
  deleteTodoItem: {
    state: 'i have will delete todo item',
    uponReceiving: 'a request for delete to item ',
    withRequest: deleteTodoItemRequest,
    willRespondWith: deleteTodoItemResponse
  }
}
