<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { OpenInstancesDir, GetInstancesDir, GetSettings, SaveSettings } from "~~/wailsjs/go/main/App";
import type { config } from "~~/wailsjs/go/models";

const instancesDir = ref('')
const backendSettings = ref<config.Settings | null>(null)

const settings = reactive({
  javaPath: '',
  minRam: 4,
  maxRam: 8,
  pingIntervalMinutes: 5,
  launchOnStartup: false,
  minimizeOnLaunch: true,
  showNotifications: true,
  autoUpdate: true,
  theme: 'dark' as 'dark' | 'light' | 'system',
})

const ramOptions = [2, 4, 6, 8, 10, 12, 14, 16]
const pingIntervalOptions = [1, 2, 5, 10, 15, 30, 60]

onMounted(async () => {
  instancesDir.value = await GetInstancesDir()
  const s = await GetSettings()
  backendSettings.value = s
  settings.javaPath = s.java?.preferredPath ?? ''
  settings.minRam = Math.round((s.java?.minMemoryMb ?? 512) / 1024)
  settings.maxRam = Math.round((s.java?.maxMemoryMb ?? 4096) / 1024)
  settings.pingIntervalMinutes = s.pingIntervalMinutes ?? 5
})

const handleSave = async () => {
  const s: config.Settings = {
    schemaVersion: backendSettings.value?.schemaVersion ?? 1,
    pingIntervalMinutes: settings.pingIntervalMinutes,
    java: {
      preferredPath: settings.javaPath || undefined,
      maxMemoryMb: settings.maxRam * 1024,
      minMemoryMb: settings.minRam * 1024,
      extraArgs: backendSettings.value?.java?.extraArgs ?? [],
    },
  }
  await SaveSettings(s)
}

const handleBrowseJava = () => {
  console.log('Browse for Java')
}

const handleBrowseInstances = () => {
  OpenInstancesDir();
}
</script>

<template>
  <div class="mx-auto max-w-3xl space-y-8">
    <!-- Header -->
    <div>
      <h1 class="font-display text-2xl font-bold text-foreground">
        Settings
      </h1>
      <p class="mt-1 text-muted-foreground">
        Configure your launcher preferences
      </p>
    </div>

    <!-- Java Settings -->
    <section class="card p-6">
      <h2 class="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-foreground">
        <Icon name="lucide:coffee" class="h-5 w-5 text-primary" />
        Java Settings
      </h2>
      <div class="space-y-4">
        <div>
          <label class="mb-2 block text-sm font-medium text-foreground">
            Java Path
          </label>
          <div class="flex gap-2">
            <input
              v-model="settings.javaPath"
              type="text"
              class="input flex-1"
            />
            <button class="btn btn-secondary" @click="handleBrowseJava">
              Browse
            </button>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="mb-2 block text-sm font-medium text-foreground">
              Minimum RAM (GB)
            </label>
            <select v-model="settings.minRam" class="input">
              <option v-for="ram in ramOptions" :key="ram" :value="ram">
                {{ ram }} GB
              </option>
            </select>
          </div>
          <div>
            <label class="mb-2 block text-sm font-medium text-foreground">
              Maximum RAM (GB)
            </label>
            <select v-model="settings.maxRam" class="input">
              <option v-for="ram in ramOptions" :key="ram" :value="ram">
                {{ ram }} GB
              </option>
            </select>
          </div>
        </div>
      </div>
    </section>

    <!-- Storage Settings -->
    <section class="card p-6">
      <h2 class="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-foreground">
        <Icon name="lucide:folder" class="h-5 w-5 text-primary" />
        Storage
      </h2>
      <div>
        <label class="mb-2 block text-sm font-medium text-foreground">
          Instances Folder
        </label>
        <div class="flex gap-2">
          <input
            v-model="instancesDir"
            type="text"
            disabled
            class="input flex-1"
          />
          <button class="btn btn-secondary" @click="handleBrowseInstances">
            Browse
          </button>
        </div>
      </div>
    </section>

    <!-- Launcher Settings -->
    <section class="card p-6">
      <h2 class="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-foreground">
        <Icon name="lucide:settings" class="h-5 w-5 text-primary" />
        Launcher
      </h2>
      <div class="space-y-4">
        <label class="flex items-center justify-between">
          <span class="text-sm text-foreground">Launch on system startup</span>
          <button
            class="relative h-6 w-11 rounded-full transition-colors"
            :class="settings.launchOnStartup ? 'bg-primary' : 'bg-secondary'"
            role="switch"
            :aria-checked="settings.launchOnStartup"
            @click="settings.launchOnStartup = !settings.launchOnStartup"
          >
            <span
              class="absolute top-1 h-4 w-4 rounded-full bg-white transition-transform"
              :class="settings.launchOnStartup ? 'left-6' : 'left-1'"
            />
          </button>
        </label>
        <label class="flex items-center justify-between">
          <span class="text-sm text-foreground">Minimize when launching game</span>
          <button
            class="relative h-6 w-11 rounded-full transition-colors"
            :class="settings.minimizeOnLaunch ? 'bg-primary' : 'bg-secondary'"
            role="switch"
            :aria-checked="settings.minimizeOnLaunch"
            @click="settings.minimizeOnLaunch = !settings.minimizeOnLaunch"
          >
            <span
              class="absolute top-1 h-4 w-4 rounded-full bg-white transition-transform"
              :class="settings.minimizeOnLaunch ? 'left-6' : 'left-1'"
            />
          </button>
        </label>
        <label class="flex items-center justify-between">
          <span class="text-sm text-foreground">Show notifications</span>
          <button
            class="relative h-6 w-11 rounded-full transition-colors"
            :class="settings.showNotifications ? 'bg-primary' : 'bg-secondary'"
            role="switch"
            :aria-checked="settings.showNotifications"
            @click="settings.showNotifications = !settings.showNotifications"
          >
            <span
              class="absolute top-1 h-4 w-4 rounded-full bg-white transition-transform"
              :class="settings.showNotifications ? 'left-6' : 'left-1'"
            />
          </button>
        </label>
        <label class="flex items-center justify-between">
          <span class="text-sm text-foreground">Auto-update launcher</span>
          <button
            class="relative h-6 w-11 rounded-full transition-colors"
            :class="settings.autoUpdate ? 'bg-primary' : 'bg-secondary'"
            role="switch"
            :aria-checked="settings.autoUpdate"
            @click="settings.autoUpdate = !settings.autoUpdate"
          >
            <span
              class="absolute top-1 h-4 w-4 rounded-full bg-white transition-transform"
              :class="settings.autoUpdate ? 'left-6' : 'left-1'"
            />
          </button>
        </label>
      </div>
    </section>

    <!-- Repository Settings -->
    <section class="card p-6">
      <h2 class="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-foreground">
        <Icon name="lucide:database" class="h-5 w-5 text-primary" />
        Repositories
      </h2>
      <div>
        <label class="mb-2 block text-sm font-medium text-foreground">
          Ping Interval
        </label>
        <select v-model="settings.pingIntervalMinutes" class="input">
          <option v-for="n in pingIntervalOptions" :key="n" :value="n">
            Every {{ n }} {{ n === 1 ? 'minute' : 'minutes' }}
          </option>
        </select>
        <p class="mt-1 text-xs text-muted-foreground">
          How often the launcher checks repository availability in the background
        </p>
      </div>
    </section>

    <!-- Save Button -->
    <div class="flex justify-end">
      <button class="btn btn-primary" @click="handleSave">
        Save Settings
      </button>
    </div>
  </div>
</template>
