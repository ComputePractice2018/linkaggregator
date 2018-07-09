<template>
  <div class="linkaggregator">
    <h1>{{ title }}</h1>

    <input type="button" value="Подписки" v-on:click="show_subscriptions"/>
    <div v-if="index_sub_show > -1">
      <div v-if="index_sub == -1">
        <p>Подписок нет</p>
      </div>

      <div v-if="index_sub > -1">
        <p>Вы подписаны на темы:</p>
        <input v-for="sub in subscriptions" v-bind:key="sub" type="button" :value="sub" v-on:click="choose_theme(sub)"/>
      </div>

      <input type="button" v-on:click="edit_subscriptions" value="Редактировать подписки"/>

      <div v-if="index_edit_sub > -1" class="flex-sub-edit">
          <div v-for="theme in themes" v-bind:key="theme">
            {{ theme }}<input v-if="index_sub_unsub[themes.indexOf(theme)] < 0" type="button" value="Подписаться" v-on:click="add_subscription(theme, themes.indexOf(theme))"/>
            <input v-if="index_sub_unsub[themes.indexOf(theme)] == 0" type="button" value="Отписаться" v-on:click="delete_subscription(theme, themes.indexOf(theme))"/>
          </div>
          <input type="button" value="Готово" v-on:click="index_edit_sub = -1"/>
      </div>
      <p></p>
      <input type="button" value="Скрыть подписки" v-on:click="index_sub_show = -1"/>
    </div>
    <p></p>

    <div id="themes" class="flex-theme">
      <input type="button" value="Theme1" v-on:click="choose_theme('Theme1')"/>
      <input type="button" value="Theme2" v-on:click="choose_theme('Theme2')"/>
      <input type="button" value="Theme3" v-on:click="choose_theme('Theme3')"/>
      <input type="button" value="Theme4" v-on:click="choose_theme('Theme4')"/>
      <input type="button" value="All news" v-on:click="show_all_news"/>
    </div>

    <div id="news" class="flex-container">
      <div v-for="post in news_theme_list" v-bind:key="post.url">
        <p>{{ post.title }}</p>
      </div>
    </div>
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
      index_edit_sub: -1, // индекс, который отвечает за блок редактирования подписок: -1 - блок не отображается
      index_sub_unsub: [], // массив, в котором хранится информация он наличии темы в подписках по порядку тем: -1 - темы нет
      news_list: [],
      news_theme_list: [],
      subscriptions: [],
      themes: ['Theme1', 'Theme2', 'Theme3', 'Theme4']
    }
  },
  mounted: function () {
    this.get_news()
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
    },
    show_subscriptions: function () {
      this.error = ''
      const url = '/api/linkaggregator/subscriptions'
      axios.get(url).then(response => {
        this.subscriptions = response.data
        this.index_sub_show++
        var length = this.subscriptions.length
        if (length == 0) {
          this.index_sub = -1
        }
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_subscriptions: function () {
      this.index_edit_sub++
      var length = this.themes.length
      for (var i = 0; i < length; i++) {
        if (this.subscriptions.indexOf(this.themes[i]) == -1) {
          this.index_sub_unsub[i] = -1 // -1 означает, что темы нет в подписках
        } else {
          this.index_sub_unsub[i] = 0 // 0 означает, что тема в подписках
        }
      }
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
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to tdis component only -->
<style scoped>
  th, td {
    text-align: center
  }
  .flex-container {
    display: inline-flex;
    width: 100%;
    height: 100%;
    align-content: flex-start;
    flex-direction: column;
  }
  .flex-container > div {
    background: #b7c9c6;
    margin: 10px;
    padding: 10px;
    border-radius: 4px;
  }
  .flex-theme {
    display: inline-flex;
    width: 100%;
  }
  .flex-theme > input {
    flex-grow: 1;
    padding: 10px;
    background: #b7c9c6;
    outline: none;
    border-color: black;
    border-width: 1px;
  }
  .flex-theme > input:hover {
    background: #9fadab;
  }
  .flex-theme > input:active {
    background: #a1a1a1;
  }
  .flex-sub-edit {
    display: inline-flex;
    width: 10%;
    height: 10%;
    align-content: flex-start;
    flex-direction: column;
  }
</style>
