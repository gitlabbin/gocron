<template>
  <el-aside width="150px">
    <el-menu
      :default-active="currentRoute"
      mode="vertical"
      :background-color="background_color"
      :text-color="text_color"
      :active-text-color="active_text_color"
      router>
      <el-menu-item index="/system">Notification</el-menu-item>
      <el-menu-item index="/system/login-log">LoginLog</el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
import uiColors from '../../utils/colors'
export default {
  name: 'system-sidebar',
  data () {
    return {
      red: '#ff0000',
      background_color: uiColors.menu.background_color,
      text_color: uiColors.menu.text_color,
      active_text_color: uiColors.menu.active_text_color
    }
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
