<template>
  <el-container >
    <task-sidebar></task-sidebar>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="180px">
        <el-input v-model="form.id" type="hidden"></el-input>
        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('job_name')" prop="name">
              <el-input v-model.trim="form.name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item :label="$t('job_tag')">
              <el-input v-model.trim="form.tag" :placeholder="$t('job_group_by_tag')"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row v-if="form.level === 1">
          <el-col>
            <el-alert
              :title="$t('jobs_main_sub_description')"
              type="info"
              :closable="false">
            </el-alert>
            <el-alert
              :title="$t('jobs_dependency')"
              type="info"
              :closable="false">
            </el-alert> <br>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="7">
            <el-form-item :label="$t('job_relation')">
              <el-select v-model.trim="form.level" :disabled="form.id !== '' ">
                <el-option
                  v-for="item in levelList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="7" v-if="form.level === 1">
            <el-form-item :label="$t('dependency')">
              <el-select v-model.trim="form.dependency_status">
                <el-option
                  v-for="item in dependencyStatusList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item :label="$t('sub_job_id')" v-if="form.level === 1">
              <el-input v-model.trim="form.dependency_task_id" :placeholder="$t('comma_separate')"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row v-if="form.level === 1">
          <el-col :span="12">
            <el-form-item :label="$t('crontab_expression')" prop="spec">
              <el-input v-model.trim="form.spec"
                        :placeholder="$t('cron_fmt')"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item :label="$t('run_type')">
              <el-select v-model.trim="form.protocol">
                <el-option
                  v-for="item in protocolList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8" v-if="form.protocol === 1 ">
            <el-form-item :label="$t('request_type')">
              <el-select key="http-method" v-model.trim="form.http_method">
                <el-option
                  v-for="item in httpMethods"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8" v-else>
            <el-form-item :label="$t('menu_nodes')">
              <el-select key="shell" v-model="selectedHosts" filterable multiple :placeholder="$t('choose_hint')">
                <el-option
                  v-for="item in hosts"
                  :key="item.id"
                  :label="item.alias + ' - ' + item.name"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="16">
            <el-form-item :label="$t('job_task_cmd')" prop="command">
              <el-input
                type="textarea"
                :rows="5"
                size="medium"
                width="100"
                :placeholder="commandPlaceholder"
                v-model="form.command">
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-alert
              :title="$t('job_timeout_hint')"
              type="info"
              :closable="false">
            </el-alert>
            <el-alert
              :title="$t('job_allow_parallel')"
              type="info"
              :closable="false">
            </el-alert> <br>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item :label="$t('job_timeout_setting')" prop="timeout">
              <el-input v-model.number.trim="form.timeout"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item :label="$t('single_instance')">
              <el-select v-model.trim="form.multi">
                <el-option
                  v-for="item in runStatusList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
        <el-col :span="12">
          <el-form-item :label="$t('job_retry')" prop="retry_times">
            <el-input v-model.number.trim="form.retry_times"
                      :placeholder="$t('retry_hint')"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item :label="$t('retry_interval')" prop="retry_interval">
            <el-input v-model.number.trim="form.retry_interval" :placeholder="$t('retry_interval_hint')"></el-input>
          </el-form-item>
        </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item :label="$t('job_notification')">
              <el-select v-model.trim="form.notify_status">
                <el-option
                  v-for="item in notifyStatusList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8" v-if="form.notify_status !== 1">
            <el-form-item :label="$t('notification_type')">
              <el-select v-model.trim="form.notify_type">
                <el-option
                  v-for="item in notifyTypes"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  >
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8"
                  v-if="form.notify_status !== 1 && form.notify_type === 2">
            <el-form-item :label="$t('receivers')">
              <el-select key="notify-mail" v-model="selectedMailNotifyIds" filterable multiple :placeholder="$t('choose_hint')">
                <el-option
                  v-for="item in mailUsers"
                  :key="item.id"
                  :label="item.username"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="8"
                  v-if="form.notify_status !== 1 && form.notify_type === 3">
            <el-form-item :label="$t('send_to_channel')">
              <el-select key="notify-slack" v-model="selectedSlackNotifyIds" filterable multiple :placeholder="$t('choose_hint')">
                <el-option
                  v-for="item in slackChannels"
                  :key="item.id"
                  :label="item.name"
                  selected="true"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row v-if="form.notify_status === 4">
          <el-col :span="12">
            <el-form-item :label="$t('notify_keywords')" prop="notify_keyword">
              <el-input v-model.trim="form.notify_keyword" :placeholder="$t('keywords_trigger_notify')"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="16">
            <el-form-item :label="$t('remark')">
              <el-input
                type="textarea"
                :rows="3"
                size="medium"
                width="100"
                v-model="form.remark">
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button type="primary" @click="submit">{{ $t('action_save') }}</el-button>
          <el-button @click="cancel">{{ $t('action_cancel') }}</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import taskSidebar from './sidebar'
import taskService from '../../api/task'
import notificationService from '../../api/notification'

export default {
  name: 'task-edit',
  data () {
    return {
      form: {
        id: '',
        name: '',
        tag: '',
        level: 1,
        dependency_status: 1,
        dependency_task_id: '',
        spec: '',
        protocol: 2,
        http_method: 1,
        command: '',
        host_id: '',
        timeout: 0,
        multi: 2,
        notify_status: 1,
        notify_type: 2,
        notify_receiver_id: '',
        notify_keyword: '',
        retry_times: 0,
        retry_interval: 0,
        remark: ''
      },
      formRules: {
        name: [
          {required: true, message: this.$t('msg_input_job_name'), trigger: 'blur'}
        ],
        spec: [
          {required: true, message: this.$t('msg_input_crontab'), trigger: 'blur'}
        ],
        command: [
          {required: true, message: this.$t('msg_input_cmd'), trigger: 'blur'}
        ],
        timeout: [
          {type: 'number', required: true, message: this.$t('msg_input_valid_timeout'), trigger: 'blur'}
        ],
        retry_times: [
          {type: 'number', required: true, message: this.$t('msg_input_valid_retry'), trigger: 'blur'}
        ],
        retry_interval: [
          {type: 'number', required: true, message: this.$t('msg_input_valid_retry_interval'), trigger: 'blur'}
        ],
        notify_keyword: [
          {required: true, message: this.$t('msg_input_job_keywords'), trigger: 'blur'}
        ]
      },
      httpMethods: [
        {
          value: 1,
          label: 'get'
        },
        {
          value: 2,
          label: 'post'
        }
      ],
      protocolList: [
        {
          value: 1,
          label: 'http'
        },
        {
          value: 2,
          label: 'shell'
        }
      ],
      levelList: [
        {
          value: 1,
          label: this.$t('main_job')
        },
        {
          value: 2,
          label: this.$t('sub_job')
        }
      ],
      dependencyStatusList: [
        {
          value: 1,
          label: this.$t('hard_dependency')
        },
        {
          value: 2,
          label: this.$t('soft_dependency')
        }
      ],
      runStatusList: [
        {
          value: 2,
          label: this.$t('yes')
        },
        {
          value: 1,
          label: this.$t('no')
        }
      ],
      notifyStatusList: [
        {
          value: 1,
          label: this.$t('no_notify')
        },
        {
          value: 2,
          label: this.$t('fail_notify')
        },
        {
          value: 3,
          label: this.$t('always_notify')
        },
        {
          value: 4,
          label: this.$t('keyword_match_notify')
        }
      ],
      notifyTypes: [
        {
          value: 2,
          label: this.$t('email_info')
        },
        {
          value: 3,
          label: 'Slack'
        },
        {
          value: 4,
          label: 'WebHook'
        }
      ],
      hosts: [],
      mailUsers: [],
      slackChannels: [],
      selectedHosts: [],
      selectedMailNotifyIds: [],
      selectedSlackNotifyIds: []
    }
  },
  computed: {
    commandPlaceholder () {
      if (this.form.protocol === 1) {
        return this.$t('msg_input_url')
      }

      return this.$t('msg_input_shell_cmd')
    }
  },
  components: {taskSidebar},
  created () {
    const id = this.$route.params.id

    taskService.detail(id, (taskData, hosts) => {
      if (id && !taskData) {
        this.$message.error(this.$tc('no_data'))
        this.cancel()
        return
      }
      this.hosts = hosts || []
      if (!taskData) {
        return
      }
      this.form.id = taskData.id
      this.form.name = taskData.name
      this.form.tag = taskData.tag
      this.form.level = taskData.level
      if (taskData.dependency_status) {
        this.form.dependency_status = taskData.dependency_status
      }
      this.form.dependency_task_id = taskData.dependency_task_id
      this.form.spec = taskData.spec
      this.form.protocol = taskData.protocol
      if (taskData.http_method) {
        this.form.http_method = taskData.http_method
      }
      this.form.command = taskData.command
      this.form.timeout = taskData.timeout
      this.form.multi = taskData.multi ? 1 : 2
      this.form.notify_keyword = taskData.notify_keyword
      this.form.notify_status = taskData.notify_status + 1
      this.form.notify_receiver_id = taskData.notify_receiver_id
      if (taskData.notify_type) {
        this.form.notify_type = taskData.notify_type + 1
      }
      this.form.retry_times = taskData.retry_times
      this.form.retry_interval = taskData.retry_interval
      this.form.remark = taskData.remark
      taskData.hosts = taskData.hosts || []
      if (this.form.protocol === 2) {
        taskData.hosts.forEach((v) => {
          this.selectedHosts.push(v.host_id)
        })
      }

      if (this.form.notify_status > 1) {
        const notifyReceiverIds = this.form.notify_receiver_id.split(',')
        if (this.form.notify_type === 2) {
          notifyReceiverIds.forEach((v) => {
            this.selectedMailNotifyIds.push(parseInt(v))
          })
        } else if (this.form.notify_type === 3) {
          notifyReceiverIds.forEach((v) => {
            this.selectedSlackNotifyIds.push(parseInt(v))
          })
        }
      }
    })

    notificationService.mail((data) => {
      this.mailUsers = data.mail_users
    })

    notificationService.slack((data) => {
      this.slackChannels = data.channels
    })
  },
  methods: {
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        if (this.form.protocol === 2 && this.selectedHosts.length === 0) {
          this.$message.error(this.$tc('msg_choose_node'))
          return false
        }
        if (this.form.notify_status > 1) {
          if (this.form.notify_type === 2 && this.selectedMailNotifyIds.length === 0) {
            this.$message.error(this.$tc('msg_choose_mail_receiver'))
            return false
          }
          if (this.form.notify_type === 3 && this.selectedSlackNotifyIds.length === 0) {
            this.$message.error(this.$tc('msg_choose_slack_channel'))
            return false
          }
        }

        this.save()
      })
    },
    save () {
      if (this.form.protocol === 2 && this.selectedHosts.length > 0) {
        this.form.host_id = this.selectedHosts.join(',')
      }
      if (this.form.notify_status > 1 && this.form.notify_type === 2) {
        this.form.notify_receiver_id = this.selectedMailNotifyIds.join(',')
      }
      if (this.form.notify_status > 1 && this.form.notify_type === 3) {
        this.form.notify_receiver_id = this.selectedSlackNotifyIds.join(',')
      }
      taskService.update(this.form, () => {
        this.$router.push('/task')
      })
    },
    cancel () {
      this.$router.push('/task')
    }
  }
}
</script>
