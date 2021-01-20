<template>
  <div class="container" id="app">
    <h3>WELCOME</h3>
    <TodoList/>
  </div>
</template>

<script>

import TodoList from '../components/TodoList.vue'

export default {
  name: 'index',
  components: {
    TodoList
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
