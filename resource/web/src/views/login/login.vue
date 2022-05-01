<template>
  <div :class="['login-wrapper', ['', 'login-form-right', 'login-form-left'][direction]]">
    <el-form
      ref="form"
      size="large"
      :model="form"
      :rules="rules"
      class="login-form ele-bg-white"
      @keyup.enter.native="doSubmit">
      <h4>{{ $t('login.title') }}</h4>
      <el-form-item prop="username">
        <el-input
          clearable
          v-model="form.account"
          prefix-icon="el-icon-user"
          :placeholder="$t('login.username')"/>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          show-password
          v-model="form.password"
          prefix-icon="el-icon-lock"
          :placeholder="$t('login.password')"/>
      </el-form-item>
      <el-form-item prop="captcha">
        <div class="login-input-group">
          <el-input
            clearable
            v-model="form.captcha"
            prefix-icon="el-icon-_vercode"
            :placeholder="$t('login.captcha')"/>
          <img
            alt=""
            v-if="captcha"
            :src="captcha"
            @click="changeCode"
            class="login-captcha"/>
        </div>
      </el-form-item>
      <div class="el-form-item">
        <el-checkbox v-model="form.remember">{{ $t('login.remember') }}</el-checkbox>
      </div>
      <div class="el-form-item">
        <el-button
          size="large"
          type="primary"
          class="login-btn"
          :loading="loading"
          @click="doSubmit">
          {{ loading ? $t('login.loading') : $t('login.login') }}
        </el-button>
      </div>
      <div class="ele-text-center" style="margin-bottom: 10px;">
        <i class="login-oauth-icon el-icon-_qq" style="background: #3492ed;"></i>
        <i class="login-oauth-icon el-icon-_wechat" style="background: #4daf29;"></i>
        <i class="login-oauth-icon el-icon-_weibo" style="background: #CF1900;"></i>
      </div>
    </el-form>
    <div class="login-copyright">Copyright © 2018~2021 www.javaweb.vip. All rights reserved.</div>
    <!-- 多语言切换 -->
    <div style="position: absolute;right: 30px;top: 20px;">
      <el-dropdown placement="bottom" @command="changeLanguage">
        <i class="el-icon-_language" style="font-size: 22px;color: #fff;cursor: pointer;"></i>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="en">
            <span :class="{'ele-text-primary': language === 'en'}">English</span>
          </el-dropdown-item>
          <el-dropdown-item command="zh_CN">
            <span :class="{'ele-text-primary': language === 'zh_CN'}">简体中文</span>
          </el-dropdown-item>
          <el-dropdown-item command="zh_TW">
            <span :class="{'ele-text-primary': language === 'zh_TW'}">繁體中文</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import setting from '@/config/setting';

export default {
  name: 'Login',
  data() {
    return {
      // 登录框方向, 0居中, 1居右, 2居左
      direction: 0,
      // 加载状态
      loading: false,
      // 表单数据
      form: {
        account: 'admin',
        password: '123456',
        captcha: '520',
        key: '',
        remember: true
      },
      // 验证码base64数据
      captcha: '',
      // 验证码内容, 实际项目去掉
      text: ''
    };
  },
  computed: {
    // 表单验证规则
    rules() {
      return {
        account: [
          {required: true, message: this.$t('login.username'), type: 'string', trigger: 'blur'}
        ],
        password: [
          {required: true, message: this.$t('login.password'), type: 'string', trigger: 'blur'}
        ],
        captcha: [
          {required: true, message: this.$t('login.captcha'), type: 'string', trigger: 'blur'}
        ]
      };
    },
    // 当前语言
    language() {
      return this.$i18n.locale;
    }
  },
  mounted() {
    if (setting.takeToken()) {
      this.goHome();
    } else {
      this.changeCode();
    }
  },
  methods: {
    /* 提交 */
    doSubmit() {
      this.$refs.form.validate((valid) => {
        if (!valid) {
          return false;
        }
        this.loading = true;
        this.$http.post('/login', this.form).then((res) => {
          this.loading = false;
          if (res.data.code === 0) {
            this.$message.success('登录成功');
            this.$store.dispatch('user/setToken', {
              token: res.data.data.access_token,
              remember: this.form.remember
            }).then(() => {
              this.goHome();
            });
          } else {
            this.$message.error(res.data.msg);
            // 重新刷新验证码
            this.changeCode()
          }
        });
      });
    },
    /* 跳转到首页 */
    goHome() {
      const query = this.$route.query;
      const path = query && query.from ? query.from : '/';
      this.$router.push(path).catch(() => {
      });
    },
    /* 更换图形验证码 */
    changeCode() {
      // 这里演示的验证码是后端返回base64格式的形式, 如果后端地址直接是图片请参考忘记密码页面
      this.$http.get('/captcha').then(res => {
        if (res.data.code === 0) {
          this.captcha = res.data.data.captcha;
          this.form.key = res.data.data.key;
          // // 实际项目后端一般会返回验证码的key而不是直接返回验证码的内容, 登录用key去验证, 可以根据自己后端接口修改
          // this.text = res.data.text;
          // // 自动回填验证码, 实际项目去掉这个
          // this.form.code = this.text;
          this.$refs.form.clearValidate();
        } else {
          this.$message.error(res.data.msg);
        }
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    /* 切换语言 */
    changeLanguage(lang) {
      this.$i18n.locale = lang;
      localStorage.setItem('i18n-lang', lang);
    }
  }
}
</script>

<style scoped>
/* 背景 */
.login-wrapper {
  padding: 50px 20px;
  position: relative;
  box-sizing: border-box;
  background-image: url("~@/assets/bg-login.png");
  background-repeat: no-repeat;
  background-size: cover;
  min-height: 100vh;
}

.login-wrapper:before {
  content: "";
  background-color: rgba(0, 0, 0, .2);
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

/* 卡片 */
.login-form {
  margin: 0 auto;
  width: 360px;
  max-width: 100%;
  padding: 25px 30px;
  position: relative;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.15);
  box-sizing: border-box;
  border-radius: 4px;
  z-index: 2;
}

.login-form-right .login-form {
  margin: 0 15% 0 auto;
}

.login-form-left .login-form {
  margin: 0 auto 0 15%;
}

.login-form h4 {
  text-align: center;
  margin: 0 0 25px 0;
}

.login-form > .el-form-item {
  margin-bottom: 25px;
}

/* 验证码 */
.login-input-group {
  display: flex;
  align-items: center;
}

.login-input-group ::v-deep .el-input {
  flex: 1;
}

.login-captcha {
  height: 38px;
  width: 102px;
  margin-left: 10px;
  border-radius: 4px;
  border: 1px solid #DCDFE6;
  text-align: center;
  cursor: pointer;
}

.login-captcha:hover {
  opacity: .75;
}

.login-btn {
  display: block;
  width: 100%;
}

/* 第三方登录图标 */
.login-oauth-icon {
  color: #FFF;
  padding: 5px;
  margin: 0 10px;
  font-size: 18px;
  border-radius: 50%;
  cursor: pointer;
}

/* 底部版权 */
.login-copyright {
  color: #eee;
  padding-top: 20px;
  text-align: center;
  position: relative;
  z-index: 1;
}

/* 响应式 */
@media screen and (min-height: 550px) {
  .login-form {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%);
    margin-top: -220px;
  }

  .login-form-right .login-form,
  .login-form-left .login-form {
    left: auto;
    right: 15%;
    transform: translateX(0);
    margin: -220px auto auto auto;
  }

  .login-form-left .login-form {
    right: auto;
    left: 15%;
  }

  .login-copyright {
    position: absolute;
    bottom: 20px;
    right: 0;
    left: 0;
  }
}

@media screen and (max-width: 768px) {
  .login-form-right .login-form,
  .login-form-left .login-form {
    left: 50%;
    right: auto;
    transform: translateX(-50%);
    margin-right: auto;
  }
}
</style>
