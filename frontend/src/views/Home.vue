<template>
  <div class="home">
    <LinkAggregator title="Новостной агрегатор"/>
  </div>
</template>

<script>
// @ is an alias to /src
import LinkAggregator from '@/components/LinkAggregator.vue'

export default {
  name: 'home',
  components: {
    LinkAggregator
  }
}
</script>
