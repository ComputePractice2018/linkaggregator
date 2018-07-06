<template>
  <div class="linkaggregator">
    <h1>{{ title }}</h1>

    <h1>Тематики:</h1>
    <table style="width:100%">
      <th><input type="button" value="Theme1" v-on:click="choose_theme('Theme1')"/></th>
      <th><input type="button" value="Theme2" v-on:click="choose_theme('Theme2')"/></th>
      <th><input type="button" value="Theme3" v-on:click="choose_theme('Theme3')"/></th>
      <th><input type="button" value="Theme4" v-on:click="choose_theme('Theme4')"/></th>
      <th><input type="button" value="All news" v-on:click="show_all_news"/></th>
    </table>

    <h1>Новости:</h1>
    <table style="width:50%">
      <tr>
        <th>Title</th>
        <th>Date</th>
        <th>URL</th>
        <th>Theme</th>
      </tr>
      <tr v-for="post in news_theme_list" v-bind:key="post.url">
        <td>{{ post.title }}</td>
        <td>{{ post.date }}</td>
        <td>{{ post.url }}</td>
        <td>{{ post.theme }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'LinkAggregator',
  props: {
    title: String
  },
  data: function () {
    return {
      news_list: [
        {
          'title': 'Title1',
          'date': 'Date1',
          'url': 'Url1',
          'theme': 'Theme1'
        },
        {
          'title': 'Title2',
          'date': 'Date2',
          'url': 'Url2',
          'theme': 'Theme2'
        },
        {
          'title': 'Title3',
          'date': 'Date3',
          'url': 'Url3',
          'theme': 'Theme3'
        },
        {
          'title': 'Title4',
          'date': 'Date4',
          'url': 'Url4',
          'theme': 'Theme4'
        },
        {
          'title': 'Title5',
          'date': 'Date5',
          'url': 'Url5',
          'theme': 'Theme1'
        },
        {
          'title': 'Title6',
          'date': 'Date6',
          'url': 'Url6',
          'theme': 'Theme2'
        },
        {
          'title': 'Title7',
          'date': 'Date7',
          'url': 'Url7',
          'theme': 'Theme1'
        }
      ],
      news_theme_list: []
    }
  },
  methods: {
    choose_theme: function (theme) {
      this.news_theme_list.length = 0
      var length = this.news_list.length
      for (var i = 0; i < length; i++) {
        if (theme == this.news_list[i].theme) {
          this.news_theme_list.push(this.news_list[i])
        }
      }
    },
    show_all_news: function () {
      this.news_theme_list.length = 0
      var length = this.news_list.length
      for (var i = 0; i < length; i++) {
        this.news_theme_list.push(this.news_list[i])
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to tdis component only -->
<style scoped>
  th, td {
    text-align: center
  }
</style>
