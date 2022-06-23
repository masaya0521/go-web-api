<template>
  <div class="container mx-auto">
    <div class="flex flex-col">
      <span class="text-gray-800 text-4xl my-16">ToDo List</span>
      <TodoAddBox />
      <button class="shadow text-gray-800 h-10 w-48 px-4 py-2 m-2">test</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import TodoAddBox from '~/components/top/TodoAddBox/TodoAddBox.vue';

interface Data {
  data: string;
}

export default Vue.extend({
  name: 'ToDo',
  components: { TodoAddBox },
  data(): Data {
    return {
      data: 'test',
    };
  },
});
</script>
