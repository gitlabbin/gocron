<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="120px" style="width: 500px;">
        <el-form-item :label="$t('new_password')" prop="new_password">
          <el-input v-model="form.new_password" type="password"></el-input>
        </el-form-item>
        <el-form-item :label="$t('confirm_new_password')" prop="confirm_new_password">
          <el-input v-model="form.confirm_new_password" type="password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submit()">{{ $t('action_save') }}</el-button>
          <el-button @click="cancel">{{ $t('action_cancel') }}</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import userService from '../../api/user'
export default {
  name: 'user-edit-password',
  data: function () {
    return {
      form: {
        id: '',
        new_password: '',
        confirm_new_password: ''
      },
      formRules: {
        new_password: [
          {required: true, message: this.$t('msg_input_new_password'), trigger: 'blur'}
        ],
        confirm_new_password: [
          {required: true, message: this.$t('msg_input_new_password_again'), trigger: 'blur'}
        ]
      }
    }
  },
  created () {
    const id = this.$route.params.id
    if (!id) {
      return
    }
    this.form.id = id
  },
  methods: {
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        this.save()
      })
    },
    save () {
      userService.editPassword(this.form, () => {
        this.$router.push('/user')
      })
    },
    cancel () {
      this.$router.push('/user')
    }
  }
}
</script>
