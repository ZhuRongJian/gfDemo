<template>
  <div class="ele-body">
    <el-card shadow="never" body-style="padding: 0;">
      <div class="ele-cell ele-cell-align-top ele-user-message">
        <el-menu
          :mode="mode"
          :default-active="active"
          class="ele-scrollbar-hide">
          <el-menu-item index="all">
            <router-link to="/user/message">
              <span>全部消息</span>
              <el-badge
                v-if="unReadNumAll"
                :value="unReadNumAll"
                class="ele-badge-static"/>
            </router-link>
          </el-menu-item>
          <el-menu-item index="notice">
            <router-link to="/user/message?type=notice">
              <span>系统通知</span>
              <el-badge
                v-if="unReadNum1"
                :value="unReadNum1"
                class="ele-badge-static"/>
            </router-link>
          </el-menu-item>
          <el-menu-item index="message">
            <router-link to="/user/message?type=message">
              <span>用户私信</span>
              <el-badge
                v-if="unReadNum2"
                :value="unReadNum2"
                class="ele-badge-static"/>
            </router-link>
          </el-menu-item>
          <el-menu-item index="todo">
            <router-link to="/user/message?type=todo">
              <span>代办事项</span>
              <el-badge
                v-if="unReadNum3"
                :value="unReadNum3"
                class="ele-badge-static"/>
            </router-link>
          </el-menu-item>
        </el-menu>
        <div class="ele-cell-content">
          <!-- 数据表格 -->
          <ele-pro-table
            ref="table"
            :loading="loading"
            :columns="columns"
            :datasource="data"
            :selection.sync="selection"
            tool-class="ele-toolbar-actions"
            @refresh="query">
            <template slot="toolbar">
              <el-button
                size="small"
                type="primary"
                @click="read">标记已读
              </el-button>
              <el-button
                size="small"
                type="primary"
                @click="readAll">全部已读
              </el-button>
              <el-button
                size="small"
                type="danger"
                @click="removeBatch">删除消息
              </el-button>
            </template>
            <template slot="state" slot-scope="{row}">
              <span :class="['ele-text-danger', 'ele-text-info'][row.state]">
                {{ ['未读', '已读'][row.state] }}
              </span>
            </template>
            <template slot="action" slot-scope="{row}">
              <el-link
                type="primary"
                :underline="false"
                icon="el-icon-view"
                @click="view(row)">查看
              </el-link>
              <el-popconfirm
                class="ele-action"
                title="确定要删除此消息吗？"
                @confirm="remove(row)">
                <el-link
                  type="danger"
                  slot="reference"
                  :underline="false"
                  icon="el-icon-delete">删除
                </el-link>
              </el-popconfirm>
            </template>
          </ele-pro-table>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'UserMessage',
  data() {
    return {
      // 导航选中
      active: 'all',
      // 加载状态
      loading: true,
      // 全部数据
      allData: [],
      // 列表显示数据
      data: [],
      // 表格列配置
      columns: [
        {
          columnKey: 'selection',
          type: 'selection',
          width: 45,
          align: 'center'
        },
        {
          columnKey: 'index',
          type: 'index',
          width: 45,
          align: 'center',
          showOverflowTooltip: true
        },
        {
          prop: 'title',
          label: '标题内容',
          showOverflowTooltip: true,
          minWidth: 110
        },
        {
          prop: 'time',
          label: '时间',
          align: 'center',
          showOverflowTooltip: true,
          width: 140
        },
        {
          prop: 'state',
          label: '状态',
          align: 'center',
          showOverflowTooltip: true,
          width: 80,
          slot: 'state'
        },
        {
          columnKey: 'action',
          label: '操作',
          align: 'center',
          showOverflowTooltip: true,
          width: 140,
          resizable: false,
          slot: 'action'
        }
      ],
      // 列表选中数据
      selection: []
    };
  },
  computed: {
    // 小屏幕水平导航
    mode() {
      return this.$store.state.theme.screenWidth < 768 ? 'horizontal' : 'vertical';
    },
    // 全部未读数量
    unReadNumAll() {
      return this.allData.filter((d) => d.state === 0).length;
    },
    // 通知未读数量
    unReadNum1() {
      return this.allData.filter((d) => d.state === 0 && d.type === 'notice').length;
    },
    // 私信未读数量
    unReadNum2() {
      return this.allData.filter((d) => d.state === 0 && d.type === 'message').length;
    },
    // 代办未读数量
    unReadNum3() {
      return this.allData.filter((d) => d.state === 0 && d.type === 'todo').length;
    }
  },
  mounted() {
    const type = this.$route.query.type;
    if (type) {
      this.active = type;
    }
    this.query();
  },
  methods: {
    /* 查询全部数据 */
    query() {
      this.loading = true;
      this.$http.get('https://images.vue.pro.javaweb.vip/message.json').then((res) => {
        this.loading = false;
        if (res.data.code === 0) {
          this.allData = res.data.data;
          this.changeType();
        } else {
          this.$message.error(res.data.msg);
        }
      }).catch((e) => {
        this.loading = false;
        this.$message.error(e.message);
      });
    },
    /* 切换消息类型 */
    changeType(index) {
      if (index) {
        this.active = index;
      }
      if (this.active === 'all') {
        this.data = [].concat(this.allData);
      } else {
        this.data = this.allData.filter((d) => d.type === this.active);
      }
      this.selection = [];
    },
    /* 查看 */
    view(row) {
      this.$message.info(row.title);
    },
    /* 删除单个 */
    remove(row) {
      this.allData.splice(this.allData.findIndex((t) => t.id === row.id), 1);
      this.$message.success('删除成功');
      this.changeType();
    },
    /* 批量删除 */
    removeBatch() {
      if (!this.selection.length) {
        return this.$message.error('请至少选择一条数据');
      }
      this.$confirm('确定要删除选中的消息吗?', '提示', {
        type: 'warning'
      }).then(() => {
        this.selection.forEach((d) => {
          this.allData.splice(this.allData.findIndex((t) => t.id === d.id), 1);
        });
        this.changeType();
        this.$message.success('删除成功');
      }).catch(() => {
      });
    },
    /* 标记已读 */
    read() {
      if (!this.selection.length) {
        return this.$message.error('请至少选择一条数据');
      }
      this.selection.forEach((d) => {
        const index = this.allData.findIndex((t) => t.id === d.id);
        this.allData[index].state = 1;
      });
      this.changeType();
    },
    /* 全部标记已读 */
    readAll() {
      this.allData.forEach((d) => {
        d.state = 1;
      });
      this.changeType();
    }
  },
  watch: {
    $route() {
      const type = this.$route.query.type;
      this.changeType(type || 'all');
    }
  }
}
</script>

<style scoped>
.ele-user-message ::v-deep .el-menu {
  width: 151px;
  flex-shrink: 0;
}

.ele-user-message .ele-cell-content {
  padding: 15px;
  box-sizing: border-box;
  overflow: auto;
}

.ele-user-message .ele-badge-static {
  margin-left: 10px;
}

@media screen and (min-width: 768px) {
  .ele-user-message ::v-deep .el-menu .el-menu-item {
    text-align: right;
  }

  .ele-user-message ::v-deep .el-menu .el-menu-item:first-child {
    margin-top: 15px;
  }
}

@media screen and (max-width: 768px) {
  .ele-user-message {
    display: block;
  }

  .ele-user-message ::v-deep .el-menu {
    width: auto;
    text-align: center;
    white-space: nowrap;
    overflow: auto;
  }

  .ele-user-message ::v-deep .el-menu .el-menu-item {
    height: 45px;
    line-height: 45px;
    padding: 0 5px;
    display: inline-block;
    float: none;
  }

  .ele-user-message .ele-badge-static {
    margin-left: 3px;
  }
}
</style>
