import axios from 'axios'

process.env.VUE_APP_API_URL = 'http://35.224.98.66'

/**
 *
 * @returns {Promise<*>}
 */
export function getTodoListItems () {
  return axios.get(`${process.env.VUE_APP_API_URL}/getTodoList`)
    .then((response) => response)
}

/**
 *
 * @param todoItemTitle
 * @returns {Promise<*>}
 */
export function addTodoListItem (todoItemTitle) {
  return axios
    .post(`${process.env.VUE_APP_API_URL}/addTodoListItem/${todoItemTitle}`)
    .then((response) => response)
}

/**
 *
 * @param itemId
 * @param completed
 * @returns {Promise<*>}
 */
export function markedTodoItem (itemId, completed) {
  return axios
    .post(`${process.env.VUE_APP_API_URL}/updateItemCompleted/${itemId}/${completed}`)
    .then((response) => response)
}

/**
 *
 * @param itemId
 * @returns {Promise<*>}
 */
export function deleteTodoItem (itemId) {
  return axios
    .post(`${process.env.VUE_APP_API_URL}/deleteTodoItem/${itemId}`)
    .then((response) => response)
}
