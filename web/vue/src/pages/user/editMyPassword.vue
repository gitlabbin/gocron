<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="120px" style="width: 500px;">
        <el-form-item :label="$t('old_password')" prop="old_password">
          <el-input v-model="form.old_password" type="password"></el-input>
        </el-form-item>
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
  name: 'user-edit-my-password',
  data: function () {
    return {
      form: {
        old_password: '',
        new_password: '',
        confirm_new_password: ''
      },
      formRules: {
        old_password: [
          {required: true, message: '请输入原密码', trigger: 'blur'}
        ],
        new_password: [
          {required: true, message: '请输入新密码', trigger: 'blur'}
        ],
        confirm_new_password: [
          {required: true, message: '请再次输入新密码', trigger: 'blur'}
        ]
      }
    }
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
      userService.editMyPassword(this.form, () => {
        this.$router.back()
      })
    },
    cancel () {
      this.$router.back()
    }
  }
}
</script>
