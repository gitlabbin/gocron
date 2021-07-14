<template>
  <el-container>
    <task-sidebar></task-sidebar>
    <el-main>
      <el-form :inline="true" >
        <el-form-item :label="$t('job_id')">
          <el-input v-model.trim="searchParams.task_id"></el-input>
        </el-form-item>
        <el-form-item :label="$t('job_type')">
          <el-select v-model.trim="searchParams.protocol" placeholder="执行方式">
            <el-option :label="$t('option_all')" value=""></el-option>
            <el-option
            v-for="item in protocolList"
            :key="item.value"
            :label="item.label"
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('job_status')">
          <el-select v-model.trim="searchParams.status">
            <el-option :label="$t('option_all')" value=""></el-option>
            <el-option
              v-for="item in statusList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" plain @click="search()">{{ $t('action_search') }}</el-button>
        </el-form-item>
      </el-form>
      <el-row type="flex" justify="end">
        <el-col :span="3">
          <el-button type="danger" icon="el-icon-delete" plain v-if="this.$store.getters.user.isAdmin" @click="clearLog">
            {{ $t('log_clean') }}</el-button>
        </el-col>
        <el-col :span="2">
          <el-button type="info" icon="el-icon-refresh" plain @click="refresh">{{ $t('action_refresh') }}</el-button>
        </el-col>
      </el-row>
      <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="logTotal"
        :page-size="20"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
      </el-pagination>
      <el-table
        :data="logs"
        border
        ref="table"
        style="width: 100%">
        <el-table-column type="expand">
          <template slot-scope="scope">
            <el-form label-position="left">
              <el-form-item>
                <li>{{ $t('job_retry_times') }} {{scope.row.retry_times}} <br></li>
                <li>{{ $t('job_cron') }}: {{scope.row.spec}} <br></li>
                <li>{{ $t('job_task_cmd') }} {{scope.row.command}}</li>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column
          prop="id"
          label="ID"
          width="80">
        </el-table-column>
        <el-table-column
          prop="task_id"
          :label="$t('job_id')"
          width="80">
        </el-table-column>
        <el-table-column
          prop="name"
          :label="$t('job_name')">
        </el-table-column>
        <el-table-column
          prop="protocol"
          :label="$t('job_type')"
          :formatter="formatProtocol"
          width="80">
        </el-table-column>
        <el-table-column
          :label="$t('job_node')"
          width="200">
          <template slot-scope="scope">
            <div v-html="scope.row.hostname">{{scope.row.hostname}}</div>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('time_took')"
          >
          <template slot-scope="scope">
            {{ $t('time_took') }}: {{scope.row.total_time > 0 ? scope.row.total_time : 1}} {{ $t('seconds') }}<br>
            <li>{{ $t('time_start') }}: {{scope.row.start_time | formatTime}}<br></li>
            <span v-if="scope.row.status !== 1"><li>{{ $t('time_end') }}: {{scope.row.end_time | formatTime}}</li></span>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('job_status')">
          <template slot-scope="scope">
            <span style="color:red" v-if="scope.row.status === 0">{{ $t('failure') }}</span>
            <span style="color:green" v-else-if="scope.row.status === 1">{{ $t('in_progress') }}</span>
            <span v-else-if="scope.row.status === 2">{{ $t('succeed') }}</span>
            <span style="color:#4499EE" v-else-if="scope.row.status === 3">{{ $t('canceled') }}</span>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('result')"
          width="140" v-if="this.isAdmin">
          <template slot-scope="scope">
            <el-button type="success" plain
                       v-if="scope.row.status === 2"
                       @click="showTaskResult(scope.row)">{{ $t('view_result') }}</el-button>
            <el-button type="warning" plain
                       v-if="scope.row.status === 0"
                       @click="showTaskResult(scope.row)" >{{ $t('view_result') }}</el-button>
            <el-button type="danger" plain
                       v-if="scope.row.status === 1 && scope.row.protocol === 2"
                       @click="stopTask(scope.row)">{{ $t('stop_job') }}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('result')"
          width="120" v-else>
          <template slot-scope="scope">
            <el-button type="success"
                       v-if="scope.row.status === 2"
                       @click="showTaskResult(scope.row)">{{ $t('view_result') }}</el-button>
            <el-button type="warning"
                       v-if="scope.row.status === 0"
                       @click="showTaskResult(scope.row)" >{{ $t('view_result') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog :title="$t('job_end_result')" :visible.sync="dialogVisible">
        <div>
          <pre>{{currentTaskResult.command}}</pre>
        </div>
        <div>
          <pre>{{currentTaskResult.result}}</pre>
        </div>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import taskSidebar from '../task/sidebar'
import taskLogService from '../../api/taskLog'

export default {
  name: 'task-log',
  data () {
    return {
      logs: [],
      logTotal: 0,
      searchParams: {
        page_size: 20,
        page: 1,
        task_id: '',
        protocol: '',
        status: ''
      },
      isAdmin: this.$store.getters.user.isAdmin,
      dialogVisible: false,
      currentTaskResult: {
        command: '',
        result: ''
      },
      protocolList: [
        {
          value: '1',
          label: 'http'
        },
        {
          value: '2',
          label: 'shell'
        }
      ],
      statusList: [
        {
          value: '1',
          label: this.$t('failure')
        },
        {
          value: '2',
          label: this.$t('in_progress')
        },
        {
          value: '3',
          label: this.$t('succeed')
        },
        {
          value: '4',
          label: this.$t('canceled')
        }
      ]
    }
  },
  components: {taskSidebar},
  created () {
    if (this.$route.query.task_id) {
      this.searchParams.task_id = this.$route.query.task_id
    }
    this.search()
  },
  methods: {
    formatProtocol (row, col) {
      if (row[col.property] === 1) {
        return 'http'
      }
      return 'shell'
    },
    changePage (page) {
      this.searchParams.page = page
      this.search()
    },
    changePageSize (pageSize) {
      this.searchParams.page_size = pageSize
      this.search()
    },
    search (callback = null) {
      taskLogService.list(this.searchParams, (data) => {
        this.logs = data.data
        this.logTotal = data.total

        if (callback) {
          callback()
        }
      })
    },
    clearLog () {
      this.$appConfirm(() => {
        taskLogService.clear(() => {
          this.searchParams.page = 1
          this.search()
        })
      })
    },
    stopTask (item) {
      taskLogService.stop(item.id, item.task_id, () => {
        this.search()
      })
    },
    showTaskResult (item) {
      this.dialogVisible = true
      this.currentTaskResult.command = item.command
      this.currentTaskResult.result = item.result
    },
    refresh () {
      this.search(() => {
        this.$message.success(this.$t('refresh_done'))
      })
    }
  }
}
</script>
<style scoped>
  pre {
    white-space: pre-wrap;
    word-wrap: break-word;
    padding:10px;
    background-color: #4C4C4C;
    color: white;
  }
</style>
