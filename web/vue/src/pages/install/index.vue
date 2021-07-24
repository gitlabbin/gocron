<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="form.db_type==='sqlite3'?formSqliteRules:formRules" label-width="100px" style="width: 700px;">
        <h3>{{ $t('database_setting') }}</h3>
        <el-form-item :label="$t('select_database')" prop="db_type">
          <el-select v-model.trim="form.db_type" @change="update_port">
            <el-option
              v-for="item in dbList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-row>
          <el-col :span="12" :hidden="form.db_type==='sqlite3'">
            <el-form-item
              :label="$t('hostname')"
              prop="db_host">
              <el-input v-model="form.db_host"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12" :hidden="form.db_type==='sqlite3'">
            <el-form-item :label="$t('port')" prop="db_port">
              <el-input v-model.number="form.db_port"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12" :hidden="form.db_type==='sqlite3'">
            <el-form-item :label="$t('db_user')" prop="db_username">
              <el-input v-model="form.db_username"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12" :hidden="form.db_type==='sqlite3'">
            <el-form-item :label="$t('db_password')" prop="db_password">
              <el-input v-model="form.db_password" type="password"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('db_name')" prop="db_name">
              <el-input v-model="form.db_name" :placeholder="$t('db_hint')"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('table_prefix')" prop="db_table_prefix">
              <el-input v-model="form.db_table_prefix"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <h3>{{ $t('system_admin_setting') }}</h3>
        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('admin_user')" prop="admin_username">
              <el-input v-model="form.admin_username"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('admin_email')" prop="admin_email">
              <el-input v-model="form.admin_email"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('admin_password')" prop="admin_password">
              <el-input v-model="form.admin_password" type="password"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('confirm_password')" prop="confirm_admin_password">
              <el-input v-model="form.confirm_admin_password" type="password"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button type="primary" @click="submit()">{{ $t('install') }}</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import installService from '../../api/install'
export default {
  name: 'index',
  data () {
    return {
      form: {
        db_type: 'mysql',
        db_host: '127.0.0.1',
        db_port: 3306,
        db_username: '',
        db_password: '',
        db_name: '',
        db_table_prefix: '',
        admin_username: '',
        admin_password: '',
        confirm_admin_password: '',
        admin_email: ''
      },
      formRules: {
        db_type: [
          {required: true, message: this.$t('msg_choose_database'), trigger: 'blur'}
        ],
        db_host: [
          {required: true, message: this.$t('msg_input_database_host'), trigger: 'blur'}
        ],
        db_port: [
          {type: 'number', required: true, message: this.$t('msg_input_database_port'), trigger: 'blur'}
        ],
        db_username: [
          {required: true, message: this.$t('msg_input_database_user'), trigger: 'blur'}
        ],
        db_password: [
          {required: true, message: this.$t('msg_input_db_password'), trigger: 'blur'}
        ],
        db_name: [
          {required: true, message: this.$t('msg_input_database_name'), trigger: 'blur'}
        ],
        admin_username: [
          {required: true, message: this.$t('msg_admin_username'), trigger: 'blur'}
        ],
        admin_email: [
          {type: 'email', required: true, message: this.$t('msg_admin_email'), trigger: 'blur'}
        ],
        admin_password: [
          {required: true, message: this.$t('msg_input_admin_password'), trigger: 'blur'},
          {min: 6, message: this.$t('msg_length_limit'), trigger: 'blur'}
        ],
        confirm_admin_password: [
          {required: true, message: this.$t('msg_repeat_admin_password'), trigger: 'blur'},
          {min: 6, message: this.$t('msg_length_limit'), trigger: 'blur'}
        ]
      },
      formSqliteRules: {
        db_type: [
          {required: true, message: this.$t('msg_choose_database'), trigger: 'blur'}
        ],
        db_host: [
          {required: false, message: this.$t('msg_input_database_host'), trigger: 'blur'}
        ],
        db_port: [
          {type: 'number', required: false, message: this.$t('msg_input_database_port'), trigger: 'blur'}
        ],
        db_username: [
          {required: false, message: this.$t('msg_input_database_user'), trigger: 'blur'}
        ],
        db_password: [
          {required: false, message: this.$t('msg_input_db_password'), trigger: 'blur'}
        ],
        db_name: [
          {required: true, message: this.$t('msg_input_database_name'), trigger: 'blur'}
        ],
        admin_username: [
          {required: true, message: this.$t('msg_admin_username'), trigger: 'blur'}
        ],
        admin_email: [
          {type: 'email', required: true, message: this.$t('msg_admin_email'), trigger: 'blur'}
        ],
        admin_password: [
          {required: true, message: this.$t('msg_input_admin_password'), trigger: 'blur'},
          {min: 6, message: this.$t('msg_length_limit'), trigger: 'blur'}
        ],
        confirm_admin_password: [
          {required: true, message: this.$t('msg_repeat_admin_password'), trigger: 'blur'},
          {min: 6, message: this.$t('msg_length_limit'), trigger: 'blur'}
        ]
      },
      dbList: [
        {
          value: 'sqlite3',
          label: 'Sqlite3'
        },
        {
          value: 'mysql',
          label: 'MySQL'
        },
        {
          value: 'postgres',
          label: 'PostgreSql'
        }
      ],
      default_ports: {
        'sqlite3': 1,
        'mysql': 3306,
        'postgres': 5432
      }
    }
  },
  methods: {
    update_port (dbType) {
      console.log(dbType)
      console.log(this.default_ports[dbType])
      this.form['db_port'] = this.default_ports[dbType]
      console.log(this.form['db_port'])
      if (dbType === 'sqlite3') {
        this.form['db_username'] = 'no_need'
        this.form['db_password'] = 'no-need'
      } else {
        this.form['db_username'] = ''
        this.form['db_password'] = ''
      }
    },
    submit () {
      this.$refs['form'].validate((valid, errors) => {
        if (!valid) {
          console.log(errors)
          return false
        }
        this.save()
      })
    },
    save () {
      installService.store(this.form, () => {
        this.$router.push('/')
      })
    }
  }
}
</script>
