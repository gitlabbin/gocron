<template>
  <el-aside width="150px">
    <el-menu
      :default-active="currentRoute"
      mode="vertical"
      background-color="#324157"
      text-color="#bfcbd9"
      active-text-color="#049eff"
      router>
      <el-menu-item index="/task">{{ $t('sidebar_jobs') }}</el-menu-item>
      <el-menu-item index="/task/log">{{ $t('sidebar_logs') }}</el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'task-sidebar',
  data () {
    return {}
  },
  computed: {
    currentRoute () {
      if (this.$route.path === '/task/log') {
        return '/task/log'
      }
      return '/task'
    }
  }
}
</script>
