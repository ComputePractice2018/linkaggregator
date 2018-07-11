<template>
  <div class="linkaggregator">
    <b-jumbotron bg-variant="info" text-variant="white" border-variant="dark">
      <template slot="header">
        {{ title }}
      </template>
    </b-jumbotron>
    <b-navbar toggleable="md" type="light" variant="secondary">
      <b-navbar-nav class="container-fluid">
        <b-nav-item class="mr-0"><b-button v-b-toggle.collapse1 variant="info" v-on:click="show_subscriptions">Подписки</b-button></b-nav-item>
        <b-nav-item><b-button variant="light" size="lg" v-on:click="show_sub_news">Рекомендации</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Спорт')">Спорт</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Политика')">Политика</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Общество')">Общество</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Экономика')">Экономика</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Культура')">Культура</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Технологии')">Технологии</b-button></b-nav-item>
        <b-nav-item><b-button variant="secondary" size="md" v-on:click="choose_theme('Наука')">Наука</b-button></b-nav-item>
        <b-nav-item class="ml-0"><b-button size="md" variant="dark" v-on:click="show_all_news">Все новости</b-button></b-nav-item>
      </b-navbar-nav>
    </b-navbar>

    <b-collapse id="collapse1" class="mt-2">
      <b-card style="width: 30%"
        bg-variant="light"
        border-variant="danger">
        <b-button v-b-toggle.collapse1 variant="secondary" v-on:click="show_subscriptions">Скрыть</b-button>
        <div v-if="index_sub == -1">
          <p>Подписок нет</p>
        </div>

        <div v-if="index_sub > -1">
          <p>Вы подписаны на темы:</p>
          <b-button variant="info" v-for="sub in subscriptions" v-bind:key="sub" style="margin: 1px" v-on:click="choose_theme(sub)">{{ sub }}</b-button>
        </div>
        <br>

        <b-button variant="primary" v-b-toggle="'collapse2'" class="m-1">Редактировать подписки</b-button>

        <b-collapse id="collapse2">
          <table class="table" style="width: 25%">
            <tbody>
              <tr v-for="theme in themes" v-bind:key="theme" >
                  <td>{{ theme }}</td>
                  <td>
                    <b-button variant="success" v-if="index_sub_unsub[themes.indexOf(theme)] < 0" v-on:click="add_subscription(theme, themes.indexOf(theme))">Подписаться</b-button>
                    <b-button variant="danger" v-if="index_sub_unsub[themes.indexOf(theme)] == 0" v-on:click="delete_subscription(theme, themes.indexOf(theme))">Отписаться</b-button>
                  </td>
              </tr>
            </tbody>
          </table>
        </b-collapse>
      </b-card>
    </b-collapse>

    <div v-if="index_sub_show > -1"></div>
    <br><br>

    <b-card style="width: 100%" v-for="post in news_theme_list" v-bind:key="post.id"
      :title="post.title"
      :sub-title="post.date"
      bg-variant="light"
      border-variant="info">
      <p class="card-text">{{ post.content }}</p>
        <b-link :href="post.url" class="card-link">{{ post.url }}</b-link>
        <br><br>
        <b-button v-on:click="show_content(post.id)">Читать</b-button>
    </b-card>
  </div>
</template>

<script>
const axios = require('axios')
export default {
  name: 'LinkAggregator',
  props: {
    title: String
  },
  data: function () {
    return {
      index_sub_show: -1, // индекс для блока подписок: -1 - подписки не показываются
      index_sub: 0, // индекс, который показывает, есть ли подписки: -1 - подписок нет
      index_sub_unsub: [], // массив, в котором хранится информация он наличии темы в подписках по порядку тем: -1 - темы нет
      news_list: [],
      news_theme_list: [],
      subscriptions: [],
      themes: ['Спорт', 'Политика', 'Общество', 'Экономика', 'Культура', 'Технологии', 'Наука']
    }
  },
  mounted: function () {
    this.get_news()
    this.show_all_news()
  },
  methods: {
    get_news: function () {
      this.error = ''
      const url = '/api/linkaggregator/news'
      axios.get(url).then(response => {
        this.news_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    choose_theme: function (theme) {
      this.news_theme_list.length = 0
      var length = this.news_list.length
      for (var i = 0; i < length; i++) {
        for (var j = 0; j < this.news_list[i].themes.length; j++) {
          if (theme == this.news_list[i].themes[j]) {
            this.news_theme_list.push(this.news_list[i])
          }
        }
      }
    },
    show_all_news: function () {
      this.news_theme_list.length = 0
      var length = this.news_list.length
      for (var i = 0; i < length; i++) {
        this.news_theme_list.push(this.news_list[i])
      }
    },
    show_sub_news: function () {
      this.news_theme_list.length = 0
      var themeLength = this.themes.length
      var newsLength = this.news_list.length
      for (var i = 0; i < themeLength; i++) {
        if (this.subscriptions.indexOf(this.themes[i]) > -1) {
          for (var j = 0; j < newsLength; j++) {
            for (var k = 0; k < this.news_list[j].themes.length; k++) {
              if (this.themes[i] == this.news_list[j].themes[k]) {
                if (this.news_theme_list.indexOf(this.news_list[j]) == -1) {
                  this.news_theme_list.push(this.news_list[j])
                }
              }
            }
          }
        }
      }
    },
    show_subscriptions: function () {
      this.error = ''
      const url = '/api/linkaggregator/subscriptions'
      axios.get(url).then(response => {
        this.subscriptions = response.data
        if (this.subscriptions.length == 0) {
          this.index_sub = -1
        }
        var length = this.themes.length
        for (var i = 0; i < length; i++) {
          if (this.subscriptions.indexOf(this.themes[i]) == -1) {
            this.index_sub_unsub[i] = -1 // -1 означает, что темы нет в подписках
          } else {
            this.index_sub_unsub[i] = 0 // 0 означает, что тема в подписках
          }
        }
        this.index_sub_show++
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_subscription: function (theme, i) {
      this.error = ''
      var number = this.subscriptions.length + 1
      const url = '/api/linkaggregator/subscriptions/' + number
      axios.put(url, theme).then(response => {
        this.index_sub_unsub[i] = 0 // 0 означает, что тема в подписках
        this.subscriptions.push(theme)
        this.index_sub++
      }).catch(response => {
        this.error = response.response.data
      })
    },
    delete_subscription: function (theme, i) {
      this.error = ''
      var number = this.subscriptions.indexOf(theme)
      const url = '/api/linkaggregator/subscriptions/' + number
      axios.delete(url).then(response => {
        this.index_sub_unsub[i] = -1 // -1 означает, что темы нет в подписках
        this.subscriptions.splice(this.subscriptions.indexOf(theme), 1)
        var exist = 0
        if (this.subscriptions.length > 0) {
          exist++
        }
        if (exist == 0) {
          this.index_sub = -1
        }
      }).catch(response => {
        this.error = response.response.data
      })
    },
    show_content: function (id) {
      this.error = ''
      const url = '/api/linkaggregator/news/content/1'
      axios.get(url).then(response => {
        this.news_theme_list[id].content = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to tdis component only -->
<style scoped>
</style>
