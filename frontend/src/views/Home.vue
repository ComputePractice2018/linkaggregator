<template>
  <div class="home">
    <LinkAggregator title="Новостной агрегатор"/>
    <Security v-if="index_log == 0" title="Регистрация"/>
    <p v-if="index_log > 0">Рабочее поле</p>
  </div>
</template>

<script>
// @ is an alias to /src
import LinkAggregator from '@/components/LinkAggregator.vue'
import Security from '@/components/Security.vue'

export default {
  name: 'home',
  components: {
    LinkAggregator,
    Security
  },
  data: function () {
    return {
      index_log: 1
    }
  },
  mounted: function () {
    var cookie = document.cookie
    if (cookie.Expires == 0) {
      this.index_log -= 1
    }
  }
}
</script>
