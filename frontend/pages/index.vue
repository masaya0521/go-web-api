<template>
  <div>
    <ToDo />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import ToDo from '~/components/top/index.vue';

export default Vue.extend({
  name: 'IndexPage',
  components: { ToDo },
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
  /*
  mounted() {
    this.$axios
      .get('api/json')
      // eslint-disable-next-line no-return-assign
      .then((response) => (this.api = response.data));
    this.$axios.post('api/post', this.post);
  },
  */
});
</script>
