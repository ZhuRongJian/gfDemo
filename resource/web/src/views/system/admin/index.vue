<template>
  <div class="ele-body">
    <el-card shadow="never">
      <!-- 数据表格 -->
      <ele-pro-table ref="table" :where="where" :datasource="url" :columns="columns" height="calc(100vh - 270px)" row-key="id" default-expand-all :need-page="false" :parse-data="parseData">
        <!-- 表头工具栏 -->
        <template slot="toolbar">
          <el-button size="small" type="primary" icon="el-icon-plus" class="ele-btn-icon" @click="openEdit(null)" v-if="permission.includes('sys:role:add')">添加
          </el-button>
        </template>

        <template slot="role_name" slot-scope="{ row }">
          <el-tag type="success" size="mini">
            {{ row.role_name }}
          </el-tag>
        </template>


        <!-- 操作列 -->
        <template slot="action" slot-scope="{ row }">
          <el-link type="primary" :underline="false" icon="el-icon-edit" @click="openEdit(row)" v-if="permission.includes('sys:role:edit')">修改
          </el-link>
          <el-link type="primary" :underline="false" icon="el-icon-finished" @click="openAuth(row)" v-if="permission.includes('sys:role:permission')">分配权限
          </el-link>
          <el-popconfirm class="ele-action" title="确定要删除吗？" @confirm="remove(row)">
            <el-link type="danger" slot="reference" :underline="false" icon="el-icon-delete" v-if="permission.includes('sys:role:delete')">删除
            </el-link>
          </el-popconfirm>
        </template>

      </ele-pro-table>
    </el-card>
    <!-- 编辑弹窗 -->
    <!-- <role-edit :data="current" :isSuper="isSuper" :roleList="roleList" :visible.sync="showEdit" @done="reload" /> -->
    <!-- 权限分配弹窗 -->
    <!-- <role-auth :data="current" :isSuper="isSuper" :visible.sync="showAuth" /> -->
  </div>
</template>

<script>
import { mapGetters } from "vuex";
// import RoleEdit from "./role-edit";
// import RoleAuth from "./role-auth";

export default {
  name: "SystemAdmin",
  // components: { RoleEdit, RoleAuth },
  computed: {
    ...mapGetters(["permission"]),
  },
  data () {
    return {
      // 表格数据接口
      url: "/admin/query",
      // 表格列配置
      columns: [
        {
          columnKey: "index",
          type: "index",
          width: 45,
          align: "center",
          showOverflowTooltip: true,
        },
        {
          prop: "nickname",
          label: "姓名",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 100,
        },
         {
          prop: "mobile",
          label: "手机号",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 200,
        },
        {
          prop: "account",
          label: "登录账号",
          align: "center",
          showOverflowTooltip: true,
          width: 200,
        },
        {
          prop: "role_name",
          label: "角色",
          align: "center",
          showOverflowTooltip: true,
          width: 200,
          slot: "role_name",
        },
         {
          prop: "partners",
          label: "所属门店",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 200,
        },
        {
          prop: "note",
          label: "备注",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 200,
        },
        {
          prop: "created_at",
          label: "创建时间",
          sortable: "custom",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 160,
          formatter: (row, column, cellValue) => {
            return this.$util.toDateString(cellValue);
          },
        },
        {
          prop: "updated_at",
          label: "更新时间",
          sortable: "custom",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 160,
          formatter: (row, column, cellValue) => {
            return this.$util.toDateString(cellValue);
          },
        },
        {
          columnKey: "action",
          label: "操作",
          width: 230,
          align: "center",
          resizable: false,
          slot: "action",
          fixed: "right",
        },
      ],
      // 表格搜索条件
      where: {  },
      // 表格选中数据
      selection: [],
      // 当前编辑数据
      current: null,
      // 是否显示编辑弹窗
      showEdit: false,
      // 是否显示导入弹窗
      showAuth: false,
      roleList: [],
    };
  },
  methods: {
    /* 解析接口返回数据 */
    parseData (res) {
      res.data = this.$util.toTreeData(res.data, "id", "pid");
      this.roleList = res.data;
      return {
        code: res.code,
        msg: res.msg,
        data: {
          records: res.data,
        },
      };
    },
    /* 刷新表格 */
    reload () {
      this.$refs.table.reload({ where: this.where });
    },
    /* 重置搜索 */
    reset () {
      this.where = {};
      this.reload();
    },
    /* 显示编辑 */
    openEdit (row) {
      this.showEdit = true;
      if(row){
          this.current = row;
      }else{
        this.current ={}
      }
    },
    /* 显示分配权限 */
    openAuth (row) {
      this.current = row;
      this.showAuth = true;
    },
    /* 删除 */
    remove (row) {
      const loading = this.$loading({ lock: true });
      this.$http
        .post("/role/delete/" + [row.id])
        .then((res) => {
          if (res.data.code === 0) {
            this.$message.success("操作成功");
            this.reload();
          }
        })
        .finally(() => {
          loading.close();
        });
    },
    /* 更改状态 */
    editStatus (row) {
      const loading = this.$loading({ lock: true });
      let params = Object.assign({
        id: row.id,
        status: row.status,
      });
      this.$http
        .post("/role/status", params)
        .then(() => {
          this.$message({ type: "success", message: "操作成功" });
        })
        .catch(() => {
          row.status = !row.status ? 1 : 0;
        }).finally(() => {
          loading.close();
        });
    },
  },
};
</script>

<style scoped></style>
