<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>

        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="websiteUrl" type="text" id="website-input" placeholder="Enter a website" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
        </form>

        <img :src="thumbnailUrl">
      </div>
    </div>
  </div>
</template>
<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

<script>
import axios from 'axios';
export default {
  name: 'App',
  data() {
    return { websiteUrl: "", thumbnailUrl : ""}
  },

  methods: {
    makeWebsiteThumbnail() {
      // axios.get("https://jsonplaceholder.typicode.com/photos/1", {})
      axios.post("http://localhost:8000/api/thumbnail", {
        url: this.websiteUrl,
      })
      .then((response) => {
        this.thumbnailUrl = response.data.thumbnailUrl
      })
      .catch((error) => {
        window.alert(`The Api returned an error: ${error}`)
      })
    }
  }

}
</script>
