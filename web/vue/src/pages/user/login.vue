<template>
    <div>
      <el-dialog
        :title="$t('user_login')"
        :visible.sync="dialogVisible"
        :close-on-click-modal="false"
        :show-close="false"
        :close-on-press-escape="false"
        width="40%">
        <el-form ref="form" @submit.prevent.native="submit" :model="form" label-width="80px"
        :rules="formRules">
          <el-form-item :label="$t('username')" prop="username" >
            <el-col :span="16">
              <el-input v-model.trim="form.username"
                        :placeholder="$t('msg_username_or_email')" autofocus>
              </el-input>
            </el-col>
          </el-form-item>
          <el-form-item :label="$t('password')" prop="password">
            <el-col :span="16">
              <el-input v-model.trim="form.password" type="password" :placeholder="$t('msg_input_password')"></el-input>
            </el-col>
          </el-form-item>
          <el-form-item>
            <el-button native-type="submit" type="primary">{{ $t('login') }}</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
</template>

<script>
import userServcie from '../../api/user'

export default {
  name: 'login',
  data () {
    return {
      form: {
        username: '',
        password: ''
      },
      formRules: {
        username: [
          {required: true, message: this.$t('msg_username_or_email'), trigger: 'blur'}
        ],
        password: [
          {required: true, message: this.$t('msg_input_password'), trigger: 'blur'}
        ]
      },
      dialogVisible: true
    }
  },
  methods: {
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        this.login()
      })
    },
    login () {
      userServcie.login(this.form.username, this.form.password, (data) => {
        this.$store.commit('setUser', {
          token: data.token,
          uid: data.uid,
          username: data.username,
          isAdmin: data.is_admin
        })
        this.$router.push(this.$route.query.redirect || '/')
      })
    }
  }
}
</script>
