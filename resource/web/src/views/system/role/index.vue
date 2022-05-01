<template>
  <div class="ele-body">
    <el-card shadow="never">
      <!-- 搜索表单 -->
      <el-form :model="where" label-width="77px" class="ele-form-search" @keyup.enter.native="reload" @submit.native.prevent>
        <el-row :gutter="15">
          <el-col :lg="6" :md="12">
            <el-form-item label="角色名称:">
              <el-input clearable v-model="where.name" placeholder="请输入角色名称" />
            </el-form-item>
          </el-col>
          <el-col :lg="6" :md="12">
            <div class="ele-form-actions">
              <el-button type="primary" icon="el-icon-search" class="ele-btn-icon" @click="reload">查询
              </el-button>
              <el-button @click="reset">重置</el-button>
            </div>
          </el-col>
        </el-row>
      </el-form>
      <!-- 数据表格 -->
      <ele-pro-table ref="table" :where="where" :datasource="url" :columns="columns" height="calc(100vh - 270px)" row-key="id" default-expand-all :need-page="false" :parse-data="parseData">
        <!-- 表头工具栏 -->
        <template slot="toolbar">
          <el-button size="small" type="primary" icon="el-icon-plus" class="ele-btn-icon" @click="openEdit(null)" v-if="permission.includes('sys:role:add')">添加
          </el-button>
        </template>

        <!-- 状态列 -->
        <template slot="status" slot-scope="{ row }">
          <el-switch v-model="row.status" @change="editStatus(row)" :active-value="1" :inactive-value="2" />
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
    <role-edit :data="current" :isSuper="isSuper" :roleList="roleList" :visible.sync="showEdit" @done="reload" />
    <!-- 权限分配弹窗 -->
    <role-auth :data="current" :isSuper="isSuper" :visible.sync="showAuth" />
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import RoleEdit from "./role-edit";
import RoleAuth from "./role-auth";

export default {
  name: "SystemRole",
  components: { RoleEdit, RoleAuth },
  computed: {
    ...mapGetters(["permission"]),
  },
  props: {
    isSuper: Boolean,
  },
  data () {
    return {
      // 表格数据接口
      url: "/role/query",
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
          prop: "name",
          label: "角色名称",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 200,
        },
         {
          prop: "code",
          label: "角色编码",
          align: "center",
          showOverflowTooltip: true,
          minWidth: 200,
        },
        {
          prop: "sort",
          label: "排序号",
          align: "center",
          showOverflowTooltip: true,
          width: 100,
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
      where: { is_super: this.isSuper },
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
