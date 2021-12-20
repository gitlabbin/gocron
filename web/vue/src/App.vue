<template>
  <el-container>
    <el-header>
      <app-header></app-header>
      <app-nav-menu></app-nav-menu>
    </el-header>
    <el-main >
      <div id="main-container" v-cloak>
        <router-view/>
      </div>
    </el-main>
    <el-footer>
      <app-footer></app-footer>
    </el-footer>
  </el-container>
</template>

<script>
import installService from './api/install'
import appHeader from './components/common/header.vue'
import appNavMenu from './components/common/navMenu.vue'
import appFooter from './components/common/footer.vue'

export default {
  name: 'App',
  data () {
    return {}
  },
  created () {
    installService.status((data) => {
      if (!data) {
        this.$router.push('/install')
      }
    })
  },
  components: {
    appHeader,
    appNavMenu,
    appFooter
  }
}
</script>
<style>
  :root {
    --red: #ff0000;
    --background_color: #324157;
    --text_color: #bfcbd9;
    --active_text_color: #049eff;
  }
  [v-cloak] {
    display: none !important;
  }
  body {
    margin:0;
    background-color: #EFF2F7;
  }
  .el-header {
    padding:0;
    margin:0;
  }
  .el-container {
    padding:0;
    margin:0;
    width: 100%;
  }
  .el-main {
    padding:0;
    margin:0;
  }
  #main-container .el-main {
    height: calc(100vh - 116px);
    margin:20px 20px 0 20px;
  }
  .el-aside .el-menu {
    height: 100%;
    background-color: var(--background_color);
    border-color: var(--background_color);
    text-color: var(--text_color);
    active-text-color: var(--active_text_color);
  }

  .el-table--border:after,
 .el-table--group:after,
 .el-table:before {
   background-color: #C0C4CC;
  }

  .el-table--border,
  .el-table--group {
    border-color: #C0C4CC;
  }

  .el-table td,
  .el-table th.is-leaf {
    border-bottom: 1px solid #C0C4CC;
  }

  .el-table--border th,
  .el-table--border th.gutter:last-of-type {
    border-bottom: 1px solid #C0C4CC;
  }

  /*
  .el-table--border td,
  .el-table--border th {
    border-right: 1px solid #CCCCCC;
  }
  */
</style>
