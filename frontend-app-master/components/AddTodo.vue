<template>
  <div id="AddTodoDiv" class="marginTopBottom50x100">
    <div class="row">
      <div class="col-md-12">
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text">Write To-Do</span>
          </div>
          <input
            id="todoText"
            class="form-control"
            aria-label="Write To-Do"
            v-model="title"/>
          <button
            id="addButton"
            type="submit"
            class="btn btn-success marginLeft20px"
            v-on:click="addTodo">
            ADD TO-DO
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { addTodoListItem } from '~/client/ApiService'

export default {
  name: 'AddTodo',
  data: function () {
    return {
      title: ''
    }
  },
  methods: {
    async addTodo (e) {
      e.preventDefault()

      const response = await addTodoListItem(this.title)
      if (response.status !== 200 && !response.data.title) {
        return console.error(response.status)
      }

      this.$emit('add-todo-item', response.data)
      this.title = null
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.marginTopBottom50x100 {
  margin-top: 50px;
  margin-bottom: 100px;
}

.marginLeft20px {
  margin-left: 20px;
}

.borderTodoList {
  border-radius: 10px;
  background-color: #efefef;
  border: 1px solid grey;
}
</style>
