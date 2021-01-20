import { mount } from '@vue/test-utils'
import TodoList from '../../components/TodoList.vue'
import { getTodoListItems } from '~/client/ApiService'
import flushPromises from 'flush-promises'

jest.mock('~/client/ApiService')

describe('Todo List', () => {
  it('GivenEmptyTodoListWhenNothingIsDoneThenICanSeeEmptyTodoLlist', () => {
    const todoWrapper = mount(TodoList)
    expect(todoWrapper.find('.todoList').exists()).toBeFalsy()
  })

  it('GivenTodoListItemsWhenCalledFillTodoListThenICanSeeItemOnTodoList', () => {
    const title = 'Buy some milk'
    const todoWrapper = mount(TodoList, {
      data () {
        return {
          todoList: [{ title: title, completed: false }]
        }
      }
    })

    expect(todoWrapper.find('.todoList').exists()).toBeTruthy()
  })

  it('GivenOpenPageWhenIRequestToTodoListThenICanSeeTodoListItemsOnTodoList', async () => {
    const response = {
      status: 200,
      data:
        [
          { id: 1, title: 'First Todo Item', completed: false },
          { id: 2, title: 'Second Todo Item', completed: false }
        ]
    }

    await getTodoListItems.mockResolvedValue(response)
    const wrapper = mount(TodoList, {
      mocks: {
        $t: () => {
        }
      }
    })

    await flushPromises()

    expect(wrapper.find('button[type="button"]').classes('todoitem-1')).toBeTruthy()
  })
})
