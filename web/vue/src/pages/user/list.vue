<template>
  <el-container>
    <el-main>
      <el-row type="flex" justify="end">
        <el-col :span="2">
          <el-button type="primary" icon="el-icon-plus" plain @click="toEdit(null)">{{ $t('action_add') }}</el-button>
        </el-col>
        <el-col :span="2">
          <el-button type="info" icon="el-icon-refresh" plain @click="refresh">{{ $t('action_refresh') }}</el-button>
        </el-col>
      </el-row>
      <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="userTotal"
        :page-size="20"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
      </el-pagination>
      <el-table
        :data="users"
        tooltip-effect="dark"
        border
        style="width: 100%">
        <el-table-column
          prop="id"
          :label="$t('userId')">
        </el-table-column>
        <el-table-column
          prop="name"
          :label="$t('username')">
        </el-table-column>
        <el-table-column
          prop="email"
          :label="$t('email')">
        </el-table-column>
        <el-table-column
          prop="is_admin"
          :formatter="formatRole"
          :label="$t('role')">
        </el-table-column>
        <el-table-column
          :label="$t('user_status')">
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.status"
              :active-value="1"
              :inactive-vlaue="0"
              active-color="#13ce66"
              @change="changeStatus(scope.row)"
              inactive-color="#ff4949">
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column :label="$t('action')" width="300" v-if="this.isAdmin">
          <template slot-scope="scope">
            <el-row>
              <el-button type="primary" plain @click="toEdit(scope.row)">{{ $t('action_edit') }}</el-button>
              <el-button type="success" plain @click="editPassword(scope.row)">{{ $t('change_password') }}</el-button>
              <el-button type="danger" plain @click="remove(scope.row)">{{ $t('action_delete') }}</el-button>
            </el-row>
            <br>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import userService from '../../api/user'
export default {
  name: 'user-list',
  data () {
    return {
      users: [],
      userTotal: 0,
      searchParams: {
        page_size: 20,
        page: 1
      },
      isAdmin: this.$store.getters.user.isAdmin
    }
  },
  created () {
    this.search()
  },
  methods: {
    changeStatus (item) {
      if (item.status) {
        userService.enable(item.id)
      } else {
        userService.disable(item.id)
      }
    },
    formatRole (row, col) {
      if (row[col.property] === 1) {
        return this.$t('admin')
      }
      return this.$t('user')
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
      userService.list(this.searchParams, (data) => {
        this.users = data.data
        this.userTotal = data.total
        if (callback) {
          callback()
        }
      })
    },
    remove (item) {
      this.$appConfirm(() => {
        userService.remove(item.id, () => {
          this.refresh()
        })
      })
    },
    toEdit (item) {
      let path = ''
      if (item === null) {
        path = '/user/create'
      } else {
        path = `/user/edit/${item.id}`
      }
      this.$router.push(path)
    },
    refresh () {
      this.search(() => {
        this.$message.success(this.$t('refresh_done'))
      })
    },
    editPassword (item) {
      this.$router.push(`/user/edit-password/${item.id}`)
    }
  }
}
</script>
