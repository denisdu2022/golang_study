<template>
  <DesktopOutlined v-if="icon==='desktop'"/>
  <HomeOutlined v-else-if="icon==='home'"/>
  <TeamOutlined v-else-if="icon==='team'"/>
  <UserOutlined v-else-if="icon==='user'"/>
  <SettingOutlined v-else-if="icon==='settings'"/>
  <CloudUploadOutlined v-else-if="icon==='cloud'"/>
  <GithubOutlined v-else-if="icon==='git'"/>
  <HistoryOutlined v-else-if="icon==='history'"/>
  <MonitorOutlined v-else-if="icon==='monitor'"/>
  <FileOutlined v-else/>
</template>

<script setup>

const props = defineProps({
  icon: String
})

import {
  CloudUploadOutlined,
  DesktopOutlined,
  FileOutlined,
  GithubOutlined,
  HistoryOutlined,
  HomeOutlined,
  MonitorOutlined,
  SettingOutlined,
  TeamOutlined,
  UserOutlined,
} from '@ant-design/icons-vue';</script>

<style scoped>

</style>