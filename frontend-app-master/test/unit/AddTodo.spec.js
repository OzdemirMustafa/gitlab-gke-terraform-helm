import { mount } from '@vue/test-utils'
import AddTodo from '../../components/AddTodo.vue'
import { addTodoListItem } from '~/client/ApiService'

jest.mock('~/client/ApiService')

describe('Todo Add', () => {
  const wrapper = mount(AddTodo)

  it('GivenOpenPageWhenPageLoadThenICanSeeInputOnPage', () => {
    expect(wrapper.find('input').exists()).toBeTruthy()
  })

  it('GivenOpenPageWhenPageLoadThenICanSeeButtonOnPage', () => {
    expect(wrapper.find('button').exists()).toBeTruthy()
  })

  it('GivenOpenPageWhenNothingIsDoneThenICanSeeEmptyValue', () => {
    expect(wrapper.find('input').element.value).toEqual('')
  })

  test('GivenOpenPageWhenWritedValueThenIWillSeeChangedInputValue', () => {
    const inputValue = 'Buy some milks'
    wrapper.find('input').setValue(inputValue)

    expect(wrapper.find('input').element.value).toEqual(inputValue)
  })

  it('GivenWrittenTodoNameWhenClickToAddTodoButtonThenIWillSeeThatAddTodoMethodWorked', () => {
    const title = 'Buy some milk'
    const wrapper = mount(AddTodo, {
      data: function () {
        return {
          title: title
        }
      }
    })

    expect(wrapper.find('input').element.value).toEqual(title)
  })

  it('GivenEmptyTodoListWhenIClickToAddButtonThenMockedFunctionWillBringResponseAndItShouldBeEmit', async () => {
    const response = { status: 200, data: { id: 1, title: 'First Todo Item', completed: false } }
    await addTodoListItem.mockResolvedValue(response)

    const wrapper = mount(AddTodo, {
      data: function () {
        return {
          title: response.data.title
        }
      },
      mocks: {
        $t: () => {
        }
      }
    })

    await wrapper.find('button').trigger('click')
    await wrapper.vm.$nextTick()

    expect(wrapper.emitted('add-todo-item')[0][0].title).toEqual(response.data.title)
  })
})
