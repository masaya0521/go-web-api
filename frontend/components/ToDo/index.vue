<template>
  <div class="justify-center">
    <button class="test h-10 w-48">test</button>
    <input type="text" class="test h-10 w-48" />
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

<style scoped>
.test {
  border-radius: 9px;
  background: #ffffff;
  box-shadow: 3px 3px 6px #bebebe, -3px -3px 6px #ffffff;
}
</style>
