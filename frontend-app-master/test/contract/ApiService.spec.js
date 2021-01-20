import { addTodoListItem, deleteTodoItem, getTodoListItems, markedTodoItem } from '@/client/ApiService'

/**
 * @jest-environment node
 */
const { pactWith } = require('jest-pact')
const interactions = require('./ApiService.pact.interactions')

pactWith(
  {
    cors: true,
    consumer: 'frontend-todo',
    provider: 'backend-todo',
    port: 4000,
    logLevel: 'INFO'
  },
  (provider) => {
    describe('addToListItem to backend-todo', () => {
      beforeEach(() => {
        return provider.addInteraction(interactions.addTodoListItemByTitle)
      })

      it('returns a successful body', async () => {
        process.env.VUE_APP_API_URL = provider.mockService.baseUrl
        const response = await addTodoListItem('FirstTodo')

        expect(response.data).toEqual(
          expect.objectContaining({ id: expect.any(Number), title: expect.any(String), completed: expect.any(Boolean) })
        )
      })
    })

    describe('getTodoListItems from backend-todo', () => {
      beforeEach(() => {
        return provider.addInteraction(interactions.getTodoListItems)
      })

      it('returns a successful body', async () => {
        process.env.VUE_APP_API_URL = provider.mockService.baseUrl
        const response = await getTodoListItems()

        expect(response.data).toEqual(
          expect.arrayContaining([
            expect.objectContaining({ id: expect.any(Number), title: expect.any(String), completed: expect.any(Boolean) }),
            expect.objectContaining({ id: expect.any(Number), title: expect.any(String), completed: expect.any(Boolean) })])
        )
      })
    })

    describe('updateItemCompleted from backend-todo', () => {
      beforeEach(() => {
        return provider.addInteraction(interactions.updateTodoItemCompleted)
      })

      it('returns a successful body', async () => {
        process.env.VUE_APP_API_URL = provider.mockService.baseUrl
        const response = await markedTodoItem(1, true)

        expect(response.data).toEqual(
          expect.objectContaining({ id: expect.any(Number), title: expect.any(String), completed: expect.any(Boolean) })
        )
      })
    })

    describe('deleteTodoItem from backend-todo', () => {
      beforeEach(() => {
        return provider.addInteraction(interactions.deleteTodoItem)
      })

      it('returns a successful body', async () => {
        process.env.VUE_APP_API_URL = provider.mockService.baseUrl
        const response = await deleteTodoItem(1)

        expect(response.data).toEqual(
          expect.objectContaining({ id: expect.any(Number), title: expect.any(String), completed: expect.any(Boolean) })
        )
      })
    })
  }
)
