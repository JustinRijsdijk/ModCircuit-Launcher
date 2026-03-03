<script setup lang="ts">
import {ref} from "vue";

type ViewType = 'home' | 'modpacks' | 'downloads' | 'instances' | 'settings' | 'about'

const currentView = ref<ViewType>('home')
const sidebarCollapsed = ref(false)

const handleNavigation = (view: ViewType) => {
  currentView.value = view
}

const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value
}
</script>

<template>
  <div class="flex h-screen flex-col overflow-hidden bg-background">
    <!-- Title Bar -->
    <LauncherTitleBar />

    <!-- Main Content -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <LauncherSidebar
          :current-view="currentView"
          :collapsed="sidebarCollapsed"
          @navigate="handleNavigation"
          @toggle="toggleSidebar"
      />

      <!-- Content Area -->
      <main class="flex-1 overflow-y-auto p-6">
        <slot />
      </main>
    </div>
  </div>
</template>

<style scoped>

</style>