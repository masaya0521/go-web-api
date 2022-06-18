<template>
  <div class="flex flex-row">
    <input
      v-model="inputText"
      type="text"
      class="shadow h-10 w-48 px-4 py-2 m-2"
      placeholder=" ToDoの入力"
    />
    <button class="shadow text-gray-800 h-10 w-16 px-4 py-2 m-2">Add</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

interface Data {
  inputText: string;
}

export default Vue.extend({
  name: 'InputBox',
  data(): Data {
    return {
      inputText: '',
    };
  },
});
</script>
