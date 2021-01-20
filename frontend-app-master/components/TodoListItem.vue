<template>
  <div>
    <button v-on:click="markedTodoItem"
            v-bind:class="'todoitem-'+todoListItem.id"
            v-bind:style="this.marked ? 'text-decoration:line-through' : 'none'" type="button"
            class="btn btn-light col-lg-12 marginTopBottom todoList">{{ todoListItem.title }}
    </button>
    <a v-on:click="deleteTodoItem" class="close closeButtonPosition delete ">x</a>
  </div>
</template>

<script>
import { deleteTodoItem, markedTodoItem } from '~/client/ApiService'

export default {
  name: 'TodoListItem',
  data () {
    return {
      marked: this.todoListItem.completed,
      itemId: this.todoListItem.id
    }
  },
  props: [
    'todoListItem'
  ],
  methods: {
    async markedTodoItem () {
      const response = await markedTodoItem(this.itemId, !this.marked)
      if (response.status === 200) {
        this.marked = !this.marked
      }
    },
    async deleteTodoItem (e) {
      e.preventDefault()
      const response = await deleteTodoItem(this.itemId)
      if (response.status === 200) {
        this.$emit('delete-todo-item', this.itemId)
      }
    }
  }

}
</script>

<style scoped>
.buttonBorder {
  border: 1px solid black !important;
}

.marginTopBottom {
  margin-top: 2px;
  margin-bottom: 2px;
}

.closeButtonPosition {
  position: relative;
  bottom: 35px;
  cursor: pointer;
}
</style>
