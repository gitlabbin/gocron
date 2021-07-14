<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="width: 500px;">
        <el-form-item>
          <el-input v-model="form.id" type="hidden"></el-input>
        </el-form-item>
        <el-form-item :label="$t('node_name')" prop="alias">
          <el-input v-model="form.alias"></el-input>
        </el-form-item>
        <el-form-item :label="$t('node_host')" prop="name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item :label="$t('node_host_port')" prop="port">
          <el-input v-model.number="form.port"></el-input>
        </el-form-item>
        <el-form-item :label="$t('node_description')">
          <el-input
            type="textarea"
            :rows="5"
            size="medium"
            width="100"
            v-model="form.remark">
          </el-input>
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
import hostService from '../../api/host'
export default {
  name: 'edit',
  data: function () {
    return {
      form: {
        id: '',
        name: '',
        port: 5921,
        alias: '',
        remark: ''
      },
      formRules: {
        name: [
          {required: true, message: this.$t('msg_input_host_name'), trigger: 'blur'}
        ],
        port: [
          {required: true, message: this.$t('msg_input_port'), trigger: 'blur'},
          {type: 'number', message: this.$t('msg_invalid_port')}
        ],
        alias: [
          {required: true, message: this.$t('msg_input_node_name'), trigger: 'blur'}
        ]
      }
    }
  },
  created () {
    const id = this.$route.params.id
    if (!id) {
      return
    }
    hostService.detail(id, (data) => {
      if (!data) {
        this.$message.error(this.$t('msg_no_data'))
        this.cancel()
        return
      }
      this.form.id = data.id
      this.form.name = data.name
      this.form.port = data.port
      this.form.alias = data.alias
      this.form.remark = data.remark
    })
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
      hostService.update(this.form, () => {
        this.$router.push('/host')
      })
    },
    cancel () {
      this.$router.push('/host')
    }
  }
}
</script>
