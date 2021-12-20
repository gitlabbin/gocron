// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App'
import i18n from './i18n'
import router from './router'
import store from './store/index'

Vue.config.productionTip = false
Vue.use(ElementUI, {
  i18n: (key, value) => i18n.t(key, value)
})

Vue.directive('focus', {
  inserted: function (el) {
    // 聚焦元素
    el.focus()
  }
})

Vue.prototype.$appConfirm = function (callback) {
  this.$confirm(this.$t('confirm_action_info'), this.$t('action_warn'), {
    confirmButtonText: this.$t('buttons_ok'),
    cancelButtonText: this.$t('action_cancel'),
    type: 'warning'
  }).then(() => {
    callback()
  }).catch(() => {})
}

Vue.filter('formatTime', function (time) {
  const fillZero = function (num) {
    return num >= 10 ? num : '0' + num
  }
  const date = new Date(time)

  const result = date.getFullYear() + '-' +
  (fillZero(date.getMonth() + 1)) + '-' +
  fillZero(date.getDate()) + ' ' +
  fillZero(date.getHours()) + ':' +
  fillZero(date.getMinutes()) + ':' +
  fillZero(date.getSeconds())

  if (result.indexOf('20') !== 0) {
    return ''
  }

  return result
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  store,
  components: { App },
  template: '<App/>'
})
