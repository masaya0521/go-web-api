<template>
  <div>
    <Tutorial/>
    <h2> {{ api }}</h2>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  data() {
    return {
      message: 'This is a message',
      api : []
    }
  },

  mounted() {
    this.$axios.get('api/json')
     // eslint-disable-next-line no-return-assign
     .then(response => this.api = response.data);
  },
  })
</script>
