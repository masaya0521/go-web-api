<template>
  <div>test22</div>
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
