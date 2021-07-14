<template>
  <div v-cloak>
    <el-menu
      :default-active="currentRoute"
      mode="horizontal"
      class="el-menu-vertical-demo"
      :background-color="background_color"
      :text-color="text_color"
      :active-text-color="active_text_color"
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
import uiColors from '../../utils/colors'
export default {
  name: 'app-nav-menu',
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
  min-height: 60px;
  position:fixed; /* fixing the position takes it out of html flow - knows
                   nothing about where to locate itself except by browser
                   coordinates */
  left:0;           /* top left corner should start at leftmost spot */
  top:0;            /* top left corner should start at topmost spot */
  width:100vw;      /* take up the full browser width */
  z-index:200;  /* high z index so other content scrolls underneath */
  border-color: var(--background_color);
}
</style>
