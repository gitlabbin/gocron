<template>
  <el-aside width="150px">
    <el-menu
      :default-active="currentRoute"
      mode="vertical"
      background-color="#324157"
      text-color="#bfcbd9"
      active-text-color="#049eff"
      router>
      <el-menu-item index="/system">Notification</el-menu-item>
      <el-menu-item index="/system/login-log">LoginLog</el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'system-sidebar',
  data () {
    return {}
  },
  computed: {
    currentRoute () {
      if (this.$route.path === '/system/login-log') {
        return '/system/login-log'
      }
      return '/system'
    }
  }
}
</script>
