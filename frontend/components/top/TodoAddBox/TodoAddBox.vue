<template>
  <div class="flex flex-row">
    <input
      v-model="post.content"
      type="text"
      class="shadow h-10 w-48 px-4 py-2 m-2"
      placeholder=" ToDoの入力"
    />
    <button
      class="shadow text-gray-800 h-10 w-16 px-4 py-2 m-2"
      @click="postTodo()"
    >
      Add
    </button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

interface Data {
  post: {
    content: string;
    author: string;
  };
}

export default Vue.extend({
  name: 'InputBox',
  data(): Data {
    return {
      post: {
        content: '',
        author: 'koshio',
      },
    };
  },
  methods: {
    async postTodo(): Promise<void> {
      await this.$axios.post('api/post', this.post);
    },
  },
});
</script>
