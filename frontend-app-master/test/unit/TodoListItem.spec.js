import { mount } from '@vue/test-utils'
import TodoListItem from '../../components/TodoListItem.vue'
import { deleteTodoItem, markedTodoItem } from '@/client/ApiService'

jest.mock('~/client/ApiService')

describe('Todo List Item', () => {
  it('GivenTodoListItemWithBuySomeMilkWhenIClickToItemICanSeeItem', () => {
    const wrapper = mount(TodoListItem, {
      propsData: {
        todoListItem: {
          id: 1,
          title: 'Buy Some Milk',
          completed: false
        }
      }
    })

    expect(wrapper.props('todoListItem').title).toEqual('Buy Some Milk')
  })

  it('GivenTodoListItemWhenIClickToItemThenMarkedMethodWillBeCalled', async () => {
    const markedTodoItem = jest.fn()
    const wrapper = mount(TodoListItem, {
      propsData: {
        todoListItem: {
          id: 1,
          title: 'Buy Some Milk',
          completed: false
        }
      },
      methods: {
        markedTodoItem
      }
    })

    await wrapper.find('button').trigger('click')
    await wrapper.vm.$nextTick()

    expect(markedTodoItem).toHaveBeenCalled()
  })

  it('GivenTodoListItemWhenIClickToItemThenICanSeeElementMarked', async () => {
    const response = { status: 200, data: { id: 1, title: 'Buy Some Milk', completed: true } }
    await markedTodoItem.mockResolvedValue(response)

    const wrapper = mount(TodoListItem, {
      propsData: {
        todoListItem: {
          id: 1,
          title: 'Buy Some Milk',
          completed: false
        }
      },
      mocks: {
        $t: () => {
        }
      }
    })

    const buttonElement = wrapper.find('button')

    await buttonElement.trigger('click')
    await wrapper.vm.$nextTick()

    expect(buttonElement.element.style.textDecoration).toBe('line-through')
  })
})

it('GivenTodoListItemWhenIClickToDeleteItemThenDeleteMethodWillBeCalled', async () => {
  const deleteTodoItem = jest.fn()
  const wrapper = mount(TodoListItem, {
    propsData: {
      todoListItem: {
        id: 1,
        title: 'Buy Some Milk',
        completed: false
      }
    },
    methods: {
      deleteTodoItem
    }
  })

  await wrapper.find('.delete').trigger('click')
  await wrapper.vm.$nextTick()

  expect(deleteTodoItem).toHaveBeenCalled()
})

it('GivenTodoListItemWhenIClickToDeleteThenICanSeeElementDelete', async () => {
  const response = { status: 200, data: { id: 1, title: 'Buy Some Milk', completed: true } }
  await deleteTodoItem.mockResolvedValue(response)

  const wrapper = mount(TodoListItem, {
    propsData: {
      todoListItem: {
        id: 1,
        title: 'Buy Some Milk',
        completed: false
      }
    },
    mocks: {
      $t: () => {
      }
    }
  })

  const deleteButton = wrapper.find('.delete')

  await deleteButton.trigger('click')
  await wrapper.vm.$nextTick()

  expect(wrapper.emitted('delete-todo-item')).toBeTruthy()
})
