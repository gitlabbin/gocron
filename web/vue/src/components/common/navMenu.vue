<template>
  <div v-cloak>
    <el-menu
      :default-active="currentRoute"
      mode="horizontal"
      class="el-menu-vertical-demo"
      background-color="#324157"
      text-color="#bfcbd9"
      active-text-color="#049eff"
      router>
      <el-row>
        <el-col :span="2">
          <el-menu-item index="/">
            <template slot="title">
              <i class="el-icon-menu"></i>
              <span slot="title" style="font-size: 200%; background: inherit;">GoCron</span>
            </template>
          </el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item index="/task">{{ $t('menu_jobs') }}</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item index="/host">{{ $t('menu_nodes') }}</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item v-if="this.$store.getters.user.isAdmin" index="/user">{{ $t('menu_users') }}</el-menu-item>
        </el-col>
        <el-col :span="2">
          <el-menu-item v-if="this.$store.getters.user.isAdmin" index="/system">{{ $t('menu_settings') }}</el-menu-item>
        </el-col>
        <el-col :span="14"></el-col>
        <el-col :span="2" style="float:right;">
          <el-submenu v-if="this.$store.getters.user.token" index="userStatus">
            <template slot="title">
              <i class="el-icon-setting"></i>
              <span slot="title">{{this.$store.getters.user.username}}</span>
            </template>
            <el-menu-item index="/user/edit-my-password">{{ $t('menu_password') }}</el-menu-item>
            <el-menu-item @click="logout" index="/user/logout">{{ $t('menu_logout') }}</el-menu-item>
          </el-submenu>
        </el-col>
      </el-row>
    </el-menu>
  </div>
</template>

<script>

export default {
  name: 'app-nav-menu',
  data () {
    return {}
  },
  computed: {
    currentRoute () {
      if (this.$route.path === '/') {
        return '/task'
      }
      const segments = this.$route.path.split('/')
      return `/${segments[1]}`
    }
  },
  methods: {
    logout () {
      this.$store.commit('logout')
      this.$router.push('/')
    }
  }
}
</script>
<style scoped lang='scss'>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 100%;
  min-height: 60px;
}
</style>
