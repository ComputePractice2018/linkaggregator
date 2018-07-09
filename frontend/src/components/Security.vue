<template>
  <div class="security">
      <h1>{{ title }}</h1>
      <div id="auth" class="flex-container">
        <div id="c1">
          <form name="signup">
            <p>Login</p>
            <input type="text" v-model="new_user.login">
            <p>Пароль</p>
            <input type="text" v-model="new_user.password">
            <p>E-mail</p>
            <input type="text" v-model="new_user.email">
            <p/>
            <button v-on:click="sign_up">Зарегистрироваться</button>
          </form>
        </div>
        <div id="c2">
          <form name="signin">
            <p>Login</p>
            <input type="text" v-model="exist_user.login">
            <p>Пароль</p>
            <input type="text" v-model="exist_user.password">
            <p/>
            <button v-on:click="sign_in">Войти</button>
          </form>
        </div>
      </div>
  </div>
</template>

<script>
const axios = require('axios')
export default {
  name: 'Security',
  props: {
    title: String
  },
  data: function () {
    return {
      new_user: {
        'login': '',
        'password': '',
        'email': ''
      },
      exist_user: {
        'login': '',
        'password': ''
      },
      error: ''
    }
  },
  methods: {
    sign_up: function () {
      this.error = ''
      const url = '/security/register'
      axios
        .post(url, this.new_user)
        .then(response => {
          if (response.data.success) {
            console.log('signed up')
          } else {
            console.log(response.data.message)
          }
        })
        .catch(response => {
          this.error = response.response.data
        })
    },
    sign_in: function () {
      this.error = ''
      const url = '/security/login'
      axios
        .post(url, this.exist_user)
        .then(response => {
          if (response.data.success) {
            console.log('signed in')
          } else {
            console.log(response.data.message)
          }
        })
        .catch(response => {
          this.error = response.response.data
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex-container input {
    box-sizing: border-box;
    width:100%;
    height:3%;
    background:#e7e7e7;
    border:1px solid #dedede;
    padding:10px;
    margin:auto;
    font-size:0.9em;
    color:#313131;
    -moz-border-radius:5px;
    -webkit-border-radius:5px;
    border-radius:5px;
    outline: none;
}
.flex-container {
  display: flex;
  width: 100%;
  height: 100%;
  margin: auto;
  align-items: flex-start;
  justify-content: center;
}

.flex-container > div {
  margin: 5%;
  padding: 1%;
  background: #a5b9a9;
  border-radius:5px;
}
</style>
