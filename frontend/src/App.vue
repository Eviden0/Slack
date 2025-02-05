<script setup lang="ts">
import { ref, onMounted } from "vue";
import global from "./global";
import Sidebar from "./components/Sidebar.vue";
import { useRoute } from "vue-router";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { CheckFileStat, UserHomeDir, Mkdir, WriteFile } from "../wailsjs/go/main/File";

const route = useRoute();
const showLogger = ref(false);

function breadcrumbItems(label: string, separator: string) {
  switch (label) {
    case "/":
      return "Home";
    case "/settings":
      return "Settings";
    default:
      return label.slice(1).split('/').join(separator);
  }
}

interface MsgInfo {
  Level: string
  Msg: string
}

const logArray = [] as string[]

onMounted(() => {
  InitDict()
  const levelClassMap: { [key: string]: string } = {
    "[INF]": "log-info",
    "[WRN]": "log-warning",
    "[ERR]": "log-error",
    "[DEB]": "log-debug",
    "[SUC]": "log-success"
  };
  EventsOn("gologger", (mi: MsgInfo) => {
    const logClass = levelClassMap[mi.Level];
    const logEntry = `<span class="${logClass}">${mi.Level}</span> ${mi.Msg}`;
    logArray.push(logEntry);

    if (logArray.length > global.Logger.length) {
      logArray.shift(); // 移除数组开头的元素
    }
    // 将数组内容拼接成字符串
    global.Logger.value = logArray.join('\n');
  });
})


// 初始化字典
async function InitDict() {
  global.PATH.PortBurstPath = await UserHomeDir() + global.PATH.PortBurstPath
  if (!(await CheckFileStat(global.PATH.PortBurstPath))) {
    if (await Mkdir(global.PATH.PortBurstPath)) {
      Mkdir(global.PATH.PortBurstPath + "/username")
      Mkdir(global.PATH.PortBurstPath + "/password")
      for (const item of global.dict.usernames) {
        WriteFile("txt", `${global.PATH.PortBurstPath}/username/${item.name}.txt`, item.dic.join("\n"))
      }
      WriteFile("txt", `${global.PATH.PortBurstPath}/password/password.txt`, global.dict.passwords.join("\n"))
    }
  }
}
</script>

<template>
  <el-container>
    <el-aside>
      <Sidebar></Sidebar>
    </el-aside>
    <el-container>
      <el-main>
        <!-- 一定要使用插槽否则keey-alive不会生效 -->
        <router-view v-slot="{ Component }">
          <keep-alive>
            <component :is="Component"></component>
          </keep-alive>
        </router-view>
      </el-main>
      <el-footer class="console-log">
        <div>
          <span>
            {{ breadcrumbItems(route.path, ' > ') }}
          </span>
          <el-button link @click="showLogger = true">
            <template #icon>
              <img src="/console.svg">
            </template>
            Console
          </el-button>
        </div>
      </el-footer>
    </el-container>
  </el-container>
  <!-- running logs -->
  <el-drawer v-model="showLogger" direction="ltr" size="50%">
    <template #header>
      <h4>运行日志</h4>
    </template>
    <div class="log-textarea" v-html="global.Logger.value"></div>
  </el-drawer>
</template>

<style>
.el-aside {
  width: 64px;
}
</style>
