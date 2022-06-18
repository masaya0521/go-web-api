<template>
  <div class="justify-center">
    <button class="shadow h-10 w-48">test</button>
    <input type="text" class="shadow h-10 w-48" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'ToDo',
  data() {
    return {
      message: 'This is a message',
      api: [],
      post: {
        content: 'hello world',
        author: 'masaya',
      },
    };
  },
});
</script>
