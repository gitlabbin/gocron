<template>
  <div><h5>version: {{ version }}</h5></div>
</template>

<script>
export default {
  name: 'app-footer',
  data () {
    return {
      version: process.env.VUE_APP_VERSION
    }
  }
}
</script>
