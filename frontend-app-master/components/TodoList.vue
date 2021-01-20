<template>
  <div class="">
    <AddTodo @add-todo-item="fillTodoList"/>

    <div class="row">
      <div class="col-md-12">
        <div class="p-5 borderTodoList">
          <h3>TODO BOARD</h3>
          <ul class="marginTop10">
            <li class="''+this.todoListItem.complete" v-bind:key="todoListItem.id" v-for="(todoListItem, index) in todoList">
              <TodoListItem @delete-todo-item="deleteTodo(index)" v-bind:todoListItem="todoListItem"/>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TodoListItem from './TodoListItem'
import AddTodo from './AddTodo'
import { getTodoListItems } from '~/client/ApiService'

export default {
  name: 'TodoList',
  components: {
    TodoListItem,
    AddTodo
  },
  data () {
    return {
      todoList: []
    }
  },
  methods: {
    fillTodoList (newTodo) {
      this.todoList = [...this.todoList, newTodo]
    },
    async getTodoList () {
      const response = await getTodoListItems()
      if (response.status === 200 && response.data !== null) {
        Object.keys(response.data).forEach(key => {
          this.todoList.push(response.data[key])
        })
      }
    },
    deleteTodo (index) {
      this.$delete(this.todoList, index)
    }
  },
  async mounted () {
    await this.getTodoList()
  }
}
</script>

<style>
.borderTodoList {
  border-radius: 10px;
  background-color: #efefef;
  border: 1px solid grey;
}

.marginTop10 {
  margin-top: 30px;
}
</style>
